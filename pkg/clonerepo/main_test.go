package clonerepo_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	goldiePkg "github.com/sebdah/goldie/v2"
	"github.com/yngvark.com/clonerepo/pkg/clonerepo"

	"github.com/yngvark.com/clonerepo/pkg/testhelper/build_executable"

	"github.com/yngvark.com/clonerepo/pkg/testhelper/execute"
	"github.com/yngvark.com/clonerepo/pkg/testhelper/store"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestCloneRepo(t *testing.T) {
	t.Parallel()
	build_executable.Run(t)

	testCases := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name: "Should show help if there are zero args",
			args: []string{},
		},
		{
			name:        "Should return error if git URI is invalid",
			args:        []string{"git@github.com-someorg-somerepo.git"},
			expectError: true,
		},
		{
			name: "Should clone repository to expected directory",
			args: []string{"git@github.com:some-org/some-1repo.git"},
		},
	}

	for _, tc := range testCases {
		tc := tc //nolint:varnamelen

		t.Run(tc.name, func(t *testing.T) {
			// Given
			t.Parallel()
			var err error

			storage, err := store.NewTemporaryStorage()
			assert.NoError(t, err)

			var stdout, stderr bytes.Buffer
			command := execute.CloneRepo(tc.args...)

			command.Env = []string{
				fmt.Sprintf("%s=%s", clonerepo.ENV_GCLONE_GIT_DIR, storage.BasePath),
				//"INTERNAL__CLONE_TEST_REPO=true",
				fmt.Sprintf("PATH=%s:%s", build_executable.ProjectBuildDir(), os.Getenv("PATH")),
			}
			command.Stdout = &stdout
			command.Stderr = &stderr

			// When
			err = command.Run()

			t.Log("PROGRAM OUTPUT:")
			t.Log("-------------------------------------------------")
			t.Log(stdout.String())
			t.Log("-------------------------------------------------")

			// Then
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err, stderr.String())
			}

			doGoldieAssert(t, stdout, stderr)
		})
	}
}

func doGoldieAssert(t *testing.T, stdout bytes.Buffer, stderr bytes.Buffer) {
	t.Helper()

	goldie := goldiePkg.New(t)
	t.Log(t.Name())

	// Remove apostrophes, so we don't break importing our code as a library
	goldieFilenameBase := strings.ReplaceAll(t.Name(), "'", "")

	goldieFilenameStdout := goldieFilenameBase + "-stdout"
	goldieFilenameStderr := goldieFilenameBase + "-stderr"

	if len(stdout.Bytes()) > 0 {
		//goldie.Update(t, goldieFilenameStdout, stdout.Bytes())
		goldie.Assert(t, goldieFilenameStdout, stdout.Bytes())
	}
	if len(stderr.Bytes()) > 0 {
		//goldie.Update(t, goldieFilenameStderr, stderr.Bytes())
		goldie.Assert(t, goldieFilenameStderr, stderr.Bytes())
	}
}
