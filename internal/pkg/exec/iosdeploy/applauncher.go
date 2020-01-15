package iosdeploy

import (
	"context"
	"strings"
)

type UDID string

// NOTE: ios-deploy can take only extracted .ipa files.
type UnarchivedAppPath string

type AppLauncher func(ctx context.Context, appPath UnarchivedAppPath, udid UDID, appArgs []string) error

func NewAppLauncher(iosDeployCmd Executor) AppLauncher {
	return func(ctx context.Context, appPath UnarchivedAppPath, udid UDID, appArgs []string) error {
		// > -i, --id <device_id>         the id of the device to connect to
		// > -I, --noninteractive         start in non interactive mode (quit when app crashes or exits)
		// > -m, --noinstall              directly start debugging without app install (-d not required)
		// > -b, --bundle <bundle.app>    the path to the app bundle to be installed
		// > -a, --args <args>            command line arguments to pass to the app when launching it
		// > -v, --verbose                enable verbose output
		//
		// SEE: About --args: https://github.com/ios-control/ios-deploy/blob/e1d8a04bd5105857b1c055f8e628404dcbfa8b9c/src/scripts/lldb.py#L49
		err := iosDeployCmd(
			ctx,
			"--verbose",
			"--id", string(udid),
			"--noninteractive",
			"--noinstall",
			"--bundle", string(appPath),
			"--args", strings.Join(appArgs, " "),
		)
		if err != nil {
			return err
		}
		return nil
	}
}

type AppInstallLauncher func(appPath UnarchivedAppPath, udid UDID, appArgs []string) error

func NewAppInstallLauncher(iosDeployCmd Executor) AppInstallLauncher {
	return func(appPath UnarchivedAppPath, udid UDID, appArgs []string) error {
		ctx := context.Background()

		// > -I, --noninteractive         start in non interactive mode (quit when app crashes or exits)
		// > -i, --id <device_id>         the id of the device to connect to
		// > -b, --bundle <bundle.app>    the path to the app bundle to be installed
		// > -a, --args <args>            command line arguments to pass to the app when launching it
		// > -v, --verbose                enable verbose output
		//
		// SEE: About --args: https://github.com/ios-control/ios-deploy/blob/e1d8a04bd5105857b1c055f8e628404dcbfa8b9c/src/scripts/lldb.py#L49
		err := iosDeployCmd(
			ctx,
			"--verbose",
			"--id", string(udid),
			"--noninteractive",
			"--bundle", string(appPath),
			"--args", strings.Join(appArgs, " "),
		)
		if err != nil {
			return err
		}
		return nil
	}
}
