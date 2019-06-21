package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"time"
)

type remoteAgentLauncherOptsStubResult struct {
	androidAppID       platforms.AndroidAppID
	androidAppIDOk     bool
	args               TransportableArgs
	iosOrAndroidDevice platforms.EitherDevice
	ipaOrApkPath       ipaOrApkPathOnLocal
	lifetime           time.Duration
}

func anyRemoteAgentLauncherOptsStubResult() remoteAgentLauncherOptsStubResult {
	return remoteAgentLauncherOptsStubResult{
		androidAppID:       "",
		androidAppIDOk:     false,
		args:               AnyTransportableArgs(),
		iosOrAndroidDevice: platforms.NewUnavailableEitherDevice(),
		ipaOrApkPath:       "/path/to/ipa-or-apk",
	}
}

func anyRemoteAgentLauncherOpts() remoteAgentLauncherOptsStub {
	return remoteAgentLauncherOptsStub{anyRemoteAgentLauncherOptsStubResult()}
}

type remoteAgentLauncherOptsStub struct {
	result remoteAgentLauncherOptsStubResult
}

var _ remoteAgentLauncherOpts = remoteAgentLauncherOptsStub{}

func (r remoteAgentLauncherOptsStub) androidAppID() (platforms.AndroidAppID, bool) {
	return r.result.androidAppID, r.result.androidAppIDOk
}

func (r remoteAgentLauncherOptsStub) args() TransportableArgs {
	return r.result.args
}

func (r remoteAgentLauncherOptsStub) iosOrAndroidDevice() platforms.EitherDevice {
	return r.result.iosOrAndroidDevice
}

func (r remoteAgentLauncherOptsStub) ipaOrApkPath() ipaOrApkPathOnLocal {
	return r.result.ipaOrApkPath
}

func (r remoteAgentLauncherOptsStub) lifetime() time.Duration {
	return r.result.lifetime
}

func (r remoteAgentLauncherOptsStub) remoteAgentSubCmd() remoteAgentSubCmd {
	return remoteAgentSubCmdIsRun
}
