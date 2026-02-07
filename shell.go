package main

import (
	"context"
	"os"
	"os/exec"
)

func RunInShell(command string) error {

	cmd := exec.CommandContext(context.Background(), "sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil

}
