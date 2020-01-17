package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
	"strings"
)

func iosOrAndroidDeviceFrom(device devicefarm.Device) platforms.EitherDevice {
	deviceName := fmt.Sprintf(
		"%s %s",
		strings.ToLower(device.Manufacturer),
		strings.ToLower(device.Model),
	)

	switch device.Platform {
	case devicefarm.PlatformIsIOS:
		return platforms.EitherDevice{
			OSName: platforms.OSIsIOS,
			IOS: platforms.NewIOSDevice(
				platforms.IOSDeviceName(deviceName),
				platforms.IOSVersion(device.OS),
			),
		}

	case devicefarm.PlatformIsAndroid:
		return platforms.EitherDevice{
			OSName: platforms.OSIsAndroid,
			Android: platforms.NewAndroidDevice(
				platforms.AndroidDeviceName(deviceName),
				platforms.AndroidVersion(device.OS),
			),
		}

	default:
		panic(fmt.Sprintf("unsupported platform: %q", device.Platform))
	}
}
