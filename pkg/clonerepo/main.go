package clonerepo

import (
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"
)

type Opts struct {
	Logger *logrus.Logger
	Gitter Gitter

	DryRun        bool
	CdToOutputDir bool
}

func Run(opts Opts, gitDir string, args []string) error {
	gitUri := args[0]

	opts.Logger.Debugln("Git dir: " + gitDir)

	org, repo, err := parse_git_uri.GetOrgAndRepoFromGitUri(gitUri)
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	dirToRunGitCloneIn := path.Join(gitDir, org)

	clonedDir := path.Join(gitDir, org, repo)

	cloneDirExists, err := dirExists(clonedDir)
	if err != nil {
		return fmt.Errorf("checking if directory '%s' exists: %w", clonedDir, err)
	}

	if !cloneDirExists {
		err = gitClone(opts, gitUri, dirToRunGitCloneIn)
		if err != nil {
			return fmt.Errorf("running git clone: %w", err)
		}
	} else {
		err = gitPull(opts, clonedDir)
		if err != nil {
			return fmt.Errorf("running git pull: %w", err)
		}
	}

	if opts.CdToOutputDir {
		fmt.Println("cd " + dirToRunGitCloneIn)
	}

	return nil
}

func dirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func gitClone(opts Opts, gitUri string, targetCloneDir string) error {
	if opts.DryRun {
		opts.Logger.Debugf("Skipping: git clone " + gitUri + " in " + targetCloneDir)

		return nil
	}

	return opts.Gitter.Clone(gitUri, targetCloneDir)
}

func gitPull(opts Opts, gitCloneDir string) error {
	if opts.DryRun {
		opts.Logger.Debugf("Skipping: git pull in " + gitCloneDir)

		return nil
	}

	return opts.Gitter.Pull(gitCloneDir)
}
