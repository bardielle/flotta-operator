package images_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestImages(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Images Suite")
}
