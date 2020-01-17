package awsdevicefarm

import (
	"errors"
	"fmt"
	"github.com/dena/devfarm/cmd/core/cli/formatter"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/dena/devfarm/cmd/core/platforms"
	"sort"
	"sync"
)

func newDeviceFinder(listDevices devicefarm.DeviceLister) platforms.DeviceFinder {
	return func(iosOrAndroidDevice platforms.EitherDevice) (bool, error) {
		devices, devicesErr := listDevices()
		if devicesErr != nil {
			return false, devicesErr
		}

		for _, device := range devices {
			if iosOrAndroidDevice == iosOrAndroidDeviceFrom(device) {
				return true, nil
			}
		}

		return false, nil
	}
}

type deviceARNFinderError struct {
	NotFound    error
	Unspecified error
}

func (e deviceARNFinderError) Error() string {
	if e.NotFound != nil {
		return e.NotFound.Error()
	}
	return e.Unspecified.Error()
}

type deviceARNFinder func(platforms.EitherDevice) (devicefarm.DeviceARN, *deviceARNFinderError)

func newDeviceARNFinder(logger logging.SeverityLogger, listDevices devicefarm.DeviceLister) deviceARNFinder {
	return func(iosOrAndroidDevice platforms.EitherDevice) (devicefarm.DeviceARN, *deviceARNFinderError) {
		logger.Info("searching AWS Device Farm device")
		devices, devicesErr := listDevices()
		if devicesErr != nil {
			logger.Error(fmt.Sprintf("failed to list devices: %s", devicesErr.Error()))
			return "", &deviceARNFinderError{Unspecified: devicesErr}
		}

		for _, device := range devices {
			if iosOrAndroidDevice == iosOrAndroidDeviceFrom(device) {
				logger.Info("the AWS Device Farm device found")
				logger.Debug(fmt.Sprintf("device ARN: %q", device.ARN))
				return device.ARN, nil
			}
		}

		message := fmt.Sprintf("no such device: %s", iosOrAndroidDevice.Desc())
		logger.Error(message)
		logger.Info(fmt.Sprintf("available ones are:\n%s", formatAvailableDevices(devices)))
		return "", &deviceARNFinderError{
			NotFound: errors.New(message),
		}
	}
}

func newDeviceARNFinderCached(findDeviceARN deviceARNFinder) deviceARNFinder {
	var mu sync.Mutex
	cache := make(map[platforms.EitherDevice]devicefarm.DeviceARN)

	return func(device platforms.EitherDevice) (devicefarm.DeviceARN, *deviceARNFinderError) {
		mu.Lock()
		defer mu.Unlock()

		if cached, ok := cache[device]; ok {
			return cached, nil
		}

		deviceARN, err := findDeviceARN(device)
		if err != nil {
			return "", err
		}

		cache[device] = deviceARN
		return deviceARN, nil
	}
}

func formatAvailableDevices(devices []devicefarm.Device) string {
	iosOrAndroidDevices := make([]platforms.EitherDevice, len(devices))
	for i, device := range devices {
		iosOrAndroidDevices[i] = iosOrAndroidDeviceFrom(device)
	}

	sort.Slice(iosOrAndroidDevices, func(i, j int) bool {
		return iosOrAndroidDevices[i].Less(iosOrAndroidDevices[j])
	})

	table := make([][]string, 1)
	table[0] = []string{"device", "os"}

	for _, iosOrAndroidDevice := range iosOrAndroidDevices {
		var row []string

		switch iosOrAndroidDevice.OSName {
		case platforms.OSIsIOS:
			iosDevice := iosOrAndroidDevice.IOS
			row = []string{string(iosDevice.DeviceName), fmt.Sprintf("ios %s", iosDevice.OSVersion)}
		case platforms.OSIsAndroid:
			androidDevice := iosOrAndroidDevice.Android
			row = []string{string(androidDevice.DeviceName), fmt.Sprintf("android %s", androidDevice.OSVersion)}
		default:
			continue
		}

		table = append(table, row)
	}

	return formatter.PrettyTSV(table)
}
