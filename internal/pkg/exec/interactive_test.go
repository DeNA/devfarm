package exec

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewInteractiveExecutor(t *testing.T) {
	cases := []struct {
		ctx            func() context.Context
		stdin          string
		command        string
		args           []string
		expectedStdout string
		expectedStderr string
		expectedErr    bool
	}{
		{
			ctx:            func() context.Context { return context.Background() },
			stdin:          "hello",
			command:        "cat",
			args:           []string{"-"},
			expectedStdout: "hello",
			expectedStderr: "",
			expectedErr:    false,
		},
		{
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				return ctx
			},
			stdin:          "",
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
			req := NewInteractiveRequest(ioutil.NopCloser(strings.NewReader(c.stdin)), stdout, stderr, c.command, c.args)

			exec := NewInteractiveExecutor(spyLogger, false)
			err := exec(c.ctx(), req)

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
