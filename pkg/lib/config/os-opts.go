package config

import "github.com/spf13/afero"

type OsOpts struct {
	UserHomeDir OsUserHomeDirFunc
	LookupEnv   OsLookupEnvFunc
	Fs          afero.Fs
}
