package wait_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGonvenience(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "wait suite")
}
