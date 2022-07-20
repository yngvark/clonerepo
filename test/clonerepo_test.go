package test_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yngvark.com/gclone/test/storage"
)

func TestSayHello(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		expect string
		// cmd    string
	}{
		{
			name: "Should work",
			// cmd:    "gclone clonerepo git@github.com:yngvark/some-repo.git",
			expect: "Hello",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var err error

			store, err := storage.NewTemporaryStorage()
			assert.NoError(t, err)

			build(t)

			var stdout, stderr bytes.Buffer
			command := exec.Command(
				"../clonerepo", "git@github.com:yngvark/some-repo.git")

			command.Env = []string{
				"GCLONE_GIT_DIR=" + store.BasePath,
				"INTERNAL__CLONE_TEST_REPO=true",
				fmt.Sprintf("PATH=../build:%s", os.Getenv("PATH")),
			}
			command.Stdout = &stdout
			command.Stderr = &stderr

			err = command.Run()
			assert.NoError(t, err)

			t.Log("Test stdout: " + stdout.String())
			t.Log("Test stderr: " + stderr.String())
		})
	}
}

func build(t *testing.T) {
	t.Helper()

	var err error

	t.Log("Building")

	var stdout, stderr bytes.Buffer

	command := exec.Command("make", "-C", "..", "build")
	command.Stdout = &stdout
	command.Stderr = &stderr

	err = command.Run()
	if err != nil {
		t.Log("Build stdout: " + stdout.String())
		t.Log("Build stderr: " + stderr.String())
	}

	require.NoError(t, err)

	t.Log("Building done")
}
