package clonerepo

import (
	"fmt"
	"io"

	"github.com/yngvark.com/gclone/pkg/lib"
)

//goland:noinspection GoSnakeCaseUsage
const (
	ENV_GCLONE_GIT_DIR = "GCLONE_GIT_DIR"
)

func cloneRepo(flags lib.Flags, out io.Writer, args []string) error {
	// org, repoDir = getOrgAndRepoFromGitUri(args[0])
	_, _ = fmt.Fprintln(out, "Hello!")
	fmt.Println("3333")

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
