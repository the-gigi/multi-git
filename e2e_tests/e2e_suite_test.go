package e2e_tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultiGit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultiGit End to End Test  Suite")
}
