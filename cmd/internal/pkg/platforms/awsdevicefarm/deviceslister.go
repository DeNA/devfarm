package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
)

func newDeviceEntryLister(listDevices devicefarm.DeviceLister) platforms.DeviceLister {
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
