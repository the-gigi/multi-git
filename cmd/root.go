package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"

	"github.com/the-gigi/multi-git/pkg/repo_manager"
)

var ignoreErrors bool

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   "multi-git",
	Short: "Runs git commands over multiple repos",
	Long: `Runs git commands over multiple repos.
           
It expects the git command to run as an argument.

For example:

multi-git status

If you want to specify multiple flags to git surround them with quotes:

multi-git 'status --short'

It also requires the following environment variables defined:	
MG_ROOT: root directory of target git repositories
MG_REPOS: list of repository names to operate on`,
	Args: cobra.ExactArgs(1),
	Run:  run,
}

func run(cmd *cobra.Command, args []string) {
	// Get managed repos from environment variables
	root := os.Getenv("MG_ROOT")
	if root[len(root)-1] != '/' {
		root += "/"
	}

	repoNames := []string{}
	if len(os.Getenv("MG_REPOS")) > 0 {
		repoNames = strings.Split(os.Getenv("MG_REPOS"), ",")
	}

	repoManager, err := repo_manager.NewRepoManager(root, repoNames, ignoreErrors)
	if err != nil {
		log.Fatal(err)
	}

	output, err := repoManager.Exec(args[0])
	if err != nil {
		fmt.Printf("command '%s' failed with error ", err)
	}

	for repo, out := range output {
		fmt.Printf("[%s]: git %s\n", path.Base(repo), args[0])
		fmt.Println(out)
	}
}

func init() {
	rootCmd.Flags().BoolVar(
		&ignoreErrors,
		"ignore-errors",
		false,
		`will continue executing the command for all repos if ignore-errors is true
                otherwise it will stop execution when an error occurs`)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
