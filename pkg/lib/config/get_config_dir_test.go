package config_test

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yngvark.com/clonerepo/pkg/lib/config"
)

const (
	userHomeDir   = "/home/bob"
	xdgConfigHome = userHomeDir + "/some-dir"
)

func TestConfigDir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                   string
		fs                     afero.Fs
		osLookupEnvFunc        config.OsLookupEnvFunc
		expectedConfigFilePath string
	}{
		{
			name: "Should return path of XDG_CONFIG_HOME when it's set",
			osLookupEnvFunc: func(key string) (string, bool) {
				if key == "XDG_CONFIG_HOME" {
					return xdgConfigHome, true
				}

				return "", false
			},
			expectedConfigFilePath: path.Join(
				xdgConfigHome, config.Dir, config.FileNameWhenInConfigFolder),
		},
		{
			name: "Should return config in $HOME/.config when it exists",
			fs: func() afero.Fs {
				fs := afero.NewMemMapFs()

				err := fs.MkdirAll(path.Join(userHomeDir, ".config"), os.ModeDir)

				require.NoError(t, err)

				return fs
			}(),
			expectedConfigFilePath: path.Join(
				userHomeDir, ".config", config.Dir, config.FileNameWhenInConfigFolder),
		},
		{
			name:                   "Should return config in $HOME in all other cases",
			expectedConfigFilePath: path.Join(userHomeDir, config.FileNameWhenHomeDir),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			osOpts := config.OsOpts{
				UserHomeDir: func() (string, error) {
					return "/home/bob", nil
				},
			}

			var fs afero.Fs
			if tc.fs == nil {
				fs = afero.NewMemMapFs()
			} else {
				fs = tc.fs
			}

			if tc.osLookupEnvFunc == nil {
				osOpts.LookupEnv = func(key string) (string, bool) {
					return "", false
				}
			} else {
				osOpts.LookupEnv = tc.osLookupEnvFunc
			}

			configFilePath, err := config.GetConfigFilePath(fs, osOpts)
			require.NoError(t, err)

			t.Log("Expected config file path:", tc.expectedConfigFilePath)
			assert.Equal(t, tc.expectedConfigFilePath, configFilePath)
		})
	}
}
