package cmd

import (
	"github.com/spf13/cobra"
	"hwang381.dev/shellmock/lib"
)

var callCommand = &cobra.Command{
	Use:   "call",
	Short: "(shellmock internal use) Record a fake executable call",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return lib.Call(args[0], args[1:])
	},
}
