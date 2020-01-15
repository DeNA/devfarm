package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func newAndroidRunnerWithRetry(logger logging.SeverityLogger, launchRemoteAgent RemoteAgentLauncher, waitRunResult RunResultWaiter, retryCount int) platforms.AndroidRunner {
	remainedRetryCount := retryCount

	var runAndroid func(plan platforms.AndroidPlan) error
	runAndroid = func(plan platforms.AndroidPlan) error {
		opts := newAndroidAgentLauncherOpts(
			plan.AndroidSpecificPart.APK,
			plan.AndroidSpecificPart.AppID,
			plan.AndroidSpecificPart.IntentExtras,
			plan.AndroidSpecificPart.Device,
			plan.CommonPart.Lifetime,
			remoteAgentSubCmdIsRun,
		)
		intermediates, launchingErr := launchRemoteAgent(plan.CommonPart.GroupName, opts)
		if launchingErr != nil {
			return launchingErr
		}

		runResult, waitErr := waitRunResult(intermediates.runARN)
		if waitErr != nil {
			return waitErr
		}

		switch runResult {
		case devicefarm.RunResultIsPassed:
			return nil

		case devicefarm.RunResultIsErrored:
			// XXX: Retry to avoid "Failed to setup network shaper" that caused by AWS Device Farm.
			//      These errors were happened on 25% runs.
			if remainedRetryCount > 0 {
				remainedRetryCount--
				logger.Info("Retry because an error occurred (and errors does not mean test failures)")

				if retryErr := runAndroid(plan); retryErr != nil {
					return retryErr
				}
				return nil
			}
			return fmt.Errorf("an error occurred (NOTE: test errors does not mean test failures): %q", runResult)
		}

		return fmt.Errorf("test not passed: %q", runResult)
	}

	return runAndroid
}
