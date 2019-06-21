package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func anySuccessfulDeviceARNFinder() deviceARNFinder {
	return stubDeviceARNFinder("arn:aws:devicefarm:ANY_DEVICE_ARN", nil)
}

func stubDeviceARNFinder(deviceARN devicefarm.DeviceARN, err *deviceARNFinderError) deviceARNFinder {
	return func(device platforms.EitherDevice) (devicefarm.DeviceARN, *deviceARNFinderError) {
		return deviceARN, err
	}
}
