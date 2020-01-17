package adb

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
	"reflect"
	"testing"
)

func TestNewExecutor(t *testing.T) {
	cases := []struct {
		findError      error
		stdout         []byte
		stderr         []byte
		execError      error
		expectedStdout []byte
		expectedError  bool
	}{
		{
			findError:      nil,
			stdout:         []byte("STDOUT FROM ADB"),
			stderr:         []byte(""),
			execError:      nil,
			expectedStdout: []byte("STDOUT FROM ADB"),
			expectedError:  false,
		},
		{
			findError:      testutil.AnyError,
			stdout:         []byte(""),
			stderr:         []byte(""),
			execError:      nil,
			expectedStdout: []byte(""),
			expectedError:  true,
		},
		{
			findError:      nil,
			stdout:         []byte(""),
			stderr:         []byte(""),
			execError:      testutil.AnyError,
			expectedStdout: []byte(""),
			expectedError:  true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("find(_) == %v, exec(_) == {stdout: []byte(%q), stderr: []byte(%q), err: %v}", c.findError, c.stdout, c.stderr, c.execError), func(t *testing.T) {
			find := exec.StubFinder(c.findError)
			exec := exec.StubExecutor(c.stdout, c.stderr, c.execError)

			execADB := NewExecutor(find, exec)

			got, err := execADB()

			if c.expectedError {
				if err == nil {
					t.Errorf("got (_, nil), want (_, error)")
					return
				}
			} else {
				if err != nil {
					t.Errorf("got (_, %v), want (_, nil)", err)
					return
				}

				if !reflect.DeepEqual(got, c.expectedStdout) {
					t.Errorf("got ([]byte(%q), nil), want ([]byte(%q), nil)", string(got), string(c.expectedStdout))
					return
				}
			}
		})
	}
}
