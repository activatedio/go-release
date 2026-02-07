package main_test

import (
	"testing"

	main "github.com/activatedio/go-release"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {

	testscript.Main(m, map[string]func(){
		"go-release": main.Main,
	})
}

func TestGoRelease(t *testing.T) {

	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
