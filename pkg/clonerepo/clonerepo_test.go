package clonerepo_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	goldiePkg "github.com/sebdah/goldie/v2"
	"github.com/yngvark.com/gclone/pkg/clonerepo"

	"github.com/yngvark.com/gclone/pkg/testhelper/build_executable"

	"github.com/yngvark.com/gclone/pkg/testhelper/execute"
	"github.com/yngvark.com/gclone/pkg/testhelper/store"

	"github.com/stretchr/testify/assert"
)

func TestCloneRepo(t *testing.T) {
	t.Parallel()
	build_executable.Run(t)

	testCases := []struct {
		name string
		args []string
		// cmd    string
	}{
		{
			name: "Should show help if there are zero args",
			args: []string{},
		},
		{
			name: "Should clone repository to directory as specified by environment variable",
			args: []string{"git@github.com:yngvark/some-repo.git"},
		},
	}

	for _, tc := range testCases {
		tc := tc //nolint:varnamelen

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var err error

			storage, err := store.NewTemporaryStorage()
			assert.NoError(t, err)

			var stdout, stderr bytes.Buffer
			command := execute.CloneRepo(tc.args...)

			command.Env = []string{
				fmt.Sprintf("%s=%s", clonerepo.ENV_GCLONE_GIT_DIR, storage.BasePath),
				"INTERNAL__CLONE_TEST_REPO=true",
				fmt.Sprintf("PATH=%s:%s", build_executable.ProjectBuildDir(), os.Getenv("PATH")),
			}
			command.Stdout = &stdout
			command.Stderr = &stderr

			err = command.Run()
			assert.NoError(t, err)

			doGoldieAssert(t, stdout)
		})
	}
}

func doGoldieAssert(t *testing.T, buffer bytes.Buffer) {
	t.Helper()

	goldie := goldiePkg.New(t)
	t.Log(t.Name())

	// Remove apostrophes, so we don't break importing our code as a library
	goldieFilename := strings.ReplaceAll(t.Name(), "'", "")

	// goldie.Update(t, goldieFilename, buffer.Bytes())
	goldie.Assert(t, goldieFilename, buffer.Bytes())
}
