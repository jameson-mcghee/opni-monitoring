package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rancher/opni-monitoring/pkg/core"
	"github.com/rancher/opni-monitoring/pkg/logger"
	"github.com/rancher/opni-monitoring/pkg/management"
	"github.com/rancher/opni-monitoring/pkg/test"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

//#region Test Setup

var _ = Describe("Management API Cluster Management Tests", Ordered, func() {
	var environment *test.Environment
	var client management.ManagementClient
	var fingerprint string
	BeforeAll(func() {
		fmt.Println("Starting test environment")
		environment = &test.Environment{
			TestBin: "../../../testbin/bin",
			Logger:  logger.New().Named("test"),
		}
		Expect(environment.Start()).To(Succeed())
		client = environment.NewManagementClient()
		Expect(json.Unmarshal(test.TestData("fingerprints.json"), &testFingerprints)).To(Succeed())

		token, err := client.CreateBootstrapToken(context.Background(), &management.CreateBootstrapTokenRequest{
			Ttl: durationpb.New(time.Minute),
		})
		Expect(err).NotTo(HaveOccurred())

		certsInfo, err := client.CertsInfo(context.Background(), &emptypb.Empty{})
		Expect(err).NotTo(HaveOccurred())
		fingerprint = certsInfo.Chain[len(certsInfo.Chain)-1].Fingerprint
		Expect(fingerprint).NotTo(BeEmpty())

		port, errC := environment.StartAgent("test-cluster-id", token, []string{fingerprint})
		promAgentPort := environment.StartPrometheus(port)
		Expect(promAgentPort).NotTo(BeZero())
		Consistently(errC).ShouldNot(Receive())
	})

	AfterAll(func() {
		fmt.Println("Stopping test environment")
		Expect(environment.Stop()).To(Succeed())
	})

	//#endregion

	//#region Happy Path Tests

	events := make(chan *management.WatchEvent, 1000)
	It("should handle watching create and delete events", func() {
		stream, err := client.WatchClusters(context.Background(), &management.WatchClustersRequest{
			KnownClusters: &core.ReferenceList{},
		})
		Expect(err).NotTo(HaveOccurred())
		go func() {
			defer close(events)
			for {
				event, err := stream.Recv()
				if err != nil {
					return
				}
				events <- event
			}
		}()
	})

	It("can get information about a specific cluster", func() {
		clusterInfo, err := client.GetCluster(context.Background(), &core.Reference{
			Id: "test-cluster-id",
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(clusterInfo.Id).To(Equal("test-cluster-id"))
		Expect(clusterInfo.Labels).To(BeNil())
	})

	It("can edit the label a cluster is using", func() {
		_, err := client.EditCluster(context.Background(), &management.EditClusterRequest{
			Cluster: &core.Reference{
				Id: "test-cluster-id",
			},
			Labels: map[string]string{
				"i": "999",
			},
		})
		Expect(err).NotTo(HaveOccurred())

		clusterInfo, err := client.GetCluster(context.Background(), &core.Reference{
			Id: "test-cluster-id",
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(clusterInfo.Id).To(Equal("test-cluster-id"))
		Expect(clusterInfo.Labels).To(HaveKeyWithValue("i", "999"))
	})

	It("can list all clusters using the same label", func() {
		token2, err := client.CreateBootstrapToken(context.Background(), &management.CreateBootstrapTokenRequest{
			Ttl: durationpb.New(time.Minute),
		})
		Expect(err).NotTo(HaveOccurred())

		certsInfo, err := client.CertsInfo(context.Background(), &emptypb.Empty{})
		Expect(err).NotTo(HaveOccurred())
		fingerprint = certsInfo.Chain[len(certsInfo.Chain)-1].Fingerprint
		Expect(fingerprint).NotTo(BeEmpty())

		_, errC := environment.StartAgent("test-cluster-id-2", token2, []string{fingerprint})
		Consistently(errC).ShouldNot(Receive())

		_, err = client.EditCluster(context.Background(), &management.EditClusterRequest{
			Cluster: &core.Reference{
				Id: "test-cluster-id-2",
			},
			Labels: map[string]string{
				"i": "999",
			},
		})
		Expect(err).NotTo(HaveOccurred())

		clusterInfo, err := client.ListClusters(context.Background(), &management.ListClustersRequest{
			MatchLabels: &core.LabelSelector{
				MatchLabels: map[string]string{
					"i": "999",
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())

		clusterItems := clusterInfo.Items
		Expect(clusterItems).To(HaveLen(2))
		for _, clusterItem := range clusterItems {
			Expect(clusterItem.Id).To(Or(Equal("test-cluster-id"), Equal("test-cluster-id-2")))
			Expect(clusterItem.Labels).To(HaveKeyWithValue("i", "999"))
		}

		_, err = client.DeleteCluster(context.Background(), &core.Reference{
			Id: "test-cluster-id-2",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	It("should handle watching create and delete events", func() {
		token3, err := client.CreateBootstrapToken(context.Background(), &management.CreateBootstrapTokenRequest{
			Ttl: durationpb.New(time.Minute),
		})
		Expect(err).NotTo(HaveOccurred())

		certsInfo, err := client.CertsInfo(context.Background(), &emptypb.Empty{})
		Expect(err).NotTo(HaveOccurred())
		fingerprint = certsInfo.Chain[len(certsInfo.Chain)-1].Fingerprint
		Expect(fingerprint).NotTo(BeEmpty())

		_, errC := environment.StartAgent("test-cluster-id-3", token3, []string{fingerprint})
		Consistently(errC).ShouldNot(Receive())

		Eventually(events).Should(Receive(WithTransform(func(event *management.WatchEvent) string {
			return event.Cluster.Id
		}, Equal("test-cluster-id-3"))))

		_, err = client.DeleteCluster(context.Background(), &core.Reference{
			Id: "test-cluster-id-3",
		})
		Expect(err).NotTo(HaveOccurred())

		Eventually(events).Should(Receive(WithTransform(func(event *management.WatchEvent) string {
			return event.Cluster.Id
		}, Equal("test-cluster-id-3"))))

	})

	It("can delete individual clusters", func() {
		_, err := client.DeleteCluster(context.Background(), &core.Reference{
			Id: "test-cluster-id",
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = client.GetCluster(context.Background(), &core.Reference{
			Id: "test-cluster-id",
		})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("failed to get cluster: not found"))
	})

	//#endregion

	//#region Edge Case Tests

	//TODO: Need to add cluster Edge Case Tests

	//#endregion
})
