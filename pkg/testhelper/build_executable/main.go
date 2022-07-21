// Package build_executable can build the program.
package build_executable

import (
	"bytes"
	"os/exec"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	buildLock  sync.Mutex //nolint:gochecknoglobals
	appIsBuilt bool       //nolint:gochecknoglobals
)

// Run builds the application into an executable, using "make build".
// We have put this here so we can do end-to-end tests on our program.
func Run(t *testing.T) {
	t.Helper()
	buildLock.Lock()

	if appIsBuilt {
		return
	}

	t.Log("Building program")

	var err error

	var stdout, stderr bytes.Buffer

	command := exec.Command("make", "-C", ProjectRoot(), "build") //nolint:gosec // This should be secure
	command.Stdout = &stdout
	command.Stderr = &stderr

	err = command.Run()
	if err != nil {
		t.Log("Build stdout: " + stdout.String())
		t.Log("Build stderr: " + stderr.String())
	}

	require.NoError(t, err)

	appIsBuilt = true
	buildLock.Unlock()
}
