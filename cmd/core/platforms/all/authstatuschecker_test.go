package all

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/testutil"
	"reflect"
	"testing"
)

func TestCheckAuthStatusOn(t *testing.T) {
	var examplePlatform1 platforms.ID = "example1"
	var examplePlatform2 platforms.ID = "example2"
	var successfulChecker = func() error { return nil }
	var failedChecker = func() error { return testutil.AnyError }

	cases := []struct {
		checkers map[platforms.ID]platforms.AuthStatusChecker
		expected map[platforms.ID]bool
	}{
		{
			checkers: map[platforms.ID]platforms.AuthStatusChecker{
				examplePlatform1: successfulChecker,
			},
			expected: map[platforms.ID]bool{
				examplePlatform1: true,
			},
		},
		{
			checkers: map[platforms.ID]platforms.AuthStatusChecker{
				examplePlatform1: failedChecker,
			},
			expected: map[platforms.ID]bool{
				examplePlatform1: false,
			},
		},
		{
			checkers: map[platforms.ID]platforms.AuthStatusChecker{
				examplePlatform1: successfulChecker,
				examplePlatform2: successfulChecker,
			},
			expected: map[platforms.ID]bool{
				examplePlatform1: true,
				examplePlatform2: true,
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("CheckAuthStatusOn(%#v)", c.checkers), func(t *testing.T) {
			table := make(map[platforms.ID]platforms.Platform)
			for platformID, checker := range c.checkers {
				p := platforms.AnyPlatform()
				p.AuthStatusCheckerFunc = checker
				table[platformID] = p
			}
			ps := Platforms{table: table}

			statusTable := ps.CheckAllAuthStatus()

			got := errorTableToBoolTable(statusTable)
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}
}

func errorTableToBoolTable(errorTable map[platforms.ID]error) map[platforms.ID]bool {
	boolTable := make(map[platforms.ID]bool, len(errorTable))

	for platform, err := range errorTable {
		boolTable[platform] = err == nil
	}

	return boolTable
}
