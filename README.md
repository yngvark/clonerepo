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
go install https://github.com/yngvark/clonerepo
```

## Getting started

* Install `clonerepo` as shown above.

* We need to tell `clonerepo` where it should store repositories.

```sh
clonerepo config gitDir=$HOME/git # Replace directory with your preference
```

* Now, try cloning a directory:

Bash/Zsh:

```bash
. clonerepo https://github.com/yngvark/clonerepo
```

Fish (this requires [fish-source](#fish-shell-support):

```fish
fs clonerepo https://github.com/yngvark/clonerepo
```

* Notice how the current directory changed to `$HOME/git` - or whatever you set your `gitDir` to in the configuration above.

## Uninstall

```sh
rm $GOBIN/clonerepo
```

## Usage

```sh
$ clonerepo -h
usage: clonerepo [-h] [-t] repoUri

git clones a repo URI to the appropriate directory. Tip: use ". clonerepo <args>"
to change directory to cloned directory.

positional arguments:
  repoUri     URI of the repo to clone

optional arguments:
  -h, --help  show this help message and exit
  -t, --temp  Clone the repository in a temporary directory
```

### Fish shell support

In Fish shell, `.` and `source` do not work. To support Fish, you can install [fish-source](https://github.com/yngvark/fish-source).

You can then use `fs` as you would `.` or `source`, like this:

```
fs clonerepo https://github.com/yngvark/gclone.git
```

