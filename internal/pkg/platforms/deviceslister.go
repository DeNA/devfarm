package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type DevicesListerBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
}

func NewDeviceListEntry(iosOrAndroidDevices EitherDevice, err error) DeviceOrError {
	return DeviceOrError{
		Device: iosOrAndroidDevices,
		Error:  err,
	}
}

func UnspecificErrorDeviceListEntry(err error) DeviceOrError {
	return NewDeviceListEntry(
		EitherDevice{OSName: OSIsUnavailable},
		err,
	)
}

type DeviceOrError struct {
	Device EitherDevice
	Error  error
}

func (entry DeviceOrError) Less(another DeviceOrError) bool {
	if entry.Device != another.Device {
		return entry.Device.Less(another.Device)
	}

	if entry.Error == nil {
		return true
	}

	if another.Error == nil {
		return false
	}

	return entry.Error.Error() < another.Error.Error()
}

type DeviceLister func(bag DevicesListerBag) ([]DeviceOrError, error)
