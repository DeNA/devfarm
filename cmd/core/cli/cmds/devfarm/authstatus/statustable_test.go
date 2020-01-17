package authstatus

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/testutil"
	"reflect"
	"testing"
)

func TestFormatAuthStatusTable(t *testing.T) {
	var examplePlatform1 platforms.ID = "example1"
	var examplePlatform2 platforms.ID = "example2"
	header := []string{"platform", "auth"}

	cases := []struct {
		table    map[platforms.ID]error
		expected [][]string
	}{
		{
			table: map[platforms.ID]error{
				examplePlatform1: nil,
			},
			expected: [][]string{
				header,
				{"example1", "success"},
			},
		},
		{
			table: map[platforms.ID]error{
				examplePlatform1: testutil.AnyError,
			},
			expected: [][]string{
				header,
				{"example1", testutil.AnyError.Error()},
			},
		},
		{
			table: map[platforms.ID]error{
				examplePlatform1: nil,
				examplePlatform2: nil,
			},
			expected: [][]string{
				header,
				{"example1", "success"},
				{"example2", "success"},
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("FormatAuthStatusTable(%v)", c.table), func(t *testing.T) {
			got := FormatAuthStatusTable(c.table)

			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}
}
