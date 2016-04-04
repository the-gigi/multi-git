# multi-git
A little GO command-line program that manages a list of git repos and runs git commands on all repos.

It accepts two command-line arguments:

--command : the git command (wrap in double quotes for multi-arguments commands)
--ignore-errors: keeps going through the list of repos even the git command failed on some of them

The list of repos is controlled by two the environment variables:

MG_ROOT : the path to a root directory that contains all the repos
MG_REPOS: the names of all managed repos inder MG_ROOT

