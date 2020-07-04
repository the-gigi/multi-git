package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"

	"github.com/the-gigi/multi-git/cmd"
)

var (
	GitTag    string
	Timestamp string
)

func main() {
	if GitTag != "" {
		fmt.Printf("Git tag : %s\nBuilt at: %s\n\n", GitTag, Timestamp)

		v := semver.MustParse(GitTag[1:])
		latest, err := selfupdate.UpdateSelf(v, "the-gigi/multi-git")
		if err != nil {
			log.Fatalf("Binary update failed: %v", err)
			return
		} else {
			fmt.Println("Current version is:", latest.Version)
		}
	}

	cmd.Execute()
}
