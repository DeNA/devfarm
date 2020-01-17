package awsdevicefarm

import "github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"

func anySuccessfulDevicePoolCreator() devicePoolCreator {
	return stubDevicePoolCreator("arn:aws:devicefarm:ANY_DEVICE_POOL", nil)
}

func stubDevicePoolCreator(devicePoolARN devicefarm.DevicePoolARN, err error) devicePoolCreator {
	return func(devicefarm.ProjectARN, devicefarm.DeviceARN) (devicefarm.DevicePoolARN, error) {
		return devicePoolARN, err
	}
}
