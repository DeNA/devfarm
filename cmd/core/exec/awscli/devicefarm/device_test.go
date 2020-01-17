package devicefarm

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDevice_UnmarshalJSON(t *testing.T) {
	deviceIOSExample := DeviceIOSExample()
	deviceAndroidExample := DeviceAndroidExample()

	cases := []struct {
		json     string
		expected *Device
	}{
		{
			json:     "",
			expected: nil,
		},
		{
			json:     deviceIOSJSONExample,
			expected: &deviceIOSExample,
		},
		{
			json:     deviceAndroidJSONExample,
			expected: &deviceAndroidExample,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Unmarshal([]byte(%q))", c.json), func(t *testing.T) {
			var got Device

			err := json.Unmarshal([]byte(c.json), &got)

			if c.expected != nil {
				if err != nil {
					t.Errorf("got (nil, %v), want (%v, nil)", err, *c.expected)
				} else if got != *c.expected {
					t.Errorf("got (%v, %v), want (%v, nil)", got, err, *c.expected)
				}
			} else {
				if err == nil {
					t.Errorf("got (nil, nil), want (nil, error)")
				}
			}
		})
	}
}
