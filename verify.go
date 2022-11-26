package main

import "log"

func Verify(config *Config) error {

	log.Println("starting verify")

	if config.Verify == "" {
		return &ErrNoCommand{Command: "verify"}
	}

	return RunInShell(config.Verify)
}
