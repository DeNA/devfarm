package devicefarm

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"github.com/dena/devfarm/internal/pkg/testutil"
	"testing"
)

func TestProjectName_ToInstanceGroupName(t *testing.T) {
	cases := []struct {
		projectName       ProjectName
		expected          platforms.InstanceGroupName
		expectedUnmanaged bool
		expectedError     bool
	}{
		{
			projectName:       "devfarm-example",
			expected:          "example",
			expectedUnmanaged: false,
			expectedError:     false,
		},
		{
			projectName:       "devfarm-",
			expected:          "",
			expectedUnmanaged: false,
			expectedError:     true,
		},
		{
			projectName:       "not-managed",
			expected:          "",
			expectedUnmanaged: true,
			expectedError:     true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v.ToInstanceGroupName()", c.projectName), func(t *testing.T) {
			got, err := c.projectName.ToInstanceGroupName()

			if c.expectedError {
				if err == nil {
					t.Errorf("got (_, nil), want (_, error)")
				} else if c.expectedUnmanaged {
					if err.Unmanaged == nil {
						t.Errorf("got (_, &InstanceGroupNameError{Unmanaged: %v}), want (_, &InstanceGroupNameError{Unmanaged: error})", err.Unspecified)
					}
				}
			} else {
				if got != c.expected {
					t.Errorf("got (%#v, %v), want (%#v, nil)", got, err, c.expected)
				}
			}
		})
	}
}

func TestProject_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		json     string
		expected *Project
	}{
		{
			json:     "",
			expected: nil,
		},
		{
			json:     projectJSONExample,
			expected: &ProjectExample,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Unmarshal([]byte(%q)", c.json), func(t *testing.T) {
			var got Project

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

func TestProject_MarshalJSON(t *testing.T) {
	cases := []Project{
		NewProject(
			"MyProjectName",
			"arn:aws:devicefarm:us-west-2:123456789101:project:5e01a8c7-c861-4c0a-b1d5-12345EXAMPLE",
			awscli.NewTimestamp(1535675814),
		),
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v", c), func(t *testing.T) {
			if err := testutil.CheckMarshalAndUnmarshalIsEquivalentToOriginal(&c); err != nil {
				t.Error(err)
			}
		})
	}
}
