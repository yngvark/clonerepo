package config

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/yngvark.com/clonerepo/pkg/lib"
	"io"
	"os"
)

type Opts struct {
	Out io.Writer
	Fs  afero.Fs
}

// Init initializes the program's configuration.
//
// Inspired by: https://github.com/spf13/cobra/blob/main/user_guide.md#create-rootcmd
func Init(initOpts Opts, cfgFile string) error {
	var err error

	opts := OsOpts{
		UserHomeDir: os.UserHomeDir,
		LookupEnv:   os.LookupEnv,
	}

	if cfgFile == "" {
		cfgFile, err = GetConfigFilePath(initOpts.Fs, opts)
		if err != nil {
			return fmt.Errorf("getting config file path: %w", err)
		}
	}

	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(initOpts.Out, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(initOpts.Out, err.Error())
	}

	return nil
}

func config(flags lib.Flags, out io.Writer, args []string) error {
	return nil
}
