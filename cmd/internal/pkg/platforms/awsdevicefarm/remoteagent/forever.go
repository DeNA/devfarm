package remoteagent

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/adb"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/iosdeploy"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
)

type Forever func() error

type ForeverBag interface {
	GetLogger() logging.SeverityLogger
	GetEnvGetter() exec.EnvGetter
	GetFinder() exec.ExecutableFinder
	GetExecutor() exec.Executor
}

func NewForever(bag ForeverBag) Forever {
	return func() error {
		getEnv := bag.GetEnvGetter()
		getEnvVars := newEnvVarsGetter(getEnv)

		envVars, envVarsErr := getEnvVars()
		if envVarsErr != nil {
			return envVarsErr
		}

		switch envVars.OSName {
		case platforms.OSIsIOS:
			getIOSEnvVars := newIOSSpecificEnvVarsGetter(getEnv)

			iosEnvVars, iosEnvVarsErr := getIOSEnvVars()
			if iosEnvVarsErr != nil {
				return iosEnvVarsErr
			}

			udid := iosdeploy.UDID(envVars.DeviceUDID)
			unarchivedAppPath := iosdeploy.UnarchivedAppPath(iosEnvVars.UnarchivedAppPath)

			iosDeployCmd := iosdeploy.NewExecutor(
				bag.GetLogger(),
				iosEnvVars.IOSDeployBin,
			)
			appForever := newIOSForever(
				bag.GetLogger(),
				iosdeploy.NewAppLauncher(iosDeployCmd),
			)
			return appForever(udid, unarchivedAppPath, platforms.IOSArgs(envVars.AppArgs), envVars.Lifetime)

		case platforms.OSIsAndroid:
			getAndroidEnvVars := newAndroidSpecificEnvVarsGetter(getEnv)
			androidEnvVars, androidEnvVarsErr := getAndroidEnvVars()
			if androidEnvVarsErr != nil {
				return androidEnvVarsErr
			}

			packageName := adb.PackageName(androidEnvVars.AppID)
			adbCmd := adb.NewExecutor(bag.GetFinder(), bag.GetExecutor())
			getProp := adb.NewPropGetter(adbCmd)
			appForever := newAndroidForever(
				bag.GetLogger(),
				packageName,
				adb.NewSerialNumberGetter(adbCmd),
				adb.NewWaitUntilBecomeReady(adb.NewReadyDetector(getProp), exec.NewWaiter()),
				adb.NewMainIntentFinder(adbCmd),
				adb.NewActivityStarter(adbCmd),
				adb.NewPIDGetter(adbCmd),
			)
			return appForever(platforms.AndroidIntentExtras(envVars.AppArgs), envVars.Lifetime)

		default:
			return fmt.Errorf("unsupported os: %q", envVars.OSName)
		}
	}
}
