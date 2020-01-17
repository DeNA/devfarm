package remoteagent

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/platforms/awsdevicefarm"
	"strconv"
	"time"
)

type DeviceUDID string

type EnvVars struct {
	OSName     platforms.OSName
	DeviceUDID DeviceUDID
	AppPath    AppPathOnRemote
	AppArgs    awsdevicefarm.TransportableArgs
	Lifetime   time.Duration
}

type IOSSpecificEnvVars struct {
	UnarchivedAppPath string
	IOSDeployBin      string
}

type AndroidSpecificEnvVars struct {
	AppID platforms.AndroidAppID
}

func NewEnvVars(
	osName platforms.OSName,
	deviceUDID DeviceUDID,
	appPath AppPathOnRemote,
	appArgs awsdevicefarm.TransportableArgs,
	lifetime time.Duration,
) EnvVars {
	return EnvVars{
		OSName:     osName,
		DeviceUDID: deviceUDID,
		AppPath:    appPath,
		AppArgs:    appArgs,
		Lifetime:   lifetime,
	}
}

func NewIOSSpecificEnvVars(unarchivedIOSAppPath string, iosDeployBin string) IOSSpecificEnvVars {
	return IOSSpecificEnvVars{
		UnarchivedAppPath: unarchivedIOSAppPath,
		IOSDeployBin:      iosDeployBin,
	}
}

func NewAndroidSpecificEnvVars(androidAppID platforms.AndroidAppID) AndroidSpecificEnvVars {
	return AndroidSpecificEnvVars{
		AppID: androidAppID,
	}
}

type envVarsGetter func() (EnvVars, error)

func newEnvVarsGetter(getEnv exec.EnvGetter) envVarsGetter {
	return func() (EnvVars, error) {
		// SEE: https://docs.aws.amazon.com/devicefarm/latest/developerguide/custom-test-environment-env.html
		// > The device platform name. It is either Android or iOS.
		unsafeOSName := getEnv("DEVICEFARM_DEVICE_PLATFORM_NAME")

		var osName platforms.OSName
		switch unsafeOSName {
		case "iOS":
			osName = platforms.OSIsIOS
		case "Android":
			osName = platforms.OSIsAndroid
		case "":
			return EnvVars{}, fmt.Errorf("$DEVICEFARM_DEVICE_PLATFORM_NAME is not defined")
		default:
			return EnvVars{}, fmt.Errorf("unsupported os: %q", unsafeOSName)
		}

		// SEE: https://docs.aws.amazon.com/devicefarm/latest/developerguide/custom-test-environment-env.html
		// > The unique identifier of the mobile device running the automated test.
		unsafeDeviceUDID := getEnv("DEVICEFARM_DEVICE_UDID")
		if len(unsafeDeviceUDID) < 1 {
			return EnvVars{}, fmt.Errorf("$DEVICEFARM_DEVICE_UDID is not defined")
		}
		deviceUDID := DeviceUDID(unsafeDeviceUDID)

		// SEE: https://docs.aws.amazon.com/devicefarm/latest/developerguide/custom-test-environment-env.html
		// > The path to the mobile app on the host machine where the tests are being executed. The app path is
		// > available for mobile apps only.
		unsafeAppPath := getEnv("DEVICEFARM_APP_PATH")
		if len(unsafeAppPath) < 1 {
			return EnvVars{}, fmt.Errorf("$DEVICEFARM_APP_PATH is not defined")
		}
		appPath := AppPathOnRemote(unsafeAppPath)

		// NOTE: Embedded by core/pkg/platforms/awsdevicefarm/testspec.go
		unsafeAppArgs := getEnv("DEVFARM_AGENT_APP_ARGS_ENCODED")
		if len(unsafeAppPath) < 1 {
			return EnvVars{}, fmt.Errorf("$DEVFARM_AGENT_APP_ARGS_ENCODED is not defined")
		}
		appArgs, appArgsErr := awsdevicefarm.DecodeAppArgs(unsafeAppArgs)
		if appArgsErr != nil {
			return EnvVars{}, fmt.Errorf("cannot decode $DEVFARM_AGENT_APP_ARGS_ENCODED: %q because %s", unsafeAppArgs, appArgsErr)
		}

		unsafeLifetime := getEnv("DEVFARM_AGENT_LIFETIME_SEC")
		if len(unsafeLifetime) < 1 {
			return EnvVars{}, fmt.Errorf("$DEVFARM_AGENT_LIFETIME_SEC is not defined")
		}
		lifetimeSec, lifetimeErr := strconv.Atoi(unsafeLifetime)
		if lifetimeErr != nil {
			return EnvVars{}, fmt.Errorf("$DEVFARM_AGENT_LIFETIME_SEC must be an integer: %q", unsafeLifetime)
		}
		lifetime := time.Duration(lifetimeSec) * time.Second

		return NewEnvVars(
			osName,
			deviceUDID,
			appPath,
			appArgs,
			lifetime,
		), nil
	}
}

type iosSpecificEnvVarsGetter func() (IOSSpecificEnvVars, error)

func newIOSSpecificEnvVarsGetter(getEnv exec.EnvGetter) iosSpecificEnvVarsGetter {
	return func() (IOSSpecificEnvVars, error) {
		// NOTE: Embedded by assets/aws-device-farm/workflows/0-shared.bash
		unarchivedIOSAppPath := getEnv("DEVFARM_AGENT_UNARCHIVED_IOS_APP_PATH")
		if len(unarchivedIOSAppPath) < 1 {
			return IOSSpecificEnvVars{}, fmt.Errorf("$DEVFARM_AGENT_UNARCHIVED_IOS_APP_PATH is not defined")
		}

		// NOTE: Embedded by assets/aws-device-farm/workflows/0-shared.bash
		iosDeployBin := getEnv("DEVFARM_AGENT_IOS_DEPLOY_BIN")
		if len(iosDeployBin) < 1 {
			return IOSSpecificEnvVars{}, fmt.Errorf("DEVFARM_AGENT_IOS_DEPLOY_BIN is not defined")
		}

		return NewIOSSpecificEnvVars(unarchivedIOSAppPath, iosDeployBin), nil
	}
}

type androidSpecificEnvVarsGetter func() (AndroidSpecificEnvVars, error)

func newAndroidSpecificEnvVarsGetter(getEnv exec.EnvGetter) androidSpecificEnvVarsGetter {
	return func() (AndroidSpecificEnvVars, error) {
		// NOTE: Embedded by assets/aws-device-farm/workflows/0-shared.bash
		unsafeAppID := getEnv("DEVFARM_AGENT_ANDROID_APP_ID_IF_ANDROID")
		if len(unsafeAppID) < 1 {
			return AndroidSpecificEnvVars{}, fmt.Errorf("$DEVFARM_AGENT_ANDROID_APP_ID_IF_ANDROID is not defined")
		}
		appID := platforms.AndroidAppID(unsafeAppID)

		return NewAndroidSpecificEnvVars(appID), nil
	}
}
