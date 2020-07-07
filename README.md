[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/the-gigi/multi-git.svg?label=release)](https://github.com/the-gigi/multi-git/releases)
<a href="https://github.com/the-gigi/multi-git/actions"><img alt="GitHub Actions status" src="https://github.com/the-gigi/multi-git/workflows/Go/badge.svg"></a> 

# Multi-git

A little Go command-line program that manages a list of git repos and runs git commands on all repos.

Here is the usage statement
```
$ multi-git -h
Runs git commands over multiple repos.

It expects the git command to run as an argument.

For example:

multi-git status

If you want to specify multiple flags to git surround them with quotes:

multi-git 'status --short'

It also requires the following environment variables defined:
MG_ROOT: root directory of target git repositories
MG_REPOS: list of repository names to operate on

Usage:
  multi-git [flags]

Flags:
      --config string   config file path (default is $HOME/multi-git.toml)
  -h, --help            help for multi-git
      --ignore-errors   will continue executing the command for all repos if ignore-errors is true
                        otherwise it will stop execution when an error occurs
```

# Makefile

multi-git has a Makefile that lets you build, install, test, create docker image and push the docker image to DockerHub. Type `make` to see this help screen:

```
$ make
Available targets:

help            This help screen
build           Build multi-git and inject the git tag and build time to variables in main
install         Install multi-git into /usr/local/bin (avoid standard go install)
unit-tests      Run only the unit tests
e2e-tests       Run only the end to end tests
test            Run all the tests
ginkgo-test     Run all the tests with ginkgo
docker-build    Dockerize multi-git
docker-push     Push multi-git to DockerHub (requires DockerHub account)
```

# Multi-git v0.1

This [initial release](https://github.com/the-gigi/multi-git/releases/tag/v0.1) of multi-git supports the TutsPlus article: [Let's Go: Command-line Programs with Golang](https://code.tutsplus.com/tutorials/lets-go-command-line-programs-with-golang--cms-26341).

Note: there is no Makefile available for multi-git v0.1.

## Command-line Arguments
It accepts two command-line arguments:

* --command : the git command (wrap in double quotes for multi-arguments 
commands)
* --ignore-errors: keeps going through the list of repos even the git command
 failed on some of them

## Environment variables
The list of repos is controlled by two the environment variables:

* MG_ROOT : the path to a root directory that contains all the repos
* MG_REPOS : the names of all managed repos under MG_ROOT
