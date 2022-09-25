package clonerepo

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"
)

type Opts struct {
	Logger             *logrus.Logger
	DryRun             bool
	PrintOutputDirFlag bool
}

func Run(opts Opts, gitDir string, args []string) error {
	gitUri := args[0]

	opts.Logger.Debugln("Git dir: " + gitDir)

	org, _, err := parse_git_uri.GetOrgAndRepoFromGitUri(gitUri)
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	// cloneDir := path.Join(gitDir, org, repo)
	targetCloneDir := path.Join(gitDir, org)

	// nolint: godox
	// TODO:
	// - If dir exists, pull it
	// - Else clone it

	err = gitClone(opts, gitUri, targetCloneDir)
	if err != nil {
		return fmt.Errorf("git cloning: %w", err)
	}

	if opts.PrintOutputDirFlag {
		fmt.Println(targetCloneDir)
	}

	return nil
}

func gitClone(opts Opts, gitUri string, targetCloneDir string) error {
	if opts.DryRun {
		opts.Logger.Infof("Skipping: git clone " + gitUri + " in " + targetCloneDir)

		return nil
	}

	cmd := exec.Command("git", "clone", gitUri)
	cmd.Dir = targetCloneDir

	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		opts.Logger.Errorf("Error! Command '%s' in directory '%s' failed. "+
			"Stderr:\n---\n%s---\n", cmd.String(), cmd.Dir, stderr.String())

		return fmt.Errorf("running command: %w", err)
	}

	return nil
}
