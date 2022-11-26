package main_test

import (
	main "github.com/activatedio/go-release"
	"github.com/rogpeppe/go-internal/testscript"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(testscript.RunMain(m, map[string]func() int{
		"go-release": main.Main,
	}))
}

func TestGoRelease(t *testing.T) {

	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
