package repo_manager

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type RepoManager struct {
	repos        []string
	ignoreErrors bool
}

func NewRepoManager(baseDir string, repoNames []string, ignoreErrors bool) (repoManager *RepoManager, err error) {
	_, err = os.Stat(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New(fmt.Sprintf("base dir: '%s' doesn't exist", baseDir))
		}
		return
	}

	if baseDir[len(baseDir)-1] != '/' {
		baseDir += "/"
	}

	if len(repoNames) == 0 {
		err = errors.New("repo list can't be empty")
		return
	}

	repoManager = &RepoManager{
		ignoreErrors: ignoreErrors,
	}
	for _, r := range repoNames {
		path := baseDir + r
		repoManager.repos = append(repoManager.repos, path)
	}

	return
}

func (m *RepoManager) GetRepos() []string {
	return m.repos
}

func (m *RepoManager) Exec(cmd string) (output map[string]string, err error) {
	output = map[string]string{}
	var components []string
	var multiWord []string
	for _, component := range strings.Split(cmd, " ") {
		if strings.HasPrefix(component, "\"") {
			multiWord = append(multiWord, component[1:])
			continue
		}

		if len(multiWord) > 0 {
			if !strings.HasSuffix(component, "\"") {
				multiWord = append(multiWord, component)
				continue
			}

			multiWord = append(multiWord, component[:len(component)-1])
			component = strings.Join(multiWord, " ")
			multiWord = []string{}
		}

		components = append(components, component)
	}

	// Restore working directory after executing the command
	wd, _ := os.Getwd()
	defer os.Chdir(wd)

	var out []byte
	for _, r := range m.repos {
		// Go to the repo's directory
		os.Chdir(r)

		// Execute the command
		out, err = exec.Command("git", components...).CombinedOutput()
		// Store the result
		output[r] = string(out)

		// Bail out if there was an error and NOT ignoring errors
		if err != nil && !m.ignoreErrors {
			return
		}
	}
	return
}
