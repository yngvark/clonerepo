package clonerepo

import (
	"fmt"
	"io"

	"github.com/yngvark.com/gclone/pkg/clonerepo/parse_git_uri"
	"github.com/yngvark.com/gclone/pkg/lib"
)

//goland:noinspection GoSnakeCaseUsage
const (
	ENV_GCLONE_GIT_DIR = "GCLONE_GIT_DIR"
)

func cloneRepo(flags lib.Flags, out io.Writer, args []string) error {
	org, repoDir, err := parse_git_uri.GetOrgAndRepoFromGitUri(args[0])
	if err != nil {
		return fmt.Errorf("parsing git organization and repository: %w", err)
	}

	_, _ = fmt.Fprintf(out, "Org: %s - Repo: %s", org, repoDir)

	return nil
}

/*
err := godotenv.Load() //nolint:ifshort
if err != nil {
	return fmt.Errorf("loading .env file: %w", err)
}

envGcloneGitDir, ok := os.LookupEnv(ENV_GCLONE_GIT_DIR)
if !ok {
	return fmt.Errorf("missing environment variable: %s", ENV_GCLONE_GIT_DIR)
}
*/
