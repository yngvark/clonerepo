# Clonerepo

`clonerepo` removes the hazzle of having to use `cd` to the preferred directory when cloning from GitHub.

It clones github repositores into a pre-determined directory structure, and then `cd`s into the cloned directory.

:information_source: Tip: See also  [`newrepo`](https://github.com/yngvark/newrepo.git).  It creates git repositories into a pre-determined directory structure, and then `cd`s intog the directory containing the repository.

**Example**

```sh
/tmp $ . clonerepo https://github.com/yngvark/newrepo.git
Cloning into directory: /home/myself/git/yngvark/newrepo
Cloning into 'newrepo'...
remote: Enumerating objects: 26, done.
remote: Counting objects: 100% (26/26), done.
remote: Compressing objects: 100% (20/20), done.
remote: Total 26 (delta 7), reused 23 (delta 4), pack-reused 0
Receiving objects: 100% (26/26), 9.02 KiB | 9.02 MiB/s, done.
Resolving deltas: 100% (7/7), done.

~/git/newrepo (main) $ 
```

Notice that `clonerepo` changed the current directory (where the parent path is configurable).

## Install

```sh
git clone https://github.com/yngvark/clonerepo.git
cd clonerepo
go install
```

Finalize with the steps below, depending on the shell you use.

### Bash

Add the following to your `.bashrc`:

```sh
source <path to cloned dir>/clonerepo_bash
```

### Fish

```sh
ln -s <path to cloned dir>/clone.fish ~/.config/fish/functions/clone.fish 
```

## Getting started

* Install `clonerepo` as shown above.

* We need to tell `clonerepo` where it should store repositories.

```sh
mkdir -p ~/.config/clonerepo
```

Replace directory `$HOME/git` below with your preferred directory for keeping repositories:

```sh
echo "gitDir: $HOME/git" >> ~/.config/clonerepo/config.yaml
```

* Now, try cloning a directory:

```bash
clone https://github.com/yngvark/clonerepo
```

* Notice how the current directory changed to `$HOME/git/clonerepo` - or whatever you set your `gitDir` to in the configuration above.

## Uninstall

* Remove binary

```sh
rm $GOBIN/clonerepo
```

* Remove the Bash or Fish specific parts added under [Install](#Install)

## Usage

```sh
$ clone -h
usage: clonerepo [-h] [-t] repoUri

git clones a repo URI to the appropriate directory. Tip: use ". clonerepo <args>"
to change directory to cloned directory.

positional arguments:
  repoUri     URI of the repo to clone

optional arguments:
  -h, --help  show this help message and exit
  -t, --temp  Clone the repository in a temporary directory
```
