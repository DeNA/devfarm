package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
)

func anySuccessfulDeviceARNFinder() deviceARNFinder {
	return stubDeviceARNFinder("arn:aws:devicefarm:ANY_DEVICE_ARN", nil)
}

func stubDeviceARNFinder(deviceARN devicefarm.DeviceARN, err *deviceARNFinderError) deviceARNFinder {
	return func(device platforms.EitherDevice) (devicefarm.DeviceARN, *deviceARNFinderError) {
		return deviceARN, err
	}
}
