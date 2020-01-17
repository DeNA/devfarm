package instruments

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDevices(t *testing.T) {
}

func TestParseDevices(t *testing.T) {
	cases := []struct {
		stdoutBytes   []byte
		expected      []DeviceEntry
		expectedError bool
	}{
		{
			stdoutBytes: []byte(`Known Devices:
example-mac [04AACE4F-5800-458E-B369-21587B62F463]
9905949119 (12.0) [b1eac0c7cc09b77d902cffdaf912178fd5c5526b]`),
			expected: []DeviceEntry{
				{
					DeviceName:  "9905949119",
					OSVersion:   "12.0",
					DeviceUDID:  "b1eac0c7cc09b77d902cffdaf912178fd5c5526b",
					IsSimulator: false,
				},
			},
		},
		{
			stdoutBytes: []byte(`9905949119 (12.0) [b1eac0c7cc09b77d902cffdaf912178fd5c5526b]`),
			expected: []DeviceEntry{
				{
					DeviceName:  "9905949119",
					OSVersion:   "12.0",
					DeviceUDID:  "b1eac0c7cc09b77d902cffdaf912178fd5c5526b",
					IsSimulator: false,
				},
			},
		},
		{
			stdoutBytes: []byte(`iPad (6th generation) (12.2) [7D316D8F-ACDA-4F3D-A128-23C845292D58] (Simulator)`),
			expected: []DeviceEntry{
				{
					DeviceName:  "iPad (6th generation)",
					OSVersion:   "12.2",
					DeviceUDID:  "7D316D8F-ACDA-4F3D-A128-23C845292D58",
					IsSimulator: true,
				},
			},
		},
		{
			stdoutBytes: []byte(`iPhone Xʀ (12.2) [27ABE406-9B5A-4D73-9DB4-6317BF608A81] (Simulator)`),
			expected: []DeviceEntry{
				{
					DeviceName:  "iPhone Xʀ",
					OSVersion:   "12.2",
					DeviceUDID:  "27ABE406-9B5A-4D73-9DB4-6317BF608A81",
					IsSimulator: true,
				},
			},
		},
		{
			stdoutBytes:   []byte(``),
			expected:      []DeviceEntry{},
			expectedError: false,
		},
		{
			stdoutBytes: []byte(`example-mac [04AACE4F-5800-458E-B369-21587B62F463]`),
			expected:    []DeviceEntry{},
		},
		{
			stdoutBytes: []byte(`iPhone Xs Max (12.2) + Apple Watch Series 4 - 44mm (5.2) [940478FB-551B-443C-9B9B-EC15532F00D0] (Simulator)`),
			expected:    []DeviceEntry{},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("parseDevices([]byte(%q))", string(c.stdoutBytes)), func(t *testing.T) {
			got, err := parseRealDevices(c.stdoutBytes)

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
