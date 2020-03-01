package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "shellmock",
	Short: "A shell mocking utility to easily create fake executables and assert on executable calls",
}

func init() {
	rootCmd.AddCommand(setupCommand)
	rootCmd.AddCommand(callCommand)
	rootCmd.AddCommand(verifyCommand)
	rootCmd.AddCommand(verifyNoInteractionCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
