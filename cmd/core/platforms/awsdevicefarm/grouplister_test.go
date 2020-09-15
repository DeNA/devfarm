package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
	"reflect"
	"testing"
)

func TestMapProjectsToInstanceGroups(t *testing.T) {
	var anyARN devicefarm.ProjectARN = "arn:devicefarm:ANY_ARN"

	cases := []struct {
		projects []devicefarm.Project
		expected []platforms.InstanceGroupListEntry
	}{
		{
			projects: []devicefarm.Project{},
			expected: []platforms.InstanceGroupListEntry{},
		},
		{
			projects: []devicefarm.Project{
				devicefarm.NewProject(
					"devfarm-example",
					anyARN,
				),
			},
			expected: []platforms.InstanceGroupListEntry{
				platforms.NewInstanceGroupListEntry(
					platforms.NewInstanceGroup("example"),
					nil,
				),
			},
		},
		{
			projects: []devicefarm.Project{
				devicefarm.NewProject(
					"not-managed",
					anyARN,
				),
			},
			expected: []platforms.InstanceGroupListEntry{},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("mapProjectsToInstanceGroups(%v)", c.projects), func(t *testing.T) {
			got := mapProjectsToInstanceGroups(c.projects)

			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("got %#v, want %#v", got, c.expected)
			}
		})
	}
}
