package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type DefaultGitter struct {
	logger Logger
}

func (g DefaultGitter) Clone(gitUri string, targetCloneDir string) error {
	cmd := exec.Command("git", "clone", gitUri)
	cmd.Dir = targetCloneDir

	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		g.logError(cmd, stderr)

		return fmt.Errorf("running command: %w", err)
	}

	return nil
}

func (g DefaultGitter) Pull(gitCloneDir string) error {
	g.logger.Debugf("Running git pull in " + gitCloneDir)

	cmd := exec.Command("git", "pull")
	cmd.Dir = gitCloneDir

	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		g.logError(cmd, stderr)

		return fmt.Errorf("running command: %w", err)
	}

	return nil
}

func (g DefaultGitter) logError(cmd *exec.Cmd, stderr *bytes.Buffer) {
	g.logger.Errorf("Error! Command '%s' in directory '%s' failed. "+
		"Stderr:\n---\n%s---\n", cmd.String(), cmd.Dir, stderr.String())
}

func New(logger Logger) DefaultGitter {
	return DefaultGitter{
		logger: logger,
	}
}
