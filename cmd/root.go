package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
	"github.com/yngvark.com/clonerepo/pkg/lib/log"

	"github.com/spf13/afero"
	"github.com/yngvark.com/clonerepo/pkg/lib/config"

	"github.com/yngvark.com/clonerepo/pkg/lib"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo"

	"github.com/spf13/cobra"
)

const cmdShort = "clonerepo clones git repositores into a pre-determined directory structure, and then `cd`s into" +
	" the cloned directory."

func Run() {
	cmd := BuildRootCommand(Opts{
		Out: os.Stdout,
		Err: os.Stderr,
	})

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

type Opts struct {
	Out        io.Writer
	Err        io.Writer
	FileSystem afero.Fs
}

func BuildRootCommand(opts Opts) *cobra.Command {
	flags := lib.Flags{}

	cmd := &cobra.Command{
		Use:          "clonerepo",
		Short:        cmdShort,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			configOpts := config.Opts{
				Fs: afero.NewOsFs(),
			}

			err := config.Init(configOpts, flags.ConfigFile)
			if err != nil {
				return fmt.Errorf("initializing config: %w", err)
			}

			logger := log.New(opts.Out, flags.Debug)
			logger.Debugln("Using config file:", viper.ConfigFileUsed())

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(opts.Out, flags.Debug)

			if len(args) == 0 {
				return cmd.Help()
			}

			return clonerepo.Run(logger, args)
		},
	}

	cmd.SetOut(opts.Out)
	cmd.SetErr(opts.Err)

	cmd.PersistentFlags().StringVar(
		&flags.ConfigFile,
		"config",
		"",
		"config file (default is: If $HOME/.config exists, it will be "+
			"$HOME/.config/clonerepo/config.yaml. If not, it will be $HOME/.clonerepo.yaml)")

	cmd.PersistentFlags().StringVarP(
		&flags.PrintOutputDirFlag,
		"print-output-dir",
		"p",
		"",
		"Use 'sh' to print a cd command to change to the resulting directory, or 'fish' "+
			"to print the resulting directory")

	cmd.PersistentFlags().BoolVarP(
		&flags.Debug,
		"debug",
		"d",
		false,
		"Enables debug output")

	return cmd
}
