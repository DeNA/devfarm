package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/executor/awscli"
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"reflect"
	"testing"
)

func TestMapProjectsToInstanceGroups(t *testing.T) {
	var anyARN devicefarm.ProjectARN = "arn:devicefarm:ANY_ARN"
	anyTimestamp := awscli.NewTimestamp(0)

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
					anyTimestamp,
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
					anyTimestamp,
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
