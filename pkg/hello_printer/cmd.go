package hello_printer

import (
	"fmt"

	"github.com/spf13/cobra"
)

func BuildHelloCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "say-hello",
		Short:        "Say hello",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(Hello())

			return nil
		},
	}

	return cmd
}
