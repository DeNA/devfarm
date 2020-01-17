package all

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
	"reflect"
	"testing"
)

func TestListInstances(t *testing.T) {
	var examplePlatform1 platforms.ID = "example1"
	var examplePlatform2 platforms.ID = "example2"
	instance1 := platforms.AnyInstance()
	instance1.State = platforms.InstanceStateIsInactive
	instance2 := platforms.AnyInstance()
	instance2.State = platforms.InstanceStateIsActive
	anyError := testutil.AnyError

	cases := []struct {
		table    map[platforms.ID]platforms.InstanceLister
		expected map[platforms.ID]InstancesOrError
	}{
		{
			table:    map[platforms.ID]platforms.InstanceLister{},
			expected: map[platforms.ID]InstancesOrError{},
		},
		{
			table: map[platforms.ID]platforms.InstanceLister{
				examplePlatform1: platforms.StubInstanceLister(
					[]platforms.InstanceOrError{},
					nil,
				),
			},
			expected: map[platforms.ID]InstancesOrError{
				examplePlatform1: {
					entries:       []platforms.InstanceOrError{},
					platformError: nil,
				},
			},
		},
		{
			table: map[platforms.ID]platforms.InstanceLister{
				examplePlatform1: platforms.StubInstanceLister(
					[]platforms.InstanceOrError{
						platforms.NewInstanceListEntry(instance1, nil),
					},
					nil,
				),
				examplePlatform2: platforms.StubInstanceLister(
					[]platforms.InstanceOrError{
						platforms.NewInstanceListEntry(instance2, nil),
					},
					nil,
				),
			},
			expected: map[platforms.ID]InstancesOrError{
				examplePlatform1: {
					entries: []platforms.InstanceOrError{
						platforms.NewInstanceListEntry(instance1, nil),
					},
					platformError: nil,
				},
				examplePlatform2: {
					entries: []platforms.InstanceOrError{
						platforms.NewInstanceListEntry(instance2, nil),
					},
					platformError: nil,
				},
			},
		},
		{
			table: map[platforms.ID]platforms.InstanceLister{
				examplePlatform1: platforms.StubInstanceLister(
					nil,
					anyError,
				),
			},
			expected: map[platforms.ID]InstancesOrError{
				examplePlatform1: {
					entries:       nil,
					platformError: anyError,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ListInstancesOn(%v)", c.table), func(t *testing.T) {
			table := make(map[platforms.ID]platforms.Platform)
			for platformID, instanceLister := range c.table {
				p := platforms.AnyPlatform()
				p.InstanceListerFunc = instanceLister
				table[platformID] = p
			}
			ps := Platforms{table: table}

			got := ps.ListInstances("ANY_GROUP")

			if c.expected != nil {
				if !reflect.DeepEqual(got, c.expected) {
					t.Errorf("got %v, want %v", got, c.expected)
				}
			} else if got != nil {
				t.Errorf("got %v, want nil", got)
			}
		})
	}
}
