package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEchoMultipleArguments(t *testing.T) {
	testEchoOutput(t, []string{"hello", "world"}, "hello world\n")
}

func TestEchoQuotedWhitespace(t *testing.T) {
	testEchoOutput(t, []string{"Go  loops"}, "Go  loops\n")
}

func TestEchoEmptyArgument(t *testing.T) {
	testEchoOutput(t, []string{"", "word"}, " word\n")
}

func TestEchoNoArguments(t *testing.T) {
	testEchoOutput(t, nil, "\n")
}

func testEchoOutput(t *testing.T, args []string, want string) {
	t.Helper()

	loops := []struct {
		name string
		echo func(io.Writer, []string) error
	}{
		{"counted", echoCounted},
		{"condition", echoCondition},
		{"range", echoRange},
		{"break", echoBreak},
	}

	for _, loop := range loops {
		t.Run(loop.name, func(t *testing.T) {
			var output bytes.Buffer
			require.NoError(t, loop.echo(&output, args))
			assert.Equal(t, want, output.String())
		})
	}
}
