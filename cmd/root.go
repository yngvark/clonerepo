package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/yngvark.com/gclone/pkg/hello_printer"
)

func Run() {
	cmd := &cobra.Command{
		Use:          "gclone",
		Short:        "My-command does this and that.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(hello_printer.BuildHelloCommand())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
