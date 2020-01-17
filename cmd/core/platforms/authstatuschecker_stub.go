package platforms

import (
	"github.com/dena/devfarm/cmd/core/testutil"
)

func AnyAuthStatusChecker() AuthStatusChecker {
	return func() error {
		return testutil.AnyError
	}
}
