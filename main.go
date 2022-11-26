package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	os.Exit(Main())
}

func Main() int {

	config, configErr := LoadConfig()

	if configErr != nil {
		printError(configErr)
		return 1
	}

	flag.Parse()

	if flag.NArg() < 1 {
		syntax()
		return 1
	}

	var err error

	switch c := flag.Arg(0); c {
	case "verify":
		err = Verify(config)
	case "perform":
		err = Verify(config)
		if err == nil || errors.Is(err, ErrNoCommand{}) {
			err = Perform(config)
		}
	default:
		err = errors.New(fmt.Sprintf("unrecognized command %s", c))
	}

	if err != nil {
		printError(err)
		return 1
	} else {
		return 0
	}
}

func syntax() {
	os.Stderr.WriteString(`usage: go-release [command]

Available commands:

  verify
  perform
`)
}

func printError(err error) {
	os.Stderr.WriteString(fmt.Sprintf("error: %s\n", err.Error()))
}
