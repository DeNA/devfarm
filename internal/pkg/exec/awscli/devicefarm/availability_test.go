package devicefarm

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/testutil"
	"testing"
)

func TestAvailability_UnmarshalJSON(t *testing.T) {
	highlyAvailable := AvailabilityIsHighlyAvailable
	available := AvailabilityIsAvailable
	busy := AvailabilityIsBusy
	temporaryNotAvailable := AvailabilityIsTemporaryNotAvailable

	cases := []struct {
		json     string
		expected *Availability
	}{
		{
			json:     "",
			expected: nil,
		},
		{
			json:     `"HIGHLY_AVAILABLE"`,
			expected: &highlyAvailable,
		},
		{
			json:     `"AVAILABLE"`,
			expected: &available,
		},
		{
			json:     `"BUSY"`,
			expected: &busy,
		},
		{
			json:     `"TEMPORARY_NOT_AVAILABLE"`,
			expected: &temporaryNotAvailable,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Unmarshal([]byte(%q)", c.json), func(t *testing.T) {
			var got Availability

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

func TestAvailability_MarshalJSON(t *testing.T) {
	cases := []Availability{
		AvailabilityIsHighlyAvailable,
		AvailabilityIsAvailable,
		AvailabilityIsBusy,
		AvailabilityIsTemporaryNotAvailable,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("json.Marshal(%#v)", c), func(t *testing.T) {
			if err := testutil.CheckMarshalAndUnmarshalIsEquivalentToOriginal(&c); err != nil {
				t.Error(err)
			}
		})
	}
}
