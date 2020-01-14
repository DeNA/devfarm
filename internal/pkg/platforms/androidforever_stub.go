package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyAndroidForever() AndroidForever {
	return FailedAndroidForever()
}

func SuccessfulAndroidForever() AndroidForever {
	return StubAndroidForever(nil)
}

func FailedAndroidForever() AndroidForever {
	return StubAndroidForever(testutil.AnyError)
}

func StubAndroidForever(err error) AndroidForever {
	return func(AndroidPlan) error {
		return err
	}
}
