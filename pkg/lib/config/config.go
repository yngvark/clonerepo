package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/spf13/viper"
)

// Init initializes the program's configuration.
//
// Inspired by: https://github.com/spf13/cobra/blob/main/user_guide.md#create-rootcmd
func Init(osOpts OsOpts, cfgFilepath string) error {
	var err error

	if cfgFilepath == "" {
		cfgFilepath, err = GetConfigFilePath(osOpts)
		if err != nil {
			return fmt.Errorf("getting config file path: %w", err)
		}
	}

	err = createMissingParentDirectories(osOpts.Fs, cfgFilepath)
	if err != nil {
		return fmt.Errorf("creating dir for config file '%s': %w", cfgFilepath, err)
	}

	viper.SetFs(osOpts.Fs)

	// We need to set config file explicitly, because viper doesn't handle creating directory that
	// do not exist.
	viper.SetConfigFile(cfgFilepath)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		err2 := viper.WriteConfig()
		if err2 != nil {
			return fmt.Errorf("writing config: %w", err)
		}
	}

	return nil
}

func createMissingParentDirectories(fs afero.Fs, cfgFilepath string) error {
	dir := filepath.Dir(cfgFilepath)

	//nolint:gomnd
	return fs.MkdirAll(dir, 0o700)
}
