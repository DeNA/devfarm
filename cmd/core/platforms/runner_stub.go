package platforms

import "github.com/dena/devfarm/cmd/core/testutil"

func AnyRunner() Runner {
	return StubRunner(*NewResults(testutil.AnyError))
}

func StubRunner(results Results) Runner {
	return func([]EitherPlan) (Results, error) {
		return results, results.Err()
	}
}
