package clonerepo_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/yngvark.com/gclone/pkg/testhelper/build"
	"github.com/yngvark.com/gclone/pkg/testhelper/execute"
	"github.com/yngvark.com/gclone/pkg/testhelper/storage"

	"github.com/stretchr/testify/assert"
)

func TestCloneRepo(t *testing.T) {
	t.Parallel()
	build.Run(t)

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

			var stdout, stderr bytes.Buffer
			command := execute.CloneRepo("git@github.com:yngvark/some-repo.git")

			command.Env = []string{
				"GCLONE_GIT_DIR=" + store.BasePath,
				"INTERNAL__CLONE_TEST_REPO=true",
				fmt.Sprintf("PATH=%s:%s", build.ProjectBuildDir(), os.Getenv("PATH")),
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