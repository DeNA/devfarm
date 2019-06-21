package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"github.com/dena/devfarm/internal/pkg/testutil"
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
		platformTable map[platforms.ID]platforms.InstanceLister
		expected      map[platforms.ID]InstancesOrError
	}{
		{
			platformTable: map[platforms.ID]platforms.InstanceLister{},
			expected:      map[platforms.ID]InstancesOrError{},
		},
		{
			platformTable: map[platforms.ID]platforms.InstanceLister{
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
			platformTable: map[platforms.ID]platforms.InstanceLister{
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
			platformTable: map[platforms.ID]platforms.InstanceLister{
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
		t.Run(fmt.Sprintf("ListInstancesOn(%v)", c.platformTable), func(t *testing.T) {
			bag := platforms.AnyInstanceListerBag()

			got := listInstancesOn(c.platformTable, "ANY_GROUP", bag)

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
