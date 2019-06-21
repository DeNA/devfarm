package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyIOSForever() IOSForever {
	return FailedIOSForever()
}

func SuccessfulIOSLauncher() IOSForever {
	return StubIOSForever(nil)
}

func FailedIOSForever() IOSForever {
	return StubIOSForever(testutil.AnyError)
}

func StubIOSForever(err error) IOSForever {
	return func(IOSPlan, IOSForeverBag) error {
		return err
	}
}

func AnyIOSForeverBag() IOSForeverBag {
	return AnyBag()
}
