package exec

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/logging"
	"testing"
)

func TestExecute(t *testing.T) {
	cases := []struct {
		cmd            string
		args           []string
		expectedStdout string
		expectedStderr string
		expectedErr    bool
	}{
		{
			cmd:            "true",
			args:           []string{},
			expectedStdout: "",
			expectedStderr: "",
			expectedErr:    false,
		},
		{
			cmd:            "false",
			args:           []string{},
			expectedStdout: "",
			expectedStderr: "",
			expectedErr:    true,
		},
		{
			cmd:            "echo",
			args:           []string{"hello"},
			expectedStdout: "hello\n",
			expectedStderr: "",
			expectedErr:    false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("got := execute(logger, %q, %v...)", c.cmd, c.args), func(t *testing.T) {
			logger := logging.NullSeverityLogger()

			got, err := execute(logger, c.cmd, c.args)

			if c.expectedErr {
				if err == nil {
					t.Errorf("got.Error == nil, want error")
					return
				}
			} else {
				if err != nil {
					t.Errorf("got.Error == %v, want nil", err)
					return
				}

				if string(got.Stdout) != c.expectedStdout {
					t.Errorf("got.Stdout == %q, want %q", string(got.Stdout), c.expectedStdout)
				}

				if string(got.Stderr) != c.expectedStderr {
					t.Errorf("got.Stderr == %q, want %q", string(got.Stderr), c.expectedStderr)
				}
			}
		})
	}
}

func TestDryRun(t *testing.T) {
	logger := logging.NullSeverityLogger()

	cmd := "something"
	args := []string{"--arg1", "--arg2"}

	_, err := dryExecute(logger, cmd, args)

	if err != nil {
		t.Errorf("got := dryExecute(logger, %q, %v...); got.Err == %v, but wanted nil", cmd, args, err)
	}
}

func TestExecutedLine(t *testing.T) {
	cases := []struct {
		cmd      string
		args     []string
		expected string
	}{
		{
			cmd:      "command",
			args:     []string{},
			expected: `["command"]`,
		},
		{
			cmd:      "command",
			args:     []string{"--arg1", "--arg2"},
			expected: `["command" "--arg1" "--arg2"]`,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("consoleLikeLine(%q, %q)", c.cmd, c.args), func(t *testing.T) {
			got := consoleLikeLine(c.cmd, c.args)

			if got != c.expected {
				t.Errorf("got %q, want %q", got, c.expected)
			}
		})
	}
}
