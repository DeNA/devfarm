package devicefarm

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestJob_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		json     string
		expected *Job
	}{
		{
			json:     completedJobJSONExample,
			expected: &completedJobExample,
		},
		{
			json:     "",
			expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Unmarshal([]byte(%q)", c.json), func(t *testing.T) {
			var got Job

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

func TestJob_MarshalJSON(t *testing.T) {
	cases := []Job{
		completedJobExample,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v", c), func(t *testing.T) {
			if err := testutil.CheckMarshalAndUnmarshalIsEquivalentToOriginal(&c); err != nil {
				t.Error(err)
			}
		})
	}
}
