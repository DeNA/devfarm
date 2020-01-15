package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"time"
)

type runScheduler func(
	osName platforms.OSName,
	projectARN devicefarm.ProjectARN,
	devicePoolARN devicefarm.DevicePoolARN,
	appUpload appUpload,
	specUpload testSpecUpload,
	packageUpload testPackageUpload,
) (devicefarm.RunARN, error)

func newRunScheduler(logger logging.SeverityLogger, scheduleRun devicefarm.RunScheduler) runScheduler {
	return func(osName platforms.OSName, projectARN devicefarm.ProjectARN, devicePoolARN devicefarm.DevicePoolARN, appUpload appUpload, specUpload testSpecUpload, testPackageUpload testPackageUpload) (devicefarm.RunARN, error) {
		logger.Info("scheduling AWS Device Farm run")
		run, runErr := scheduleRun(
			projectARN,
			devicePoolARN,
			devicefarm.NewTestProp(
				devicefarm.TestTypeIsAppiumNode,
				testPackageUpload.arn,
				specUpload.arn,
			),
			devicefarm.ExecutionConfiguration{
				JobTimeout:         devicefarm.JobTimeout(time.Hour),
				AccountsCleanup:    false,
				AppPackagesCleanup: false,
				VideoCapture:       true,
				// NOTE: Resign is unnecessary only if for Android.
				SkipAppResign: osName == platforms.OSIsAndroid,
			},
			appUpload.arn,
		)
		if runErr != nil {
			logger.Error(fmt.Sprintf("failed to schedule AWS Device Farm run: %s", runErr.Error()))
			return "", runErr
		}

		logger.Info("successfully scheduled")
		return run.ARN, nil
	}
}
