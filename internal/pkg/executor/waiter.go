package executor

import (
	"fmt"
	"time"
)

type Waiter func(cond func() (bool, error), label string, interval time.Duration, timeout time.Duration) error

func NewWaiter() Waiter {
	return func(cond func() (bool, error), label string, interval time.Duration, timeout time.Duration) error {
		begin := time.Now()

		for {
			if shouldWait, err := cond(); !shouldWait || err != nil {
				return err
			}

			elapsed := time.Since(begin)
			if elapsed > timeout {
				break
			}

			time.Sleep(interval)
		}

		return fmt.Errorf("timeout exceeded: %q %f sec", label, interval.Seconds())
	}
}
