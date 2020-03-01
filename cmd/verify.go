package cmd

import (
	"github.com/spf13/cobra"
	"hwang381.dev/shellmock/lib"
)

var verifyCommand = &cobra.Command{
	Use:   "verify",
	Short: "Verify whether a call has been made. Call this consecutively to verify calls in order.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return lib.Verify(args[0], args[1:])
	},
}

var verifyNoInteractionCommand = &cobra.Command{
	Use:   "verifyNoInteraction",
	Short: "Verify if there are no more interactions",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return lib.VerifyNoInteraction(args[0])
	},
}
