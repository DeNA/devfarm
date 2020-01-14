package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyIOSForever() IOSForever {
	return FailedIOSForever()
}

func SuccessfulIOSForever() IOSForever {
	return StubIOSForever(nil)
}

func FailedIOSForever() IOSForever {
	return StubIOSForever(testutil.AnyError)
}

func StubIOSForever(err error) IOSForever {
	return func(IOSPlan) error {
		return err
	}
}
