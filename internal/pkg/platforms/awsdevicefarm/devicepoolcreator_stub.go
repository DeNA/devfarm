package awsdevicefarm

import "github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"

func anySuccessfulDevicePoolCreatorIfNotExists() devicePoolCreatorIfNotExists {
	return stubDevicePoolCreatorIfNotExists("arn:aws:devicefarm:ANY_DEVICE_POOL", nil)
}

func stubDevicePoolCreatorIfNotExists(devicePoolARN devicefarm.DevicePoolARN, err error) devicePoolCreatorIfNotExists {
	return func(devicefarm.ProjectARN, devicefarm.DeviceARN) (devicefarm.DevicePoolARN, error) {
		return devicePoolARN, err
	}
}
