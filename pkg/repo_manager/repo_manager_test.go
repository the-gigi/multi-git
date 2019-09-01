package repo_manager

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/the-gigi/multi-git/pkg/helpers"
)

const baseDir = "/tmp/test-multi-git"

var repoList = []string{"dir-1"}

var _ = Describe("In-memory link manager tests", func() {
	var err error
	BeforeEach(func() {
		err = CreateDir(baseDir, "dir-1", true)
		Ω(err).Should(BeNil())
	})

	It("Should fail with invalid base dir", func() {
		_, err := NewRepoManager("/no-such-dir", repoList, true)
		Ω(err).ShouldNot(BeNil())
	})

	It("Should fail with empty repo list", func() {
		_, err := NewRepoManager(baseDir, []string{}, true)
		Ω(err).ShouldNot(BeNil())
	})
})
