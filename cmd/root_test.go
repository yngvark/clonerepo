package cmd_test

import (
	"bytes"
	"strings"
	"testing"

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
		asserts     func(t *testing.T, opts testOpts)
	}{
		{
			name: "Should show help if there are zero args",
			args: []string{},
		},
		{
			name:        "Should return error if git URI is invalid",
			args:        []string{"git@github.com:someorg-somerepo.git"}, // Correct is git@github.com:someorg/somerepo.git
			expectError: true,
		},
		{
			name:        "Should return error if config is missing",
			args:        []string{"git@github.com:someorg/somerepo.git"},
			expectError: true,
		},
		//{
		//	name: "Should write initial configuration if it's missing",
		//	args: []string{"--verbose", "git@github.com:some-org/some-repo.git"},
		//	asserts: func(t *testing.T, opts testOpts) {
		//		t.Helper()
		//
		//		configFilename := "TODO"
		//		configFile, err := opts.cmdOpts.FileSystem.Open(configFilename)
		//		require.NoError(t, err)
		//
		//		configFileContents, err := ioutil.ReadAll(configFile)
		//		require.NoError(t, err)
		//
		//		goldie := goldiePkg.New(t)
		//		goldie.Assert(t, t.Name(), configFileContents)
		//	},
		//},
		{
			name: "Should clone repository to expected directory",
			args: []string{"git@github.com:some-org/some-repo.git"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			// Given
			t.Parallel()
			var err error
			opts := testOpts{}

			/*
				storage, err := store.NewTemporaryStorage()
				assert.NoError(t, err)

				//c.Env = []string{
				//	fmt.Sprintf("%s=%s", clonerepo.ENV_GCLONE_GIT_DIR, storage.BasePath),
				//	//"INTERNAL__CLONE_TEST_REPO=true",
				//	fmt.Sprintf("PATH=%s:%s", build_executable.ProjectBuildDir(), os.Getenv("PATH")),
				//}
			*/

			var stdout, stderr bytes.Buffer

			cmdOpts := cmd.Opts{
				Out:        &stdout,
				Err:        &stderr,
				FileSystem: afero.NewMemMapFs(),
			}
			opts.cmdOpts = cmdOpts

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
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err, stderr.String())
			}

			if tc.asserts != nil {
				tc.asserts(t, opts)
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
		// goldie.Update(t, goldieFilenameStdout, stdout.Bytes())
		goldie.Assert(t, goldieFilenameStdout, stdout.Bytes())
	}

	if len(stderr.Bytes()) > 0 {
		// goldie.Update(t, goldieFilenameStderr, stderr.Bytes())
		goldie.Assert(t, goldieFilenameStderr, stderr.Bytes())
	}
}

type testOpts struct {
	cmdOpts cmd.Opts
}
