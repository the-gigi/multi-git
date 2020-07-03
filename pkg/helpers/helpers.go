package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func ConfigureGit() (err error) {
	err = exec.Command("git", "config", "--global", "user.name", "the-gigi").Run()
	if err != nil {
		return
	}
	err = exec.Command("git", "config", "--global", "user.email", "the.gigi@gmail.com").Run()
	return
}

func CreateDir(baseDir string, name string, initGit bool) (err error) {
	dirName := path.Join(baseDir, name)
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return
	}

	if !initGit {
		return
	}

	currDir, err := os.Getwd()
	if err != nil {
		return
	}
	defer os.Chdir(currDir)
	os.Chdir(dirName)
	err = exec.Command("git", "init").Run()
	return
}

func AddFiles(baseDir string, dirName string, commit bool, filenames ...string) (err error) {
	dir := path.Join(baseDir, dirName)
	for _, f := range filenames {
		data := []byte("data for" + f)
		err = ioutil.WriteFile(path.Join(dir, f), data, 0777)
		if err != nil {
			return
		}
	}

	if !commit {
		return
	}

	currDir, err := os.Getwd()
	if err != nil {
		return
	}
	defer os.Chdir(currDir)
	os.Chdir(dir)
	err = exec.Command("git", "add", "-A").Run()
	if err != nil {
		return
	}

	err = exec.Command("git", "commit", "-m", "added some files...").Run()
	return
}

func RunMultiGit(command string, ignoreErrors bool, mgRoot string, mgRepos string, useConfigFile bool) (output string, err error) {
	out, err := exec.Command("which", "multi-git").CombinedOutput()
	if err != nil {
		return
	}

	if len(out) == 0 {
		err = errors.New("multi-git is not in the PATH")
		return
	}

	components := []string{command}
	env := os.Environ()
	if useConfigFile {
		configFile := path.Join(mgRoot, "multi-git-test-config.toml")
		data := fmt.Sprintf("root = \"%s\"\nrepos = \"%s\"\nignore-errors = %v\n", mgRoot, mgRepos, ignoreErrors)
		err = ioutil.WriteFile(configFile, []byte(data), 0644)
		if err != nil {
			return
		}
		components = append(components, "--config", configFile)
	} else {
		if ignoreErrors {
			components = append(components, "--ignore-errors")
		}
		env = append(env, "MG_ROOT="+mgRoot, "MG_REPOS="+mgRepos)
	}

	cmd := exec.Command("multi-git", components...)
	cmd.Env = env
	out, err = cmd.CombinedOutput()
	output = string(out)
	return
}
