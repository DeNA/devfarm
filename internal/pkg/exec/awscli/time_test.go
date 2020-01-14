package awscli

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/testutil"
	"testing"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	time1 := NewTimestamp(1535675814)

	cases := []struct {
		json     string
		expected *Timestamp
	}{
		{
			json:     "",
			expected: nil,
		},
		{
			json:     `1535675814`,
			expected: &time1,
		},
		{
			json:     `1535675814.000`,
			expected: &time1,
		},
		{
			// NOTE: https://docs.aws.amazon.com/devicefarm/latest/developerguide/how-to-create-project.html
			json:     `1535675814.414`,
			expected: &time1,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Unmarshal([]byte(%q)", c.json), func(t *testing.T) {
			var got Timestamp

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

func TestTime_MarshalJSON(t *testing.T) {
	time1 := NewTimestamp(1535675814)

	cases := []Timestamp{
		time1,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("json.Marshal(%#v)", c), func(t *testing.T) {
			if err := testutil.CheckMarshalAndUnmarshalIsEquivalentToOriginal(&c); err != nil {
				t.Error(err)
			}
		})
	}
}
