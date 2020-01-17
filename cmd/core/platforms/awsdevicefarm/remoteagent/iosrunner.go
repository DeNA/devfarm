package remoteagent

import (
	"context"
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/iosdeploy"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/dena/devfarm/cmd/core/platforms"
	"time"
)

type iosAppRunner func(udid iosdeploy.UDID, unarchivedAppPath iosdeploy.UnarchivedAppPath, iosArgs platforms.IOSArgs, lifetime time.Duration) error

func newIOSAppRunner(
	logger logging.SeverityLogger,
	launchApp iosdeploy.AppLauncher,
) iosAppRunner {
	return func(udid iosdeploy.UDID, unarchivedAppPath iosdeploy.UnarchivedAppPath, iosArgs platforms.IOSArgs, lifetime time.Duration) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), lifetime)
		defer cancelFunc()

		if err := launchApp(ctx, unarchivedAppPath, udid, []string(iosArgs)); err != nil {
			logger.Debug(fmt.Sprintf("app forever: failed to launch app: %s", err.Error()))
			return err
		}
		return nil
	}
}
