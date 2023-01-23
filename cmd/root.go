package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yngvark.com/clonerepo/pkg/git"

	"github.com/spf13/viper"
	"github.com/yngvark.com/clonerepo/pkg/lib/log"

	"github.com/spf13/afero"
	"github.com/yngvark.com/clonerepo/pkg/lib/config"

	"github.com/yngvark.com/clonerepo/pkg/lib"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo"

	"github.com/spf13/cobra"
)

const cmdShort = "clonerepo clones git repositores into a pre-determined directory structure, and " +
	"then cd-s into the cloned directory."

type Opts struct {
	Out    io.Writer
	Err    io.Writer
	Logger *logrus.Logger
	Gitter clonerepo.Gitter
	OsOpts config.OsOpts
}

func Run() {
	logger := log.New(os.Stdout)

	cmd := BuildRootCommand(Opts{
		Out:    os.Stdout,
		Err:    os.Stderr,
		Logger: logger,
		Gitter: git.New(logger),
		OsOpts: config.OsOpts{
			UserHomeDir: os.UserHomeDir,
			LookupEnv:   os.LookupEnv,
			Fs:          afero.NewOsFs(),
		},
	})

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// nolint:funlen
func BuildRootCommand(opts Opts) *cobra.Command {
	flags := lib.Flags{}

	cmd := &cobra.Command{
		Use:          "clonerepo",
		Short:        cmdShort,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := config.Init(opts.OsOpts, flags.ConfigFile)
			if err != nil {
				return fmt.Errorf("initializing config: %w", err)
			}

			if flags.Verbose {
				opts.Logger.SetLevel(logrus.DebugLevel)
			}

			opts.Logger.Debugln("Using config file:", viper.ConfigFileUsed())

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			gitDir := viper.GetString("gitDir")

			clonerepoOpts := clonerepo.Opts{
				Out:    opts.Out,
				Logger: opts.Logger,
				Gitter: opts.Gitter,
				Flags:  flags,
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
