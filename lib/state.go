package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type State struct {
	ExecutablePath string     `json:"executable_path"`
	Calls          [][]string `json:"calls"`
	VerifyingCall  int        `json:"verifying_call"`
}

func newState(executablePath string, args []string) *State {
	return &State{
		ExecutablePath: executablePath,
		Calls:          [][]string{args},
		VerifyingCall:  0,
	}
}

func readState(path string) (*State, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var state State
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func writeState(path string, state *State) error {
	bytes, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(
		path,
		bytes,
		os.FileMode(0644),
	)
	if err != nil {
		return err
	}
	return nil
}
