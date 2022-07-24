package cmd

import (
	"github.com/spf13/afero"
	"github.com/yngvark.com/clonerepo/pkg/lib/config"
	"io"
	"os"

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
			return config.Init(config.Opts{
				Out: opts.Out,
				Fs:  afero.NewOsFs(),
			}, flags.ConfigFile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			return clonerepo.Run(flags, opts.Out, args)
		},
	}

	cmd.SetOut(opts.Out)
	cmd.SetErr(opts.Err)

	cmd.PersistentFlags().StringVar(&flags.ConfigFile, "config", "", "config file (default is: If"+
		" $HOME/.config exists, it will be $HOME/.config/clonerepo/config.yaml. If not, it will be $HOME/.clonerepo.yaml)")

	cmd.PersistentFlags().StringVarP(&flags.PrintOutputDirFlag, "print-output-dir", "p", "",
		"Use 'sh' to print a cd command to change to the resulting directory, or 'fish' to print the resulting directory")

	return cmd
}
