package build_executable

import (
	"path"
	"path/filepath"
	"runtime"
)

var (
	_, caller, _, _  = runtime.Caller(0)                          //nolint:gochecknoglobals
	basepathThisFile = filepath.Dir(caller)                       //nolint:gochecknoglobals
	projectRoot      = path.Clean(basepathThisFile + "/../../..") //nolint:gochecknoglobals
)

func ProjectRoot() string {
	return projectRoot
}

func ProjectBuildDir() string {
	return path.Join(ProjectRoot(), "build")
}
