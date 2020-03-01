package lib

import (
	"os"
)

func Call(executablePath string, args []string) error {
	stateFilepath := getStateFilepath(executablePath)
	if _, err := os.Stat(stateFilepath); err == nil {
		// subsequent calls, append to the file
		state, err := readState(stateFilepath)
		if err != nil {
			return err
		}

		newCalls := append(state.Calls, args)
		newState := State{
			ExecutablePath: state.ExecutablePath,
			Calls:          newCalls,
		}

		err = writeState(stateFilepath, &newState)
		if err != nil {
			return err
		}

		return nil
	} else if os.IsNotExist(err) {
		// first call, create the file
		err = writeState(stateFilepath, newState(executablePath, args))
		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}
