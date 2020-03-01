package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"hwang381.dev/shellmock/lib"
	"os"
)

var setupCommand = &cobra.Command{
	Use: "setup",
	Short: "Setup a fake executable",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Geteuid() != 0 {
			return errors.New("shellmock setup should be run under root so that it can replace executables")
		}
		return lib.Setup(args[0])
	},
}
