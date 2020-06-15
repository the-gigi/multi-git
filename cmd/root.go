package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"

	"github.com/the-gigi/multi-git/pkg/repo_manager"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var configFilename string

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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	// Get managed repos from environment variables
	root := viper.GetString("root")
	if root[len(root)-1] != '/' {
		root += "/"
	}

	repoNames := []string{}
	repoNames = strings.Split(viper.GetString("repos"), ",")

	repoManager, err := repo_manager.NewRepoManager(root, repoNames, viper.GetBool("ignore-errors"))
	check(err)

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
	cobra.OnInitialize(initConfig)

	// Find home directory.
	home, err := homedir.Dir()
	check(err)

	defaultConfigFilename := path.Join(home, ".config/multi-git.toml")
	rootCmd.Flags().StringVar(&configFilename,
		"config",
		defaultConfigFilename,
		"config file path (default is $HOME/multi-git.toml)")
	rootCmd.Flags().Bool(
		"ignore-errors",
		false,
		`will continue executing the command for all repos if ignore-errors is true
                otherwise it will stop execution when an error occurs`)
	err = viper.BindPFlag("ignore-errors", rootCmd.Flags().Lookup("ignore-errors"))
	check(err)
}

func initConfig() {
	_, err := os.Stat(configFilename)
	if os.IsNotExist(err) {
		check(err)
	}

	viper.SetConfigFile(configFilename)
	err = viper.ReadInConfig()
	check(err)

	viper.SetEnvPrefix("MG")
	err = viper.BindEnv("root")
	check(err)

	err = viper.BindEnv("repos")
	check(err)
}

func Execute() {
	err := rootCmd.Execute()
	check(err)
}
