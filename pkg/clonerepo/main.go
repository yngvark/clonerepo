package clonerepo

import (
	"fmt"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"
)

func Run(logger *logrus.Logger, args []string) error {
	// nolint: godox
	gitDir := "todo!" // TODO: Get from config

	org, repo, err := parse_git_uri.GetOrgAndRepoFromGitUri(args[0])
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	gitRepoDir := path.Join(gitDir, org, repo)
	logger.Infof("%s\n", gitRepoDir)

	// nolint: godox
	// TODO:
	// - If dir exists, pull it
	// - Else clone it

	return nil
}
