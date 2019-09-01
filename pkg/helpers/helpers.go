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
	err = exec.Command("git", "init").Run()
	return
}

func Addfiles(baseDir string, dirName string, commit bool, filenames ... string) (err error) {
	for _, f := range filenames {
		data := []byte("data for" + f)
		err = ioutil.WriteFile(path.Join(baseDir, dirName, f), data, 0777)
		if err != nil {
			return
		}
	}
	return
}

func DeleteDir(baseDir string, dirName string) (err error) {
	err = os.RemoveAll(path.Join(baseDir, dirName))
	return
}
