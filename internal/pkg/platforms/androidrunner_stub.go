package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyAndroidRunner() AndroidRunner {
	return FailedAndroidRunner()
}

func SuccessfulAndroidRunner() AndroidRunner {
	return StubAndroidRunner(nil)
}

func FailedAndroidRunner() AndroidRunner {
	return StubAndroidRunner(testutil.AnyError)
}

func StubAndroidRunner(err error) AndroidRunner {
	return func(AndroidPlan) error {
		return err
	}
}
