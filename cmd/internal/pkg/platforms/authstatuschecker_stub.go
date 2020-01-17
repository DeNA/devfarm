package platforms

import (
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
)

func AnyAuthStatusChecker() AuthStatusChecker {
	return func() error {
		return testutil.AnyError
	}
}
