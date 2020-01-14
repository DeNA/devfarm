package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestHaltAll(t *testing.T) {
	var examplePlatform1 platforms.ID = "example1"
	var examplePlatform2 platforms.ID = "example2"

	cases := []struct {
		table    map[platforms.ID]platforms.Halt
		expected ResultTable
	}{
		{
			table:    map[platforms.ID]platforms.Halt{},
			expected: ResultTable{},
		},
		{
			table: map[platforms.ID]platforms.Halt{
				examplePlatform1: platforms.StubHalt([]error{}),
			},
			expected: ResultTable{
				examplePlatform1: *platforms.NewResults(),
			},
		},
		{
			table: map[platforms.ID]platforms.Halt{
				examplePlatform1: platforms.StubHalt([]error{nil}),
			},
			expected: ResultTable{
				examplePlatform1: *platforms.NewResults(nil),
			},
		},
		{
			table: map[platforms.ID]platforms.Halt{
				examplePlatform1: platforms.StubHalt([]error{Error1{}, Error2{}}),
				examplePlatform2: platforms.StubHalt([]error{}),
			},
			expected: ResultTable{
				examplePlatform1: *platforms.NewResults(Error1{}, Error2{}),
				examplePlatform2: *platforms.NewResults(),
			},
		},
		{
			table: map[platforms.ID]platforms.Halt{
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
		t.Run(fmt.Sprintf("HaltAllOn(%v)", c.table), func(t *testing.T) {
			groupName := platforms.InstanceGroupName("ANY_GROUP")
			table := make(map[platforms.ID]platforms.Platform)
			for platformID, halt := range c.table {
				p := platforms.AnyPlatform()
				p.HaltFunc = halt
				table[platformID] = p
			}
			ps := Platforms{table: table}

			got, _ := ps.HaltAll(groupName)

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
