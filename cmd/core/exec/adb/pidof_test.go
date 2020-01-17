package adb

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestNewPIDGetter(t *testing.T) {
	cases := []struct {
		stdoutBytes   []byte
		adbError      *ExecutorError
		expectedPID   PID
		expectedError bool
	}{
		{
			stdoutBytes:   []byte("1234"),
			adbError:      nil,
			expectedPID:   1234,
			expectedError: false,
		},
		{
			stdoutBytes:   []byte(""),
			adbError:      nil,
			expectedError: true,
		},
		{
			stdoutBytes:   []byte("NOT_NUMBER"),
			adbError:      nil,
			expectedError: true,
		},
		{
			stdoutBytes:   []byte(""),
			adbError:      &ExecutorError{NoSuchCommand: testutil.AnyError},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("adb.Execute(_) == ([]byte(%q), %v)", c.stdoutBytes, c.adbError), func(t *testing.T) {
			execute := StubExecutor(c.stdoutBytes, c.adbError)
			getPID := NewPIDGetter(execute)

			got, err := getPID("ANY-SERIAL-NUMBER", "com.example.package")
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

				if got != c.expectedPID {
					t.Errorf("got (%v, nil), want (%v, nil)", got, c.expectedPID)
					return
				}
			}
		})
	}
}
