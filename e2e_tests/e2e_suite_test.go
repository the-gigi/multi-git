package e2e_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestMultiGit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultiGit End to End Test  Suite")
}
