package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func newIOSRunner(launchRemoteAgent remoteAgentLauncher, waitRunResult runResultWaiter) platforms.IOSForever {
	return func(
		plan platforms.IOSPlan,
		bag platforms.IOSForeverBag,
	) error {
		opts := newIOSAgentLauncherOpts(
			plan.IOSSpecificPart.IPA,
			plan.IOSSpecificPart.Args,
			plan.IOSSpecificPart.Device,
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
