package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/logging"
)

type devicePoolARNFinder func(projectARN devicefarm.ProjectARN, deviceARN devicefarm.DeviceARN) (devicefarm.DevicePoolARN, *devicePoolARNFinderError)

type devicePoolARNFinderError struct {
	notFound    error
	unspecified error
}

func (e devicePoolARNFinderError) Error() string {
	if e.notFound != nil {
		return e.notFound.Error()
	}
	return e.unspecified.Error()
}

func newDevicePoolARNFinder(logger logging.SeverityLogger, listDevicePools devicefarm.DevicePoolLister) devicePoolARNFinder {
	return func(projectARN devicefarm.ProjectARN, deviceARN devicefarm.DeviceARN) (devicefarm.DevicePoolARN, *devicePoolARNFinderError) {
		logger.Info("listing AWS Device Farm device pools")
		logger.Debug(fmt.Sprintf("device pool ARN to search: %q", deviceARN))
		devicePools, devicePoolsErr := listDevicePools(projectARN)
		if devicePoolsErr != nil {
			logger.Error("failed to list AWS Device Farm device pools")
			return "", &devicePoolARNFinderError{
				unspecified: devicePoolsErr,
			}
		}

		for _, devicePool := range devicePools {
			if devicePool.Name == devicePoolName(deviceARN) {
				logger.Info("found the AWS Device Farm device pool")
				return devicePool.ARN, nil
			}
		}

		logger.Info("no such AWS Device Farm device pools")
		return "", &devicePoolARNFinderError{
			notFound: fmt.Errorf("no such device pools for device: %q (on %s)", deviceARN, projectARN),
		}
	}
}
