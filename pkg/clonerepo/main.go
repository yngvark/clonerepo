package clonerepo

import (
	"fmt"
	"io"
	"path"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"
)

func Run(out io.Writer, args []string) error {
	gitDir := "todo!" // TODO: Get from config

	org, repo, err := parse_git_uri.GetOrgAndRepoFromGitUri(args[0])
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	gitRepoDir := path.Join(gitDir, org, repo)
	_, _ = fmt.Fprintf(out, "%s\n", gitRepoDir)

	// TODO:
	// - If dir exists, pull it
	// - Else clone it

	return nil
}
