package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func anySuccessfulRunScheduler() runScheduler {
	return stubRunScheduler("arn:aws:devicefarm:ANY_RUN", nil)
}

func stubRunScheduler(runARN devicefarm.RunARN, err error) runScheduler {
	return func(platforms.OSName, devicefarm.ProjectARN, devicefarm.DevicePoolARN, appUpload, testSpecUpload, testPackageUpload) (devicefarm.RunARN, error) {
		return runARN, err
	}
}
