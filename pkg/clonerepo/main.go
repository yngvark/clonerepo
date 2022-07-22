package clonerepo

import (
	"fmt"
	"io"
	"path"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"
	"github.com/yngvark.com/clonerepo/pkg/lib"
)

func Run(flags lib.Flags, out io.Writer, args []string) error {
	gitDir := "todo!"

	org, repo, err := parse_git_uri.GetOrgAndRepoFromGitUri(args[0])
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	gitRepoDir := path.Join(gitDir, org, repo)
	_, _ = fmt.Fprintf(out, "%s\n", gitRepoDir)

	return nil
}
