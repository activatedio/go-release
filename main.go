// Package main is the entry point for the go-release
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

const (
	CommandVerify  = "verify"
	CommandPerform = "perform"
)

var cli struct {
	Increment               string   `help:"version increment level - patch|minor|major - defaults to minor"`
	SkipPush                *bool    `help:"skip git push after commit"`
	SkipCleanWorkspaceCheck *bool    `help:"skip check for a clean workspace"`
	Verify                  struct{} `cmd:"" help:"verify release but do not perform"`
	Perform                 struct{} `cmd:"" help:"perform release"`
}

func main() {
	Main()
}

func Main() {

	ctx := kong.Parse(&cli)

	// Load up the config and then override with cli
	config, configErr := LoadConfig()

	if configErr != nil {
		printError(configErr)
		os.Exit(1)
	}

	var err error

	err = overrideAndValidateConfig(config)

	if err != nil {
		printError(err)
		os.Exit(1)
	}

	switch ctx.Command() {
	case CommandVerify:
		err = Verify(config)
	case CommandPerform:
		err = Verify(config)
		if err == nil || errors.Is(err, ErrNoCommand{}) {
			err = Perform(config)
		}
	default:
		err = errors.New("invalid command specified")
	}

	if err != nil {
		printError(err)
		os.Exit(1)
	}
}

func overrideAndValidateConfig(config *Config) error {

	if cli.SkipPush != nil {
		config.SkipPush = *cli.SkipPush
	}
	if cli.SkipCleanWorkspaceCheck != nil {
		config.SkipCleanWorkspaceCheck = *cli.SkipCleanWorkspaceCheck
	}
	if cli.Increment != "" {
		config.Increment = cli.Increment
	}

	switch config.Increment {
	case IncrementMajor, IncrementMinor, IncrementPatch:
		// All good
	case "":
		config.Increment = IncrementMinor
	default:
		return fmt.Errorf("unrecognized increment level %s", config.Increment)
	}

	return nil
}

func printError(err error) {
	_, _err := fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	mustNoError(_err)
}
