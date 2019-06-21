package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyAuthStatusCheckerBag() AuthStatusCheckerBag {
	return AnyBag()
}

func AnyAuthStatusChecker() AuthStatusChecker {
	return func(AuthStatusCheckerBag) error {
		return testutil.AnyError
	}
}
