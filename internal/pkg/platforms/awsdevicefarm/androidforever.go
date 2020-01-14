package awsdevicefarm

import "github.com/dena/devfarm/internal/pkg/platforms"

func newAndroidForever(launchRemoteAgent RemoteAgentLauncher) platforms.AndroidForever {
	return func(plan platforms.AndroidPlan) error {
		opts := newAndroidAgentLauncherOpts(
			plan.AndroidSpecificPart.APK,
			plan.AndroidSpecificPart.AppID,
			plan.AndroidSpecificPart.IntentExtras,
			plan.AndroidSpecificPart.Device,
			plan.CommonPart.Lifetime,
			remoteAgentSubCmdIsForever,
		)
		if _, err := launchRemoteAgent(plan.CommonPart.GroupName, opts); err != nil {
			return err
		}
		return nil
	}
}
