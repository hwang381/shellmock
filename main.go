package main

import (
	"fmt"
	"hwang381.dev/shellmock/cmd"
	"os"
)

func main() {
	if os.Getenv("SHELLMOCK_ENABLE") != "1" {
		fmt.Println("shellmock could replace arbitrary executables, " +
			"so it explicitly requires the environment variable SHELLMOCK_ENABLE=1 " +
			"so that you don't accidentally run it on your host system")
		os.Exit(1)
	}
	cmd.Execute()
}
