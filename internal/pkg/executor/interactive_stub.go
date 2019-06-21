package executor

import (
	"context"
	"github.com/dena/devfarm/internal/pkg/testutil"
)

var AnySuccessfulInteractiveExecutor = StubInteractiveExecutor(nil)
var AnyFailedInteractiveExecutor = StubInteractiveExecutor(testutil.AnyError)

func StubInteractiveExecutor(err error) InteractiveExecutor {
	return func(context.Context, InteractiveRequest) error {
		return err
	}
}
