package adb

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseDevices(t *testing.T) {
	cases := []struct {
		stdoutBytes   []byte
		expected      []DeviceEntry
		expectedError bool
	}{
		{
			stdoutBytes:   []byte(``),
			expected:      []DeviceEntry{},
			expectedError: false,
		},
		{
			stdoutBytes: []byte(`List of devices attached
emulator-5554	device
`),
			expected: []DeviceEntry{
				{
					Name:  "emulator-5554",
					State: DeviceIsConnected,
				},
			},
			expectedError: false,
		},
		{
			stdoutBytes:   []byte("something\twent\twrong"),
			expected:      nil,
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("parseDevices([]byte(%q))", string(c.stdoutBytes)), func(t *testing.T) {
			got, err := parseDevices(c.stdoutBytes)

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

				if !reflect.DeepEqual(got, c.expected) {
					t.Errorf("got (%v, nil), want (%v, nil)", got, c.expected)
					return
				}
			}
		})
	}
}
