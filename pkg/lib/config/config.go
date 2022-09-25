package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type Opts struct {
	Out io.Writer
	Fs  afero.Fs
}

// Init initializes the program's configuration.
//
// Inspired by: https://github.com/spf13/cobra/blob/main/user_guide.md#create-rootcmd
func Init(initOpts Opts, cfgFilepath string) error {
	var err error

	opts := OsOpts{
		UserHomeDir: os.UserHomeDir,
		LookupEnv:   os.LookupEnv,
	}

	if cfgFilepath == "" {
		cfgFilepath, err = GetConfigFilePath(initOpts.Fs, opts)
		if err != nil {
			return fmt.Errorf("getting config file path: %w", err)
		}
	}

	err = createMissingParentDirectories(cfgFilepath)
	if err != nil {
		return fmt.Errorf("creating dir for config file '%s': %w", cfgFilepath, err)
	}

	// We need to set config file explicitly, because viper doesn't handle creating directory that
	// do not exist.
	viper.SetConfigFile(cfgFilepath)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(initOpts.Out, "Using config file:", viper.ConfigFileUsed())
	} else {
		err2 := viper.WriteConfig()
		if err2 != nil {
			return fmt.Errorf("!!!!!!!!!!!!! writing config: %w", err)
		}
	}

	return nil
}

func createMissingParentDirectories(cfgFilepath string) error {
	dir := filepath.Dir(cfgFilepath)
	return os.MkdirAll(dir, 0o700)
}
