package exec

import (
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
	"time"
)

func AnySuccessfulWaiter() Waiter {
	return AnyImmediatelyBackWaiter(nil)
}

func AnyFailedWaiter() Waiter {
	return AnyImmediatelyBackWaiter(testutil.AnyError)
}

func AnyImmediatelyBackWaiter(err error) Waiter {
	return func(func() (bool, error), string, time.Duration, time.Duration) error {
		return err
	}
}

func StubWaiter(err error) (*func() (bool, error), Waiter) {
	var capturedCond func() (bool, error)

	wait := func(cond func() (bool, error), _ string, _ time.Duration, _ time.Duration) error {
		capturedCond = cond
		return err
	}

	return &capturedCond, wait
}
