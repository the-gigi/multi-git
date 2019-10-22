package helpers

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

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
