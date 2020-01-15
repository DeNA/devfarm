package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"sync"
)

func devicePoolName(deviceARN devicefarm.DeviceARN) string {
	return fmt.Sprintf("%s%s", devfarmUploadNamePrefix, deviceARN)
}

type devicePoolCreator func(projectARN devicefarm.ProjectARN, deviceARN devicefarm.DeviceARN) (devicefarm.DevicePoolARN, error)

func newDevicePoolCreator(logger logging.SeverityLogger, createDevicePool devicefarm.DevicePoolCreator) devicePoolCreator {
	return func(projectARN devicefarm.ProjectARN, deviceARN devicefarm.DeviceARN) (devicefarm.DevicePoolARN, error) {
		logger.Info("creating AWS Device Farm device pool")
		devicePool, err := createDevicePool(
			projectARN,
			devicePoolName(deviceARN),
			"Auto generated device pool by devfarm",
			[]devicefarm.DevicePoolRule{
				devicefarm.NewDeviceARNBasedDevicePoolRule(deviceARN),
			},
		)

		if err != nil {
			logger.Error(fmt.Sprintf("failed to create the AWS Device Farm device pool: %s", err.Error()))
			return "", err
		}

		logger.Info("created the AWS Device Farm device pool")
		logger.Debug(fmt.Sprintf("device pool ARN: %q", devicePool.ARN))
		return devicePool.ARN, nil
	}
}

func newDevicePoolCreatorIfNotExists(logger logging.SeverityLogger, findDevicePoolARN devicePoolARNFinder, createDevicePool devicePoolCreator) devicePoolCreator {
	return func(projectARN devicefarm.ProjectARN, deviceARN devicefarm.DeviceARN) (devicefarm.DevicePoolARN, error) {
		logger.Info("searching AWS Device Farm device pool to skip creation")
		devicePoolARN, err := findDevicePoolARN(projectARN, deviceARN)
		if err != nil {
			if err.notFound != nil {
				devicePool, devicePoolErr := createDevicePool(projectARN, deviceARN)
				if devicePoolErr != nil {
					return "", devicePoolErr
				}
				return devicePool, nil
			}

			return "", err.unspecified
		}

		logger.Info("skipping to create AWS Device Farm device pool (because already exists)")
		return devicePoolARN, nil
	}
}

func newDevicePoolCreatorCached(createDevicePool devicePoolCreator) devicePoolCreator {
	var mu sync.Mutex
	cache := make(map[projectARNAndDeviceARN]devicefarm.DevicePoolARN)

	return func(projectARN devicefarm.ProjectARN, deviceARN devicefarm.DeviceARN) (devicefarm.DevicePoolARN, error) {
		mu.Lock()
		defer mu.Unlock()

		key := projectARNAndDeviceARN{
			projectARN: projectARN,
			deviceARN:  deviceARN,
		}

		if cached, ok := cache[key]; ok {
			return cached, nil
		}

		devicePoolARN, err := createDevicePool(projectARN, deviceARN)
		if err != nil {
			return "", err
		}

		cache[key] = devicePoolARN
		return devicePoolARN, nil
	}
}

type projectARNAndDeviceARN struct {
	projectARN devicefarm.ProjectARN
	deviceARN  devicefarm.DeviceARN
}
