package cmd_test

import "github.com/yngvark.com/clonerepo/pkg/git"

type TestGitter struct{}

func (t TestGitter) Clone(gitUri string, targetCloneDir string) error {
	return nil
}

func (t TestGitter) Pull(gitCloneDir string) error {
	return nil
}

func newTestGitter() git.Gitter {
	return TestGitter{}
}
