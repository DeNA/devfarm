package awsdevicefarm

import "github.com/dena/devfarm/cmd/internal/pkg/platforms"

func newIOSForever(launchRemoteAgent RemoteAgentLauncher) platforms.IOSForever {
	return func(plan platforms.IOSPlan) error {
		opts := newIOSAgentLauncherOpts(
			plan.IOSSpecificPart.IPA,
			plan.IOSSpecificPart.Args,
			plan.IOSSpecificPart.Device,
			plan.CommonPart.Lifetime,
			remoteAgentSubCmdIsForever,
		)
		if _, err := launchRemoteAgent(plan.CommonPart.GroupName, opts); err != nil {
			return err
		}
		return nil
	}
}
