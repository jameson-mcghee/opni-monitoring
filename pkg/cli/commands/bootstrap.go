package commands

import (
	"context"
	"time"

	"github.com/rancher/opni-monitoring/pkg/config/meta"
	"github.com/rancher/opni-monitoring/pkg/config/v1beta1"
	"github.com/rancher/opni-monitoring/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/polymorphichelpers"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

func BuildBootstrapCmd() *cobra.Command {
	lg := logger.New()

	var namespace, kubeconfig, gatewayAddress, token string
	var pins []string
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "Bootstrap an agent",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			rules := clientcmd.NewDefaultClientConfigLoadingRules()
			rules.ExplicitPath = kubeconfig
			apiConfig, err := rules.Load()
			if err != nil {
				lg.With(
					zap.Error(err),
				).Fatal("Failed to load kubeconfig")
			}
			restConfig, err := clientcmd.NewDefaultClientConfig(
				*apiConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
			if err != nil {
				lg.With(
					zap.Error(err),
				).Fatal("Failed to load kubeconfig")
			}
			clientset := kubernetes.NewForConfigOrDie(restConfig)

			lg.Info("Checking for pending agents...")

			dep, err := getAgentDeployment(ctx, clientset, namespace)
			if err != nil {
				lg.With(
					zap.Error(err),
				).Fatal("Failed to look up agent deployment")
			}

			// If any replicas are available (not unavailable), the agent has already
			// been bootstrapped.
			if dep.Status.UnavailableReplicas < pointer.Int32Deref(dep.Spec.Replicas, 1) {
				lg.Info("Agent has already been bootstrapped")
				return
			}

			// If the agent config secret exists, the agent has already been
			// bootstrapped
			_, err = clientset.CoreV1().
				Secrets(namespace).
				Get(ctx, "agent-config", metav1.GetOptions{})
			if err == nil {
				lg.Info("Agent has already been bootstrapped")
				return
			}
			if !k8serrors.IsNotFound(err) {
				lg.With(
					zap.Error(err),
				).Fatal("Failed to look up agent config secret")
			}
			lg.Info("Bootstrapping agent...")

			agentConfig := v1beta1.AgentConfig{
				TypeMeta: meta.TypeMeta{
					APIVersion: "v1beta1",
					Kind:       "AgentConfig",
				},
				Spec: v1beta1.AgentConfigSpec{
					ListenAddress:    ":8080",
					GatewayAddress:   gatewayAddress,
					IdentityProvider: "kubernetes",
					Storage: v1beta1.StorageSpec{
						Type: v1beta1.StorageTypeSecret,
					},
					Bootstrap: &v1beta1.BootstrapSpec{
						Token: token,
						Pins:  pins,
					},
				},
			}

			configData, err := yaml.Marshal(agentConfig)
			if err != nil {
				lg.With(
					zap.Error(err),
				).Fatal("Failed to marshal agent config")
			}

			secret := corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "agent-config",
					Namespace: namespace,
				},
				Data: map[string][]byte{
					"config.yaml": configData,
				},
			}
			_, err = clientset.CoreV1().
				Secrets(namespace).
				Create(ctx, &secret, metav1.CreateOptions{})
			if err != nil {
				lg.With(
					zap.Error(err),
				).Fatal("Failed to create agent config secret")
			}

			lg.Info("Agent bootstrapped successfully.")

			// Check if the deployment needs to be restarted
			for _, cond := range dep.Status.Conditions {
				if cond.Type == appsv1.DeploymentProgressing &&
					cond.Status == corev1.ConditionFalse &&
					cond.Reason == "ProgressDeadlineExceeded" {
					lg.Info("Agent deployment is stuck, restarting...")
					if err := doRolloutRestart(ctx, clientset, dep); err != nil {
						lg.With(
							zap.Error(err),
						).Error("Failed to restart agent deployment. You might have to restart the agent manually.")
					}
				}
			}

			lg.Info("Waiting for agent to be ready...")

			for {
				dep, err = getAgentDeployment(ctx, clientset, namespace)
				if err != nil {
					lg.With(
						zap.Error(err),
					).Error("Failed to look up agent deployment")
				}
				if dep.Status.AvailableReplicas == pointer.Int32Deref(dep.Spec.Replicas, 1) {
					break
				}
				time.Sleep(time.Second)
			}

			lg.Info("Done.")
		},
	}

	bootstrapCmd.Flags().StringVarP(&gatewayAddress, "address", "a", "", "Gateway address")
	bootstrapCmd.Flags().StringVarP(&token, "token", "t", "", "Token to use for bootstrapping")
	bootstrapCmd.Flags().StringSliceVar(&pins, "pin", []string{}, "Gateway server public key to pin (repeatable)")
	bootstrapCmd.Flags().StringVarP(&namespace, "namespace", "n", "opni-monitoring-agent", "Namespace where the agent is installed")
	bootstrapCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to kubeconfig file (optional)")

	bootstrapCmd.MarkFlagRequired("address")
	bootstrapCmd.MarkFlagRequired("token")
	bootstrapCmd.MarkFlagRequired("pin")

	return bootstrapCmd
}

func doRolloutRestart(ctx context.Context, clientset *kubernetes.Clientset, dep *appsv1.Deployment) error {
	patch, err := polymorphichelpers.ObjectRestarterFn(dep)
	if err != nil {
		return err
	}
	_, err = clientset.AppsV1().
		Deployments(dep.Namespace).
		Patch(ctx,
			dep.Name, types.StrategicMergePatchType, patch, metav1.PatchOptions{
				FieldManager: "kubectl-rollout",
			},
		)
	return err
}

func getAgentDeployment(ctx context.Context, clientset *kubernetes.Clientset, namespace string) (*appsv1.Deployment, error) {
	return clientset.AppsV1().
		Deployments(namespace).
		Get(ctx, "opni-monitoring-agent", metav1.GetOptions{})
}
