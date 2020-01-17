package devicefarm

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestPlatform_UnmarshalJSON(t *testing.T) {
	ios := PlatformIsIOS
	android := PlatformIsAndroid

	cases := []struct {
		json     string
		expected *Platform
	}{
		{
			json:     "",
			expected: nil,
		},
		{
			json:     `"IOS"`,
			expected: &ios,
		},
		{
			json:     `"ANDROID"`,
			expected: &android,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Unmarshal([]byte(%q))", c.json), func(t *testing.T) {
			var got Platform

			err := json.Unmarshal([]byte(c.json), &got)

			if c.expected != nil {
				if err != nil {
					t.Errorf("got (nil, %v), want (%v, nil)", err, *c.expected)
				} else if got != *c.expected {
					t.Errorf("got (%v, %v), but wanted (%v, nil)", got, err, *c.expected)
				}
			} else {
				if err == nil {
					t.Errorf("got (nil, nil), want (nil, error)")
				}
			}
		})
	}
}

func TestPlatform_MarshalJSON(t *testing.T) {
	cases := []Platform{PlatformIsIOS, PlatformIsAndroid}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v", c), func(t *testing.T) {
			if err := testutil.CheckMarshalAndUnmarshalIsEquivalentToOriginal(&c); err != nil {
				t.Error(err)
			}
		})
	}
}
