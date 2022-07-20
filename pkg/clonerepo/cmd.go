package clonerepo

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yngvark.com/gclone/pkg/lib"
)

const cmdShort = "`clonerepo` clones git repositores into a pre-determined directory structure, and then `cd`s into" +
	" the cloned directory."

func BuildCommand(flags lib.Flags) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "clonerepo",
		Short:        cmdShort,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cloneRepo(flags, os.Stdout, args)
		},
	}

	return cmd
}
