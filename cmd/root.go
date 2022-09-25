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

const cmdShort = "clonerepo clones git repositores into a pre-determined directory structure, and " +
	"then `cd`s into the cloned directory."

func Run() {
	cmd := BuildRootCommand(Opts{
		Out:        os.Stdout,
		Err:        os.Stderr,
		FileSystem: afero.NewOsFs(),
	})

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func BuildRootCommand(opts Opts) *cobra.Command {
	flags := lib.Flags{}

	cmd := &cobra.Command{
		Use:          "clonerepo",
		Short:        cmdShort,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := config.Init(opts.FileSystem, flags.ConfigFile)
			if err != nil {
				return fmt.Errorf("initializing config: %w", err)
			}

			logger := log.New(opts.Out, flags.Verbose)
			logger.Debugln("Using config file:", viper.ConfigFileUsed())

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(opts.Out, flags.Verbose)

			if len(args) == 0 {
				return cmd.Help()
			}

			gitDir := viper.GetString("gitDir")

			clonerepoOpts := clonerepo.Opts{
				Logger:        logger,
				DryRun:        flags.DryRun,
				CdToOutputDir: flags.CdToOutputDir,
			}

			return clonerepo.Run(clonerepoOpts, gitDir, args)
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

	cmd.PersistentFlags().BoolVarP(
		&flags.CdToOutputDir,
		"cd-to-output-dir",
		"o",
		true,
		"Outputs 'cd <cloned dir>' to stdout, which can be used for sourcing in a shell script.")

	cmd.PersistentFlags().BoolVarP(
		&flags.DryRun,
		"dry-run",
		"d",
		false,
		"Dry run, don't do any changes")

	cmd.PersistentFlags().BoolVarP(
		&flags.Verbose,
		"verbose",
		"v",
		false,
		"Enables verbose output")

	return cmd
}

type Opts struct {
	Out        io.Writer
	Err        io.Writer
	FileSystem afero.Fs
}
