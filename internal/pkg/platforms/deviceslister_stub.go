package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyDeviceOrError() DeviceOrError {
	return DeviceOrError{
		Device: AnyIOSOrAndroidDevice(),
		Error:  testutil.AnyError,
	}
}

func StubDeviceLister(pairs []DeviceOrError, err error) DeviceLister {
	return func() ([]DeviceOrError, error) {
		return pairs, err
	}
}

func AnyDeviceLister() DeviceLister {
	return StubDeviceLister([]DeviceOrError{}, testutil.AnyError)
}
