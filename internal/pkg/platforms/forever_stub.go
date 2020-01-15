package platforms

import "github.com/dena/devfarm/internal/pkg/testutil"

func AnyForever() Forever {
	return StubForever(*NewResults(testutil.AnyError))
}

func StubForever(results Results) Forever {
	return func([]EitherPlan) (Results, error) {
		return results, results.Err()
	}
}
