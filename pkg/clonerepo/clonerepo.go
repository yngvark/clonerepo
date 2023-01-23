package clonerepo

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/yngvark.com/clonerepo/pkg/lib"
	"io"
	"path"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"
)

type Opts struct {
	Out    io.Writer
	Logger *logrus.Logger
	Fs     afero.Fs
	Gitter Gitter

	Flags lib.Flags
}

func Run(opts Opts, gitDir string, args []string) error {
	// Validate
	err := validation.Validate(gitDir, validation.Required)
	if err != nil {
		return fmt.Errorf("gitDir validation error: %w", err)
	}

	err = validation.Validate(args[0], validation.Required)
	if err != nil {
		return fmt.Errorf("args[0] validation error: %w", err)
	}

	gitUri := args[0]

	opts.Logger.Debugln("Git dir: " + gitDir)

	// Get org and repo
	org, repo, err := parse_git_uri.GetOrgAndRepoFromGitUri(gitUri)
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	// Get paths
	dirToRunGitCloneIn := path.Join(gitDir, org)
	clonedDir := path.Join(gitDir, org, repo)

	// Git clone or pull

	cloneDirExists, err := afero.DirExists(opts.Fs, clonedDir)
	if err != nil {
		return fmt.Errorf("checking if directory '%s' exists: %w", clonedDir, err)
	}

	if !cloneDirExists {
		// Organization directory might not exist, so we need to create it. For instance, if cloning
		// https://github.com/some-org/hello.git
		// we want it to be cloned into /home/myself/git/someOrg/hello
		// However, someOrg might not have been created yet.
		err = opts.Fs.MkdirAll(dirToRunGitCloneIn, 0755)
		if err != nil {
			return fmt.Errorf("creating directory '%s': %w", dirToRunGitCloneIn, err)
		}

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

	if opts.Flags.CdToOutputDir {
		fmt.Fprintln(opts.Out, clonedDir)
	}

	return nil
}

func gitClone(opts Opts, gitUri string, targetCloneDir string) error {
	if opts.Flags.DryRun {
		opts.Logger.Debugf("Skipping: git clone " + gitUri + " in " + targetCloneDir)

		return nil
	}

	return opts.Gitter.Clone(gitUri, targetCloneDir)
}

func gitPull(opts Opts, gitCloneDir string) error {
	if opts.Flags.DryRun {
		opts.Logger.Debugf("Skipping: git pull in " + gitCloneDir)

		return nil
	}

	return opts.Gitter.Pull(gitCloneDir)
}
