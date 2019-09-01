package repo_manager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRepoManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RepoManager Suite")
}
