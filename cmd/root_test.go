package cmd_test

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yngvark.com/clonerepo/pkg/lib/config"

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
		preRun      func(t *testing.T, cmdOpts cmd.Opts)
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
			name: "Should clone to root tmp directory if temporaryGitDir is unset",
			args: []string{"-t", "git@github.com:some-org/some-repo.git"},
		},
		{
			name: "Should clone to temporary directory if temporaryGitDir is set",
			args: []string{"-t", "git@github.com:some-org/some-repo.git"},
			preRun: func(t *testing.T, cmdOpts cmd.Opts) {
				fs := cmdOpts.OsOpts.Fs
				configFile := getConfigFilename()
				file, err := fs.OpenFile(configFile, os.O_APPEND, 0o600)
				require.NoError(t, err)

				_, err = file.WriteString(fmt.Sprintf("%s: /home/bob/tmp", cmd.TemporaryGitDirKey))

				err = file.Close()
				require.NoError(t, err)
			},
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

			var stdout = &bytes.Buffer{}
			var stderr = &bytes.Buffer{}

			logger := log.New(stdout)

			cmdOpts := buildCmdOpts(stdout, stderr, logger)

			configFilename := createConfigFile(t, cmdOpts.OsOpts.Fs)

			command := cmd.BuildRootCommand(cmdOpts)
			command.SetArgs(tc.args)

			if tc.preRun != nil {
				tc.preRun(t, cmdOpts)
			}

			printConfigFile(t, cmdOpts, configFilename)

			// When
			err = command.Execute()

			printOutput(t, stdout, stderr)

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

func printConfigFile(t *testing.T, cmdOpts cmd.Opts, configFilename string) {
	t.Log("CONFIG FILE:")
	t.Log("-------------------------------------------------")

	configFile, err := cmdOpts.OsOpts.Fs.Open(configFilename)
	require.NoError(t, err)

	scanner := bufio.NewScanner(configFile)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text()) // token in unicode-char
	}
}

func printOutput(t *testing.T, stdout *bytes.Buffer, stderr *bytes.Buffer) {
	t.Log("PROGRAM OUTPUT:")
	t.Log("-------------------------------------------------")
	t.Log("stdout:")
	t.Log(stdout.String())
	t.Log("stderr:")
	t.Log(stderr.String())
	t.Log("-------------------------------------------------")
}

func buildCmdOpts(stdout *bytes.Buffer, stderr *bytes.Buffer, logger *logrus.Logger) cmd.Opts {
	return cmd.Opts{
		Out:    stdout,
		Err:    stderr,
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
}

func createConfigFile(t *testing.T, fs afero.Fs) string {
	t.Helper()

	err := fs.MkdirAll(fmt.Sprintf("/home/bob/.config/%s", config.Dir), 0o700)
	require.NoError(t, err)

	configFile, err := fs.Create(getConfigFilename())
	require.NoError(t, err)

	err = fs.MkdirAll("/home/bob/git", 0o777)
	require.NoError(t, err)

	_, err = configFile.WriteString("gitDir: /home/bob/git\n")
	require.NoError(t, err)

	return configFile.Name()
}

func getConfigFilename() string {
	return fmt.Sprintf(
		"/home/bob/.config/%s/%s", config.Dir, config.FileNameWhenInConfigFolder)
}

func doGoldieAssert(t *testing.T, stdout *bytes.Buffer, stderr *bytes.Buffer) {
	t.Helper()

	goldie := goldiePkg.New(t)
	t.Log(t.Name())

	// Remove apostrophes, so we don't break importing our code as a library
	var goldieFilenameBase string = t.Name()
	goldieFilenameBase = strings.ReplaceAll(goldieFilenameBase, "'", "")

	goldieFilenameStdout := goldieFilenameBase + "-stdout"
	goldieFilenameStderr := goldieFilenameBase + "-stderr"

	//goldie.Update(t, goldieFilenameStdout, stdout.Bytes())
	goldie.Assert(t, goldieFilenameStdout, stdout.Bytes())

	//goldie.Update(t, goldieFilenameStderr, stderr.Bytes())
	goldie.Assert(t, goldieFilenameStderr, stderr.Bytes())
}
