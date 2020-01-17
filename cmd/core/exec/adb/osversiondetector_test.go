package adb

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestNewOSVersionDetector(t *testing.T) {
	cases := []struct {
		propValue     string
		propError     error
		expected      OSVersion
		expectedError bool
	}{
		{
			propValue:     "9",
			expected:      "9",
			expectedError: false,
		},
		{
			propValue:     "",
			propError:     nil,
			expected:      "",
			expectedError: true,
		},
		{
			propValue:     "",
			propError:     testutil.AnyError,
			expected:      "",
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("getProp(_, _) == (%q, %v)", c.propValue, c.propError), func(t *testing.T) {
			getProp := StubPropGetter(c.propValue, c.propError)
			detect := NewOSVersionDetector(getProp)

			got, err := detect("ANY_SERIAL_NUMBER")

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
		})
	}
}
