package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/platforms"
	"time"
)

type remoteAgentLauncherOpts interface {
	testSpecEmbeddedData
	iosOrAndroidDevice() platforms.EitherDevice
	ipaOrApkPath() ipaOrApkPathOnLocal
}

type iosAgentLauncherOpts struct {
	ipaPath     platforms.IPAPathOnLocal
	iosArgs     platforms.IOSArgs
	iosDevice   platforms.IOSDevice
	life        time.Duration
	agentSubCmd remoteAgentSubCmd
}

func newIOSAgentLauncherOpts(
	ipaPath platforms.IPAPathOnLocal,
	iosArgs platforms.IOSArgs,
	iosDevice platforms.IOSDevice,
	lifetime time.Duration,
	remoteAgentSubCmd remoteAgentSubCmd,
) iosAgentLauncherOpts {
	return iosAgentLauncherOpts{
		ipaPath:     ipaPath,
		iosArgs:     iosArgs,
		iosDevice:   iosDevice,
		life:        lifetime,
		agentSubCmd: remoteAgentSubCmd,
	}
}

var _ remoteAgentLauncherOpts = iosAgentLauncherOpts{}

func (i iosAgentLauncherOpts) iosOrAndroidDevice() platforms.EitherDevice {
	return platforms.EitherDevice{OSName: platforms.OSIsIOS, IOS: i.iosDevice}
}

func (i iosAgentLauncherOpts) args() TransportableArgs {
	return TransportableArgs(i.iosArgs)
}

func (i iosAgentLauncherOpts) androidAppID() (platforms.AndroidAppID, bool) {
	return "", false
}

func (i iosAgentLauncherOpts) ipaOrApkPath() ipaOrApkPathOnLocal {
	return ipaOrApkPathOnLocal(i.ipaPath)
}

func (i iosAgentLauncherOpts) lifetime() time.Duration {
	return i.life
}

func (i iosAgentLauncherOpts) remoteAgentSubCmd() remoteAgentSubCmd {
	return i.agentSubCmd
}

type androidAgentLauncherOpts struct {
	apkPath       platforms.APKPathOnLocal
	appID         platforms.AndroidAppID
	intentExtras  platforms.AndroidIntentExtras
	androidDevice platforms.AndroidDevice
	life          time.Duration
	agentSubCmd   remoteAgentSubCmd
}

func newAndroidAgentLauncherOpts(
	apkPath platforms.APKPathOnLocal,
	appID platforms.AndroidAppID,
	intentExtras platforms.AndroidIntentExtras,
	androidDevice platforms.AndroidDevice,
	lifetime time.Duration,
	remoteAgentSubCmd remoteAgentSubCmd,
) androidAgentLauncherOpts {
	return androidAgentLauncherOpts{
		apkPath:       apkPath,
		appID:         appID,
		intentExtras:  intentExtras,
		androidDevice: androidDevice,
		life:          lifetime,
		agentSubCmd:   remoteAgentSubCmd,
	}
}

var _ remoteAgentLauncherOpts = androidAgentLauncherOpts{}

func (a androidAgentLauncherOpts) iosOrAndroidDevice() platforms.EitherDevice {
	return platforms.EitherDevice{OSName: platforms.OSIsAndroid, Android: a.androidDevice}
}

func (a androidAgentLauncherOpts) args() TransportableArgs {
	return TransportableArgs(a.intentExtras)
}

func (a androidAgentLauncherOpts) androidAppID() (platforms.AndroidAppID, bool) {
	return a.appID, true
}

func (a androidAgentLauncherOpts) ipaOrApkPath() ipaOrApkPathOnLocal {
	return ipaOrApkPathOnLocal(a.apkPath)
}

func (a androidAgentLauncherOpts) lifetime() time.Duration {
	return a.life
}

func (a androidAgentLauncherOpts) remoteAgentSubCmd() remoteAgentSubCmd {
	return a.agentSubCmd
}
