package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/afero"
)

const (
	Dir                        = "clonerepo"
	FileNameWhenInConfigFolder = "config.yaml"
	FileNameWhenHomeDir        = ".clonerepo.yaml"
)

type OsOpts struct {
	UserHomeDir OsUserHomeDirFunc
	LookupEnv   OsLookupEnvFunc
}

type OsUserHomeDirFunc func() (string, error)

type OsLookupEnvFunc func(string) (string, bool)

func GetConfigFilePath(fs afero.Fs, opts OsOpts) (string, error) {
	home, err := opts.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user's home directory: %w", err)
	}

	xdgConfigHome, ok := opts.LookupEnv("XDG_CONFIG_HOME")
	if ok {
		// Example: /home/bob/some-dir/.config/clonerepo/config.yaml
		return path.Join(xdgConfigHome, Dir, FileNameWhenInConfigFolder), nil
	}

	//
	// Use dir $HOME/.config if it exists, if not use $HOME
	//
	pathHomeConfig := path.Join(home, ".config")
	_, err = fs.Stat(pathHomeConfig)

	switch {
	case err != nil && !os.IsNotExist(err):
		return "", fmt.Errorf("getting user's home directory: %w", err)
	case err != nil && os.IsNotExist(err): // home/.config does not exist
		// Example: /home/bob/.clonerepo.yaml
		return path.Join(home, ".", FileNameWhenHomeDir), nil
	default:
		// Example: /home/bob/.config/clonerepo/config.yaml
		return path.Join(pathHomeConfig, Dir, "config.yaml"), nil
	}
}
