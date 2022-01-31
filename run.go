package main

import (
	"os/exec"
)

func ExecuteCommand(config *Configuration) (string, error) {
	cmd := exec.Command(config.Command, config.Args...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(output), nil
}
