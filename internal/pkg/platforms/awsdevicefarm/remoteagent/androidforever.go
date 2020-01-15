package remoteagent

import (
	"context"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/adb"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"time"
)

type androidForever func(intentExtras platforms.AndroidIntentExtras, lifetime time.Duration) error

func newAndroidForever(
	logger logging.SeverityLogger,
	packageName adb.PackageName,
	getSerialNumber adb.SerialNumberGetter,
	waitUntilBecomeReady adb.WaitUntilBecomeReady,
	findMainIntent adb.MainIntentFinder,
	startActivity adb.ActivityStarter,
	getPID adb.PIDGetter,
) androidForever {
	return func(intentExtras platforms.AndroidIntentExtras, lifetime time.Duration) error {
		serialNumber, serialNumberErr := getSerialNumber()
		if serialNumberErr != nil {
			logger.Debug(fmt.Sprintf("app forever: failed to get serial number: %s", serialNumberErr.Error()))
			return serialNumberErr
		}

		if waitErr := waitUntilBecomeReady(serialNumber); waitErr != nil {
			logger.Debug(fmt.Sprintf("app forever: failed to wait until become ready: %s", waitErr.Error()))
			return waitErr
		}

		mainIntent, intentErr := findMainIntent(serialNumber, packageName)
		if intentErr != nil {
			logger.Debug(fmt.Sprintf("app forever: failed to find a main intent: %s", intentErr.Error()))
			return intentErr
		}

		ctx, cancel := context.WithTimeout(context.Background(), lifetime)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				logger.Debug(fmt.Sprintf("app forever: canceled because: %s", ctx.Err().Error()))
				return nil
			default:
			}

			pid, pidErr := getPID(serialNumber, packageName)
			if pidErr == nil {
				logger.Debug(fmt.Sprintf("app forever: app is still running at %d", pid))

				time.Sleep(time.Second * 5)
				continue
			}

			if pidErr.NotRunning != nil {
				logger.Debug(fmt.Sprintf("app forever: app is not running"))

				startErr := startActivity(serialNumber, mainIntent, adb.IntentExtras(intentExtras))
				if startErr != nil {
					logger.Debug(fmt.Sprintf("app forever: failed to start an activity %q: %s", mainIntent, startErr.Error()))
					return startErr
				}
				continue
			}
			logger.Debug(fmt.Sprintf("app forever: app is not running"))

			return pidErr
		}
	}
}
