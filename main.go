package main

import (
	"fmt"

	"github.com/the-gigi/multi-git/cmd"
)

var (
	GitTag    string
	Timestamp string
)

func main() {
	if GitTag != "" {
		fmt.Printf("Git tag : %s\nBuilt at: %s\n\n", GitTag, Timestamp)
	}

	cmd.Execute()
}
