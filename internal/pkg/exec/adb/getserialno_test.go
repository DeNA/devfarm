package adb

import (
	"testing"
)

func TestNewSerialNumberGetter(t *testing.T) {
	cases := []struct {
		stdoutBytes   []byte
		expected      SerialNumber
		expectedError bool
	}{
		{
			stdoutBytes:   []byte("emulator-5554\n"),
			expected:      "emulator-5554",
			expectedError: false,
		},
		{
			stdoutBytes:   []byte{},
			expected:      "",
			expectedError: true,
		},
	}

	for _, c := range cases {
		execute := StubExecutor(c.stdoutBytes, nil)
		getSerialNumber := NewSerialNumberGetter(execute)

		got, err := getSerialNumber()

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

			if got != c.expected {
				t.Errorf("got (%v, nil), want (%v, nil)", got, c.expected)
				return
			}
		}
	}
}
