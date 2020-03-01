package lib

import (
	"fmt"
)

func Verify(executablePath string, expectedArgs []string) error {
	stateFilepath := getStateFilepath(executablePath)
	state, err := readState(stateFilepath)
	if err != nil {
		return err
	}

	verifyingCall := state.VerifyingCall
	if verifyingCall >= len(state.Calls) {
		return fmt.Errorf("%s had no more calls to be verified", executablePath)
	}

	actualArgs := state.Calls[verifyingCall]
	i := 0
	for i < len(expectedArgs) && i < len(actualArgs) {
		if expectedArgs[i] != actualArgs[i] {
			return fmt.Errorf(
				"%s, call #%v, argument %v, expected %s, got %s",
				executablePath, verifyingCall, i, expectedArgs[i], actualArgs[i],
			)
		}
		i = i + 1
	}
	if len(expectedArgs) != len(actualArgs) {
		return fmt.Errorf(
			"%s, expected %v calls, got %v calls",
			executablePath, len(expectedArgs), len(actualArgs),
		)
	}

	newState := State{
		ExecutablePath: state.ExecutablePath,
		Calls:          state.Calls,
		VerifyingCall:  verifyingCall + 1,
	}
	err = writeState(stateFilepath, &newState)
	if err != nil {
		return err
	}

	return nil
}

func VerifyNoInteraction(executablePath string) error {
	stateFilepath := getStateFilepath(executablePath)
	state, err := readState(stateFilepath)
	if err != nil {
		return err
	}

	if state.VerifyingCall < len(state.Calls) {
		return fmt.Errorf(
			"%s, still expecting %v more calls",
			executablePath, len(state.Calls)-state.VerifyingCall,
		)
	}

	return nil
}
