package e2e_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/the-gigi/multi-git/pkg/helpers"
	"os"
	"strings"
)

const baseDir = "/tmp/multi-git"

var repoList string

var _ = Describe("multi-git e2e tests", func() {
	var err error

	removeAll := func() {
		err = os.RemoveAll(baseDir)
		Ω(err).Should(BeNil())
	}

	BeforeEach(func() {
		removeAll()
		err = CreateDir(baseDir, "", false)
		Ω(err).Should(BeNil())
	})

	AfterSuite(removeAll)

	Context("Tests for empty/undefined environment failure cases", func() {
		It("Should fail with invalid base dir", func() {
			output, err := RunMultiGit("status", false, "/no-such-dir", repoList, false)
			Ω(err).ShouldNot(BeNil())
			suffix := "base dir: '/no-such-dir/' doesn't exist\n"
			Ω(output).Should(HaveSuffix(suffix))

		})
	})

	Context("Tests for success cases", func() {
		It("Should do git init successfully", func() {
			err = CreateDir(baseDir, "dir-1", false)
			Ω(err).Should(BeNil())
			err = CreateDir(baseDir, "dir-2", false)
			Ω(err).Should(BeNil())
			repoList = "dir-1,dir-2"

			output, err := RunMultiGit("init", false, baseDir, repoList, true)
			Ω(err).Should(BeNil())
			count := strings.Count(output, "Initialized empty Git repository")
			Ω(count).Should(Equal(2))
		})

		It("Should do git status successfully for git directories", func() {
			err = CreateDir(baseDir, "dir-1", true)
			Ω(err).Should(BeNil())
			err = CreateDir(baseDir, "dir-2", true)
			Ω(err).Should(BeNil())
			repoList = "dir-1,dir-2"

			output, err := RunMultiGit("status", false, baseDir, repoList, true)
			Ω(err).Should(BeNil())
			count := strings.Count(output, "nothing to commit")
			Ω(count).Should(Equal(2))

			output, err = RunMultiGit("status", false, baseDir, repoList, false)
			Ω(err).Should(BeNil())
			count = strings.Count(output, "nothing to commit")
			Ω(count).Should(Equal(2))

		})

		It("Should create branches successfully", func() {
			err = CreateDir(baseDir, "dir-1", true)
			Ω(err).Should(BeNil())
			err = CreateDir(baseDir, "dir-2", true)
			Ω(err).Should(BeNil())
			repoList = "dir-1,dir-2"

			output, err := RunMultiGit("checkout -b test-branch", false, baseDir, repoList, true)
			Ω(err).Should(BeNil())
			count := strings.Count(output, "Switched to a new branch 'test-branch'")
			Ω(count).Should(Equal(2))

			output, err = RunMultiGit("checkout -b test-branch", false, baseDir, repoList, false)
			Ω(err).Should(BeNil())
			count = strings.Count(output, "Switched to a new branch 'test-branch'")
			Ω(count).Should(Equal(2))
		})
	})

	Context("Tests for non-git directories", func() {
		It("Should fail git status", func() {
			err = CreateDir(baseDir, "dir-1", false)
			Ω(err).Should(BeNil())
			err = CreateDir(baseDir, "dir-2", false)
			Ω(err).Should(BeNil())
			repoList = "dir-1,dir-2"

			output, err := RunMultiGit("status", false, baseDir, repoList, true)
			Ω(err).Should(BeNil())
			Ω(output).Should(ContainSubstring("fatal: not a git repository"))

			output, err = RunMultiGit("status", false, baseDir, repoList, false)
			Ω(err).Should(BeNil())
			Ω(output).Should(ContainSubstring("fatal: not a git repository"))

		})
	})

	Context("Tests for ignoreErrors flag", func() {
		Context("First directory is invalid", func() {
			When("ignoreErrors is true", func() {
				It("git status should succeed for the second directory", func() {
					err = CreateDir(baseDir, "dir-1", false)
					Ω(err).Should(BeNil())
					err = CreateDir(baseDir, "dir-2", true)
					Ω(err).Should(BeNil())
					repoList = "dir-1,dir-2"

					output, err := RunMultiGit("status", true, baseDir, repoList, true)
					Ω(err).Should(BeNil())
					Ω(output).Should(ContainSubstring("[dir-1]: git status\nfatal: not a git repository"))
					Ω(output).Should(ContainSubstring("[dir-2]: git status\nOn branch master"))

					output, err = RunMultiGit("status", true, baseDir, repoList, false)
					Ω(err).Should(BeNil())
					Ω(output).Should(ContainSubstring("[dir-1]: git status\nfatal: not a git repository"))
					Ω(output).Should(ContainSubstring("[dir-2]: git status\nOn branch master"))

				})
			})

			When("ignoreErrors is false", func() {
				It("Should fail on first directory and bail out", func() {
					err = CreateDir(baseDir, "dir-1", false)
					Ω(err).Should(BeNil())
					err = CreateDir(baseDir, "dir-2", true)
					Ω(err).Should(BeNil())
					repoList = "dir-1,dir-2"

					output, err := RunMultiGit("status", false, baseDir, repoList, true)
					Ω(err).Should(BeNil())
					Ω(output).Should(ContainSubstring("[dir-1]: git status\nfatal: not a git repository"))
					Ω(output).ShouldNot(ContainSubstring("[dir-2]"))

					output, err = RunMultiGit("status", false, baseDir, repoList, false)
					Ω(err).Should(BeNil())
					Ω(output).Should(ContainSubstring("[dir-1]: git status\nfatal: not a git repository"))
					Ω(output).ShouldNot(ContainSubstring("[dir-2]"))
				})
			})
		})
	})
})
