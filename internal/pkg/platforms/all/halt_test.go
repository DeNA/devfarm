package all

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"reflect"
	"testing"
)

func TestHaltAll(t *testing.T) {
	var examplePlatform1 platforms.ID = "example1"
	var examplePlatform2 platforms.ID = "example2"

	cases := []struct {
		platformTable map[platforms.ID]platforms.Halt
		expected      ResultTable
	}{
		{
			platformTable: map[platforms.ID]platforms.Halt{},
			expected:      ResultTable{},
		},
		{
			platformTable: map[platforms.ID]platforms.Halt{
				examplePlatform1: platforms.StubHalt([]error{}),
			},
			expected: ResultTable{
				examplePlatform1: *platforms.NewResults(),
			},
		},
		{
			platformTable: map[platforms.ID]platforms.Halt{
				examplePlatform1: platforms.StubHalt([]error{nil}),
			},
			expected: ResultTable{
				examplePlatform1: *platforms.NewResults(nil),
			},
		},
		{
			platformTable: map[platforms.ID]platforms.Halt{
				examplePlatform1: platforms.StubHalt([]error{Error1{}, Error2{}}),
				examplePlatform2: platforms.StubHalt([]error{}),
			},
			expected: ResultTable{
				examplePlatform1: *platforms.NewResults(Error1{}, Error2{}),
				examplePlatform2: *platforms.NewResults(),
			},
		},
		{
			platformTable: map[platforms.ID]platforms.Halt{
				examplePlatform1: platforms.StubHalt([]error{Error1{}}),
				examplePlatform2: platforms.StubHalt([]error{Error2{}}),
			},
			expected: ResultTable{
				examplePlatform1: *platforms.NewResults(Error1{}),
				examplePlatform2: *platforms.NewResults(Error2{}),
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("HaltAllOn(%v)", c.platformTable), func(t *testing.T) {
			bag := platforms.AnyBag()
			groupName := platforms.InstanceGroupName("ANY_GROUP")

			got, _ := haltAllOn(c.platformTable, groupName, bag)

			if !reflect.DeepEqual(got, c.expected) {
				t.Error(cmp.Diff(c.expected, got))
			}
		})
	}
}

type Error1 struct{}

func (Error1) Error() string { return "Error1" }

type Error2 struct{}

func (Error2) Error() string { return "Error2" }
