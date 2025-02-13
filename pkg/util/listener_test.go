package util_test

import (
	"context"
	"net"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rancher/opni-monitoring/pkg/util"
)

var _ = Describe("Listener", func() {
	When("the given address uses the tcp or tcp4 scheme", func() {
		It("should return a tcp listener", func() {
			listener, err := util.NewProtocolListener(context.Background(), "tcp://:0")
			Expect(err).NotTo(HaveOccurred())
			Expect(listener).To(BeAssignableToTypeOf(&net.TCPListener{}))
			listener.Close()

			listener, err = util.NewProtocolListener(context.Background(), "tcp4://:0")
			Expect(err).NotTo(HaveOccurred())
			Expect(listener).To(BeAssignableToTypeOf(&net.TCPListener{}))
			listener.Close()
		})
	})
	When("the given address uses the unix scheme", func() {
		It("should return a socket listener", func() {
			listener, err := util.NewProtocolListener(context.Background(),
				"unix:///tmp/opni-monitoring-test-util.sock")
			Expect(err).NotTo(HaveOccurred())
			Expect(listener).To(BeAssignableToTypeOf(&net.UnixListener{}))
			listener.Close()
		})
		It("should create the socket's parent directory if needed", func() {
			os.RemoveAll("/tmp/opni-monitoring-test-util-dir")
			listener, err := util.NewProtocolListener(context.Background(),
				"unix:///tmp/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock")
			Expect(err).NotTo(HaveOccurred())
			Expect(listener).To(BeAssignableToTypeOf(&net.UnixListener{}))
			_, err = os.Stat("/tmp/opni-monitoring-test-util-dir")
			Expect(err).NotTo(HaveOccurred())
			listener.Close()
			Expect(os.Remove("/tmp/opni-monitoring-test-util-dir")).To(Succeed())
		})
		It("should clean up existing sockets before creating a new one", func() {
			By("creating a socket")
			listener, err := util.NewProtocolListener(context.Background(),
				"unix:///tmp/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock")
			Expect(err).NotTo(HaveOccurred())
			Expect(listener).To(BeAssignableToTypeOf(&net.UnixListener{}))

			By("ensuring the socket exists")
			_, err = os.Stat("/tmp/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock")
			Expect(err).NotTo(HaveOccurred())

			By("creating a new socket in its place")
			listener, err = util.NewProtocolListener(context.Background(),
				"unix:///tmp/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock")
			Expect(err).NotTo(HaveOccurred())
			Expect(listener).To(BeAssignableToTypeOf(&net.UnixListener{}))
			listener.Close()

			By("creating a non-empty directory where the socket should be")
			Expect(os.Mkdir("/tmp/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock", 0700)).To(Succeed())
			os.Create("/tmp/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock/foo")

			By("creating a new socket where the non-empty directory is")
			listener, err = util.NewProtocolListener(context.Background(),
				"unix:///tmp/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock")
			Expect(err).To(HaveOccurred())

			os.RemoveAll("/tmp/opni-monitoring-test-util-dir")
		})
		When("the user does not have permissions to create the requested directory", func() {
			It("should return an error", func() {
				listener, err := util.NewProtocolListener(context.Background(),
					"unix:///var/lib/opni-monitoring-test-util-dir/opni-monitoring-test-util.sock")
				Expect(err).To(HaveOccurred())
				Expect(listener).To(BeNil())
			})
		})
	})
	When("the given address uses an unsupported scheme", func() {
		It("should return an error", func() {
			_, err := util.NewProtocolListener(context.Background(), "foo://:0")
			Expect(err).To(HaveOccurred())
		})
	})
	When("an invalid address is given", func() {
		It("should return an error", func() {
			_, err := util.NewProtocolListener(context.Background(), "")
			Expect(err).To(HaveOccurred())

			_, err = util.NewProtocolListener(context.Background(), "tcp://")
			Expect(err).To(HaveOccurred())

			_, err = util.NewProtocolListener(context.Background(), string([]byte{0x7f}))
			Expect(err).To(HaveOccurred())
		})
	})
})
