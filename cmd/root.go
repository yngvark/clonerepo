package cmd

import (
	"os"

	"github.com/yngvark.com/clonerepo/pkg/lib"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo"

	"github.com/spf13/cobra"
)

const cmdShort = "Gclone removes the hazzle of having to use `cd` to the preferred directory when cloning and" +
	" creating repositories."

func Run() {
	cmd := BuildCommand()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func BuildCommand() *cobra.Command {
	flags := lib.Flags{}

	cmd := &cobra.Command{
		Use:          "gclone",
		Short:        cmdShort,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(clonerepo.BuildCommand(flags))

	cmd.PersistentFlags().StringVarP(&flags.PrintOutputDirFlag, "print-output-dir", "p", "",
		"Use 'sh' to print a cd command to change to the resulting directory, or 'fish' to print the resulting directory")

	return cmd
}
