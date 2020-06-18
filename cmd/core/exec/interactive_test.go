package exec

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/google/go-cmp/cmp"
	"io"
	"strings"
	"testing"
	"time"
)

func TestNewInteractiveExecutor(t *testing.T) {
	timeout := 1000 * time.Millisecond

	cases := []struct {
		ctx            func() context.Context
		stdin          io.Reader
		command        string
		args           []string
		expectedStdout string
		expectedStderr string
		expectedErr    bool
	}{
		{
			ctx: func() context.Context {
				ctx, _ := context.WithTimeout(context.Background(), timeout)
				return ctx
			},
			stdin:          strings.NewReader("hello"),
			command:        "cat",
			args:           []string{"-"},
			expectedStdout: "hello",
			expectedStderr: "",
			expectedErr:    false,
		},
		{
			ctx: func() context.Context {
				ctx, _ := context.WithTimeout(context.Background(), timeout)
				return ctx
			},
			stdin:          nil,
			command:        "cat",
			args:           []string{"-"},
			expectedStdout: "",
			expectedStderr: "",
			expectedErr:    false,
		},
		{
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			stdin:          strings.NewReader(""),
			command:        "sleep",
			args:           []string{"100"},
			expectedStdout: "",
			expectedStderr: "",
			expectedErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("echo %q | %s %q", c.stdin, c.command, c.args), func(t *testing.T) {
			spyLogger := logging.SpySeverityLogger()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			req := NewInteractiveRequest(c.stdin, stdout, stderr, c.command, c.args)

			exec := NewInteractiveExecutor(spyLogger, false)
			err := exec.Execute(c.ctx(), req)

			if stdout.String() != c.expectedStdout {
				t.Error(cmp.Diff(c.expectedStdout, stdout.String()))
			}
			if stderr.String() != c.expectedStderr {
				t.Error(cmp.Diff(c.expectedStderr, stderr.String()))
			}

			if c.expectedErr {
				if err == nil {
					t.Error("got nil, want error")
					return
				}
			} else {
				if err != nil {
					t.Errorf("got %v, want nil", err)
					return
				}
			}
		})
	}
}
