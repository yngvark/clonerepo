# Gclone

Gclone removes the hazzle of having to use `cd` to the preferred directory when cloning and creating repositories.

It consists of two commands:

* `clonerepo` clones git repositores into a pre-determined directory structure, and then `cd`s into the cloned directory.
* `newrepo` creates git repositories into a pre-determined directory structure, and then `cd`s into the directory containing the repository.

**Example: clonerepo**

```sh
/tmp $ . clonerepo https://github.com/yngvark/gclone.git
Cloning into directory: /home/myself/git/yngvark/gclone
Cloning into 'gclone'...
remote: Enumerating objects: 26, done.
remote: Counting objects: 100% (26/26), done.
remote: Compressing objects: 100% (20/20), done.
remote: Total 26 (delta 7), reused 23 (delta 4), pack-reused 0
Receiving objects: 100% (26/26), 9.02 KiB | 9.02 MiB/s, done.
Resolving deltas: 100% (7/7), done.

~/git/gclone (main) $ 
```

Notice that `clonerepo` changed the current directory (where the parent path is configurable).

**Example: newrepo**

```sh
/tmp $ . newrepo my-new-repo
Command: gh repo create --clone my-github-username/my-new-repo --public
Successfully created public repository in directory /home/myself/git/my-github-username/my-new-repo
cd /home/myself/git/my-github-username/my-new-repo
```

Notice that `newrepo` changed the current directory to the new repository's directory (where the parent path is configurable).

## Install

```sh
cd wherever-you-put-your-applications-or-repos
git clone https://github.com/yngvark/gclone.git
mkdir -p ~/.local/bin # Make sure this is in your PATH
ln -s $(pwd)/gclone/clonerepo ~/.local/bin/clonerepo
ln -s $(pwd)/gclone/newrepo ~/.local/bin/newrepo
```

In your .bashrc/.zshrc, or wherever you want your environment variables to live, add:

```sh
export GCLONE_GIT_DIR=/home/myself/git
export GCLONE_GIT_TEMP_DIR="/tmp/git"
export REPONEW_DEFAULT_ORGANIZATION="my-git-username"
```

ToDo: Put this into config file or something instead.

### Uninstall

```sh
cd wherever-you-put-your-applications-or-repos
rm -rf gclone
rm ~/.local/bin/clonerepo
rm ~/.local/bin/newrepo
```

## Usage

### Clone repositories

```sh
$ clonerepo -h
usage: gclone_repo [-h] [-t] repoUri

git clones a repo URI to the appropriate directory. Tip: use ". clonerepo <args>"
to change directory to cloned directory.

positional arguments:
  repoUri     URI of the repo to clone

optional arguments:
  -h, --help  show this help message and exit
  -t, --temp  Clone the repository in a temporary directory
```

### Create repositories

Requirements:
* [gh](https://cli.github.com/)

```sh
$ newrepo -h
usage: gclone_reponew [-h] [-n | --dry-run | --no-dry-run] [-p | --private | --no-private] [-t TEMPLATE] [-d DESCRIPTION] repoId

Creates a new Github repository. Tip: use ". newrepo <args>"
to change directory to cloned directory.

positional arguments:
  repoId                Organization (optinal) and repository name. Example: myorg/myrepo

optional arguments:
  -h, --help            show this help message and exit
  -n, --dry-run, --no-dry-run
                        Don't make any changes
  -p, --private, --no-private
                        Make the new repository private
  -t TEMPLATE, --template TEMPLATE
                        repository for template. For instance 'myorg/mytemplaterepo'
  -d DESCRIPTION, --description DESCRIPTION
                        the description for the repository
```

### Fish shell support

In Fish shell, `.` and `source` don't work. To support Fish, you can install [fs](https://github.com/yngvark/fs).

You can then replace `.` in the above commands with `fs`, for instance

```
fs clonerepo https://github.com/yngvark/gclone.git
```

