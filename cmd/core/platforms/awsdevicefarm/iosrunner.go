package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
)

func newIOSRunner(launchRemoteAgent RemoteAgentLauncher, waitRunResult RunResultWaiter) platforms.IOSRunner {
	return func(plan platforms.IOSPlan) error {
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
