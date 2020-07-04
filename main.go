package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"

	"github.com/the-gigi/multi-git/cmd"
)

const (
	version = "v0.8.18"
)

func main() {
	fmt.Println("version: ", version)
	v := semver.MustParse(version[1:])
	latest, err := selfupdate.UpdateSelf(v, "the-gigi/multi-git")
	if err != nil {
		log.Fatalf("Binary update failed: %v", err)
		return
	} else {
		fmt.Println("Current version is:", latest.Version)
	}

	cmd.Execute()
}
