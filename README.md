<a href="https://github.com/the-gigi/multi-git/actions"><img alt="GitHub Actions status" src="https://github.com/the-gigi/multi-git/workflows/Go/badge.svg"></a> 
<a href="https://github.com/the-gigi/multi-git/releases"><img alt="GitHub Actions status" src="https://github.com/the-gigi/multi-git/workflows/Create%20Release/badge.svg"></a>

# Multi-git
A little Go command-line program that manages a list of git repos and runs git commands on all repos.

This program supports the article: [Let's Go: Command-line Programs with Golang](https://code.tutsplus.com/tutorials/lets-go-command-line-programs-with-golang--cms-26341).

# Command-line Arguments
It accepts two command-line arguments:

* --command : the git command (wrap in double quotes for multi-arguments 
commands)
* --ignore-errors: keeps going through the list of repos even the git command
 failed on some of them

# Environment variables
The list of repos is controlled by two the environment variables:

* MG_ROOT : the path to a root directory that contains all the repos
* MG_REPOS : the names of all managed repos under MG_ROOT
