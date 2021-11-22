package main

import (
	"os"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	// Mock os.Args
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	os.Args = []string{
		t.Name(),
		"-nick", "foo",
		"-room", "bar",
	}

	// Mock os.Exit of util.OsExit
	oldOsExit := util.OsExit
	defer func() {
		util.OsExit = oldOsExit
	}()

	var exitCode int

	util.OsExit = func(code int) {
		exitCode = code
	}

	// Mock os.Stdin
	oldOsStdin := os.Stdin
	defer func() {
		os.Stdin = oldOsStdin
	}()

	input := []byte("/quit")

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	if _, err = w.Write(input); err != nil {
		t.Error(err)
	}

	w.Close()

	// Run main
	out := capturer.CaptureStdout(func() {
		os.Stdin = r

		main()
	})

	// Assertions
	require.Equal(t, 1, exitCode, "")
	assert.Empty(t, out)
}
