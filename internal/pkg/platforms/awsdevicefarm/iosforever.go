package awsdevicefarm

import "github.com/dena/devfarm/internal/pkg/platforms"

func newIOSForever(launchRemoteAgent remoteAgentLauncher) platforms.IOSForever {
	return func(
		plan platforms.IOSPlan,
		bag platforms.IOSForeverBag,
	) error {
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
