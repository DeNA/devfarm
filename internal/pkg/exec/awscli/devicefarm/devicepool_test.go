package devicefarm

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/testutil"
	"reflect"
	"testing"
)

func TestDevicePool_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		json     string
		expected *DevicePool
	}{
		{
			json:     devicePoolJSONExample,
			expected: &devicePoolExample,
		},
		{
			json:     "",
			expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Unmarshal([]byte(%q)", c.json), func(t *testing.T) {
			var got DevicePool

			err := json.Unmarshal([]byte(c.json), &got)

			if c.expected != nil {
				if err != nil {
					t.Errorf("got (nil, %v), want (%v, nil)", err, *c.expected)
				} else if !reflect.DeepEqual(got, *c.expected) {
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

func TestDevicePool_MarshalJSON(t *testing.T) {
	cases := []DevicePool{
		devicePoolExample,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("sut := %#v", c), func(t *testing.T) {
			if err := testutil.CheckMarshalAndUnmarshalIsEquivalentToOriginal(&c); err != nil {
				t.Error(err)
			}
		})
	}
}
