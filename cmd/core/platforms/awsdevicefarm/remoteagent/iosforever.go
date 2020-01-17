package remoteagent

import (
	"context"
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/iosdeploy"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/dena/devfarm/cmd/core/platforms"
	"time"
)

type iosForever func(udid iosdeploy.UDID, unarchivedAppPath iosdeploy.UnarchivedAppPath, iosArgs platforms.IOSArgs, lifetime time.Duration) error

func newIOSForever(
	logger logging.SeverityLogger,
	launchApp iosdeploy.AppLauncher,
) iosForever {
	return func(udid iosdeploy.UDID, unarchivedAppPath iosdeploy.UnarchivedAppPath, iosArgs platforms.IOSArgs, lifetime time.Duration) error {
		ctx, cancel := context.WithTimeout(context.Background(), lifetime)
		defer cancel()

		if err := launchApp(ctx, unarchivedAppPath, udid, []string(iosArgs)); err != nil {
			logger.Debug(fmt.Sprintf("app forever: failed to launch app: %s", err.Error()))
			return err
		}

		for {
			select {
			case <-ctx.Done():
				logger.Debug(fmt.Sprintf("app forever: canceled because: %s", ctx.Err().Error()))
				return nil
			default:
			}

			logger.Debug(fmt.Sprintf("app forever: launching app again"))

			if err := launchApp(ctx, unarchivedAppPath, udid, []string(iosArgs)); err != nil {
				logger.Debug(fmt.Sprintf("app forever: failed to launch app: %s", err.Error()))
				return err
			}

			logger.Debug(fmt.Sprintf("app forever: app seemed exited or crashed"))
		}
	}
}
