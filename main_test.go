package main

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	out := capturer.CaptureStdout(func() {
		main()
	})

	assert.Equal(t, out, "Hello, Gopher!\n")
}
