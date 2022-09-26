package cmd_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/yngvark.com/clonerepo/pkg/lib/config"
	"strings"
	"testing"

	"github.com/yngvark.com/clonerepo/pkg/lib/log"

	"github.com/spf13/afero"
	"github.com/yngvark.com/clonerepo/cmd"

	goldiePkg "github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
)

func TestCloneRepo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		args        []string
		expectError bool
		asserts     func(t *testing.T, opts cmd.Opts)
	}{
		{
			name: "Should show help if there are zero args",
			args: []string{},
		},
		{
			name: "Should clone repository to expected directory",
			args: []string{"git@github.com:some-org/some-repo.git"},
		},
		{
			name:        "Should return error if git URI is invalid",
			args:        []string{"git@github.com:someorg-somerepo.git"}, // Correct is git@github.com:someorg/somerepo.git
			expectError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Given
			var err error

			var stdout, stderr bytes.Buffer

			logger := log.New(&stdout)

			cmdOpts := cmd.Opts{
				Out:    &stdout,
				Err:    &stderr,
				Logger: logger,
				Gitter: newTestGitter(),
				OsOpts: config.OsOpts{
					UserHomeDir: func() (string, error) {
						return "/home/bob", nil
					},
					LookupEnv: func(key string) (string, bool) {
						return "", false
					},
					Fs: afero.NewMemMapFs(),
				},
			}

			createConfigFile(t, cmdOpts.OsOpts.Fs)

			command := cmd.BuildRootCommand(cmdOpts)
			command.SetArgs(tc.args)

			// When
			err = command.Execute()

			t.Log("PROGRAM OUTPUT:")
			t.Log("-------------------------------------------------")
			t.Log("stdout:")
			t.Log(stdout.String())
			t.Log("stderr:")
			t.Log(stderr.String())
			t.Log("-------------------------------------------------")

			// Then
			if tc.asserts != nil {
				tc.asserts(t, cmdOpts)
			}

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err, stderr.String())
			}

			doGoldieAssert(t, stdout, stderr)
		})
	}
}

func createConfigFile(t *testing.T, fs afero.Fs) {
	err := fs.MkdirAll(fmt.Sprintf("/home/bob/.config/%s", config.Dir), 0700)
	require.NoError(t, err)

	configFile, err := fs.Create(fmt.Sprintf(
		"/home/bob/.config/%s/%s", config.Dir, config.FileNameWhenInConfigFolder))
	require.NoError(t, err)

	err = fs.MkdirAll("/home/bob/git", 0o777)
	require.NoError(t, err)

	_, err = configFile.WriteString("gitDir: /home/bob/git")
	require.NoError(t, err)

}

func doGoldieAssert(t *testing.T, stdout bytes.Buffer, stderr bytes.Buffer) {
	t.Helper()

	goldie := goldiePkg.New(t)
	t.Log(t.Name())

	// Remove apostrophes, so we don't break importing our code as a library
	goldieFilenameBase := strings.ReplaceAll(t.Name(), "'", "")

	goldieFilenameStdout := goldieFilenameBase + "-stdout"
	goldieFilenameStderr := goldieFilenameBase + "-stderr"

	//goldie.Update(t, goldieFilenameStdout, stdout.Bytes())
	goldie.Assert(t, goldieFilenameStdout, stdout.Bytes())

	//goldie.Update(t, goldieFilenameStderr, stderr.Bytes())
	goldie.Assert(t, goldieFilenameStderr, stderr.Bytes())
}
