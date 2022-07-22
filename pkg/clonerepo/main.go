package clonerepo

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"
	"github.com/yngvark.com/clonerepo/pkg/lib"
)

//goland:noinspection GoSnakeCaseUsage
const (
	ENV_GCLONE_GIT_DIR = "GCLONE_GIT_DIR"
)

func cloneRepo(flags lib.Flags, out io.Writer, args []string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("loading .env file: %w", err)
	}

	envGcloneGitDir, ok := os.LookupEnv(ENV_GCLONE_GIT_DIR)
	if !ok {
		return fmt.Errorf("missing environment variable: %s", ENV_GCLONE_GIT_DIR)
	}

	org, repo, err := parse_git_uri.GetOrgAndRepoFromGitUri(args[0])
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	gitRepoDir := path.Join(envGcloneGitDir, org, repo)
	_, _ = fmt.Fprintf(out, "%s\n", gitRepoDir)

	return nil
}
