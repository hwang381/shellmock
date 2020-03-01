package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Setup(executablePath string) error {
	if !filepath.IsAbs(executablePath) {
		return fmt.Errorf("%s is not an absoluate path", executablePath)
	}

	stateFilepath := getStateFilepath(executablePath)
	if _, err := os.Stat(stateFilepath); err == nil {
		fmt.Printf("removing old state file %s\n", stateFilepath)
		err = os.Remove(stateFilepath)
		if err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	fmt.Printf("setting up %s\n", executablePath)
	shellFile := fmt.Sprintf(`#!/usr/bin/env sh
shellmock call %s "$@"
`, executablePath)
	return ioutil.WriteFile(
		executablePath,
		[]byte(shellFile),
		os.FileMode(0544),
	)
}
