package keyring_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/atomic"

	"github.com/rancher/opni-monitoring/pkg/keyring"
	"github.com/rancher/opni-monitoring/pkg/pkp"
)

var _ = Describe("Keyring", func() {
	When("creating an empty keyring", func() {
		It("should function correctly", func() {
			By("creating a new keyring")
			kr := keyring.New()
			Expect(kr).NotTo(BeNil())

			By("ensuring functions passed to Try are never called")
			counter := atomic.NewInt32(0)
			kr.Try(func(*keyring.SharedKeys) {
				counter.Inc()
			}, func(*keyring.PKPKey) {
				counter.Inc()
			})
			Expect(counter.Load()).To(Equal(int32(0)))

			By("ensuring ForEach is never called")
			counter.Store(0)
			kr.ForEach(func(key interface{}) {
				counter.Inc()
			})
			Expect(counter.Load()).To(Equal(int32(0)))

			By("ensuring Marshal returns an empty object")
			j, err := kr.Marshal()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(j)).To(Equal("{}"))
		})
	})
	When("creating a keyring with one key", func() {
		It("should function correctly", func() {
			By("creating a new keyring")
			kr := keyring.New(keyring.NewPKPKey([]*pkp.PublicKeyPin{
				{
					Algorithm:   "sha256",
					Fingerprint: []byte("test"),
				},
			}))
			Expect(kr).NotTo(BeNil())

			By("ensuring Try calls the correct function")
			counter := atomic.NewInt32(0)
			kr.Try(func(keys *keyring.SharedKeys) {
				Fail("Try called the wrong function")
			}, func(key *keyring.PKPKey) {
				counter.Inc()
				Expect(key.PinnedKeys[0].Algorithm).To(BeEquivalentTo("sha256"))
				Expect(key.PinnedKeys[0].Fingerprint).To(BeEquivalentTo("test"))
			})
			Expect(counter.Load()).To(Equal(int32(1)))

			By("ensuring ForEach is called once")
			counter.Store(0)
			kr.ForEach(func(key interface{}) {
				counter.Inc()
				Expect(key).To(BeAssignableToTypeOf(&keyring.PKPKey{}))
				Expect(key.(*keyring.PKPKey).PinnedKeys[0].Algorithm).To(BeEquivalentTo("sha256"))
				Expect(key.(*keyring.PKPKey).PinnedKeys[0].Fingerprint).To(BeEquivalentTo("test"))
			})
			Expect(counter.Load()).To(Equal(int32(1)))

			By("ensuring Marshal returns the correct object")
			j, err := kr.Marshal()
			Expect(err).NotTo(HaveOccurred())
			jsonString := `{"pkpKey":{"pinnedKeys":[{"alg":"sha256","fingerprint":"dGVzdA=="}]}}`
			Expect(string(j)).To(Equal(jsonString))

			By("ensuring Unmarshal returns the correct object")
			kr2, err := keyring.Unmarshal([]byte(jsonString))
			Expect(err).NotTo(HaveOccurred())
			Expect(kr2).NotTo(BeNil())
			Expect(kr).To(BeEquivalentTo(kr2))
		})
	})
	When("creating a keyring with multiple keys", func() {
		It("should function correctly", func() {
			By("creating a new keyring")
			kr := keyring.New(keyring.NewPKPKey([]*pkp.PublicKeyPin{
				{
					Algorithm:   "sha256",
					Fingerprint: []byte("test"),
				},
			}), keyring.NewSharedKeys(make([]byte, 64)))
			Expect(kr).NotTo(BeNil())

			By("ensuring Try calls all functions")
			counterA := atomic.NewInt32(0)
			counterB := atomic.NewInt32(0)
			kr.Try(func(keys *keyring.SharedKeys) {
				counterA.Inc()
				Expect(keys.ClientKey).To(HaveLen(64))
				Expect(keys.ServerKey).To(HaveLen(64))
			}, func(key *keyring.PKPKey) {
				counterB.Inc()
				Expect(key.PinnedKeys[0].Algorithm).To(BeEquivalentTo("sha256"))
				Expect(key.PinnedKeys[0].Fingerprint).To(BeEquivalentTo("test"))
			})
			Expect(counterA.Load()).To(Equal(int32(1)))
			Expect(counterB.Load()).To(Equal(int32(1)))

			By("ensuring ForEach is called for each key")
			counter := atomic.NewInt32(0)
			kr.ForEach(func(key interface{}) {
				counter.Inc()
			})
			Expect(counter.Load()).To(Equal(int32(2)))

			By("ensuring Marshal followed by Unmarshal returns the same data")
			data, err := kr.Marshal()
			Expect(err).NotTo(HaveOccurred())
			kr2, err := keyring.Unmarshal(data)
			Expect(err).NotTo(HaveOccurred())
			Expect(kr2).NotTo(BeNil())
			Expect(kr).To(BeEquivalentTo(kr2))
		})
	})
	It("should handle errors", func() {
		Expect(func() {
			keyring.New("not_an_allowed_keytype")
		}).To(PanicWith(keyring.ErrInvalidKeyType))
		kr := keyring.New()
		Expect(func() {
			kr.Try("not_a_function")
		}).To(PanicWith("invalid UseKeyFn"))
		Expect(func() {
			kr.Try(func(a, b string) {})
		}).To(PanicWith("invalid UseKeyFn (requires one parameter)"))
		kr, err := keyring.Unmarshal([]byte("not_json"))
		Expect(kr).To(BeNil())
		Expect(err).To(HaveOccurred())
		Expect(func() {
			keyring.NewSharedKeys([]byte("not_64_bytes"))
		}).To(PanicWith("shared secret must be 64 bytes"))
	})
})
