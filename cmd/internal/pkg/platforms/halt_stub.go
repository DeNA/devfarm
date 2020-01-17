package platforms

import "github.com/dena/devfarm/cmd/internal/pkg/testutil"

func AnyHalt() Halt {
	return StubHalt(*NewResults(testutil.AnyError))
}

func StubHalt(results Results) Halt {
	return func(InstanceGroupName) (Results, error) {
		return results, results.Err()
	}
}
