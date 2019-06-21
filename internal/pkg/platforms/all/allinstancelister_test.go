package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"github.com/dena/devfarm/internal/pkg/testutil"
	"reflect"
	"testing"
)

func TestListAllInstances(t *testing.T) {
	var examplePlatform1 platforms.ID = "example1"
	var examplePlatform2 platforms.ID = "example2"
	instance1 := platforms.AnyInstance()
	instance1.State = platforms.InstanceStateIsInactive
	instance2 := platforms.AnyInstance()
	instance2.State = platforms.InstanceStateIsActive
	anyError := testutil.AnyError

	cases := []struct {
		platformTable map[platforms.ID]platforms.AllInstanceLister
		expected      map[platforms.ID]InstancesOrError
	}{
		{
			platformTable: map[platforms.ID]platforms.AllInstanceLister{},
			expected:      map[platforms.ID]InstancesOrError{},
		},
		{
			platformTable: map[platforms.ID]platforms.AllInstanceLister{
				examplePlatform1: platforms.StubAllInstanceLister(
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
			platformTable: map[platforms.ID]platforms.AllInstanceLister{
				examplePlatform1: platforms.StubAllInstanceLister(
					[]platforms.InstanceOrError{
						platforms.NewInstanceListEntry(instance1, nil),
					},
					nil,
				),
				examplePlatform2: platforms.StubAllInstanceLister(
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
			platformTable: map[platforms.ID]platforms.AllInstanceLister{
				examplePlatform1: platforms.StubAllInstanceLister(
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
		t.Run(fmt.Sprintf("ListAllInstancesOn(%v)", c.platformTable), func(t *testing.T) {
			bag := platforms.AnyAllInstanceListerBag()
			got := listAllInstancesOn(c.platformTable, bag)

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
