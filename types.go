package main

import "fmt"

type ErrNoCommand struct {
	Command string
}

func (e ErrNoCommand) Error() string {
	return fmt.Sprintf("%s command not specified", e.Command)
}
