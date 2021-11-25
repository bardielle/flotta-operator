package images_test

import (
	"context"
	"encoding/base64"
	"github.com/jakub-dzon/k4e-operator/internal/images"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

const (
	testAuthFile = `{
    "auths": {
        "https://index.docker.io/v1/": {
            "auth": "dGVzdC10ZXN0Cg=="
        }
    }
}`
)

var (
	k8sClient          client.Client
	testAuthFileBase64 = make([]byte, base64.StdEncoding.EncodedLen(len(testAuthFile)))
)

var _ = FDescribe("Images", func() {

	var (
		testEnv  *envtest.Environment
		repoAuth *images.RepositoryAuth
	)

	BeforeEach(func() {
		By("bootstrapping test environment")
		testEnv = &envtest.Environment{}
		var err error
		cfg, err := testEnv.Start()
		Expect(err).NotTo(HaveOccurred())
		Expect(cfg).NotTo(BeNil())

		k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
		Expect(err).NotTo(HaveOccurred())

		base64.StdEncoding.Encode(testAuthFileBase64, []byte(testAuthFile))

		repoAuth = images.NewRepositoryAuth(k8sClient)
	})

	AfterEach(func() {
		err := testEnv.Stop()
		Expect(err).NotTo(HaveOccurred())
	})

	It("should get authfile from a secret", func() {
		// given
		data := map[string][]byte{
			".dockerconfigjson": testAuthFileBase64,
		}
		createSecret("default", "test-auth", data)

		// when
		authFile, err := repoAuth.GetAuthFileFromSecret(context.TODO(), "default", "test-auth")

		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(authFile).To(BeEquivalentTo(testAuthFile))
	})
})

func createSecret(namespace, name string, data map[string][]byte) {
	secret := corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: name, Namespace: namespace},
		Data:       data,
	}
	err := k8sClient.Create(context.TODO(), &secret)
	ExpectWithOffset(1, err).ToNot(HaveOccurred())
}
