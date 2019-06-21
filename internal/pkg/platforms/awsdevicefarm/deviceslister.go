package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

type deviceEntryLister func() ([]platforms.DeviceOrError, error)

func newDeviceEntryLister(listDevices devicefarm.DeviceLister) deviceEntryLister {
	return func() ([]platforms.DeviceOrError, error) {
		devices, err := listDevices()
		if err != nil {
			return nil, err
		}

		entries := make([]platforms.DeviceOrError, len(devices))
		for i, device := range devices {
			entries[i] = platforms.NewDeviceListEntry(iosOrAndroidDeviceFrom(device), nil)
		}

		return entries, nil
	}
}
