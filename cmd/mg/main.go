package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/the-gigi/multi-git/pkg/repo_manager"
)

func main() {
	command := flag.String("command", "", "The git command")
	ignoreErros := flag.Bool(
		"ignore-errors",
		false,
		"Keep running after error if true")
	flag.Parse()

	// Get managed repos from environment variables
	root := os.Getenv("MG_ROOT")
	if root[len(root)-1] != '/' {
		root += "/"
	}

	repoNames := []string{}
	if len(os.Getenv("MG_REPOS")) > 0 {
		repoNames = strings.Split(os.Getenv("MG_REPOS"), ",")
	}

	repoManager, err := repo_manager.NewRepoManager(root, repoNames, *ignoreErros)
	if err != nil {
		log.Fatal(err)
	}

	output, err := repoManager.Exec(*command)
	if err != nil {
		fmt.Printf("command '%s' failed with error ", err)
	}

	for repo, out := range output {
		fmt.Printf("[%s]: git %s\n", path.Base(repo), *command)
		fmt.Println(out)
	}
}
