package platforms

import "github.com/dena/devfarm/cmd/core/testutil"

func AnyIOSRunner() IOSRunner {
	return FailedIOSRunner()
}

func SuccessfulIOSRunner() IOSRunner {
	return StubIOSRunner(nil)
}

func FailedIOSRunner() IOSRunner {
	return StubIOSRunner(testutil.AnyError)
}

func StubIOSRunner(err error) IOSRunner {
	return func(IOSPlan) error {
		return err
	}
}
