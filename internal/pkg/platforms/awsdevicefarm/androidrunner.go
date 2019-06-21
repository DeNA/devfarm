package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func newAndroidRunner(launchRemoteAgent remoteAgentLauncher, waitRunResult runResultWaiter) platforms.AndroidRunner {
	return func(plan platforms.AndroidPlan, bag platforms.AndroidRunnerBag) error {
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

		if runResult != devicefarm.RunResultIsPassed {
			return fmt.Errorf("test not passed: %q", runResult)
		}
		return nil
	}
}
