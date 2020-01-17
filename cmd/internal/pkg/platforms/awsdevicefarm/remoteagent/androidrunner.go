package remoteagent

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/adb"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"time"
)

type androidRunner func(intentExtras platforms.AndroidIntentExtras, lifetime time.Duration) error

func newAndroidRunner(
	logger logging.SeverityLogger,
	packageName adb.PackageName,
	getSerialNumber adb.SerialNumberGetter,
	waitUntilBecomeReady adb.WaitUntilBecomeReady,
	findMainIntent adb.MainIntentFinder,
	startActivity adb.ActivityStarter,
	watch androidWatcher,
) androidRunner {
	return func(intentExtras platforms.AndroidIntentExtras, lifetime time.Duration) error {
		serialNumber, serialNumberErr := getSerialNumber()
		if serialNumberErr != nil {
			logger.Debug(fmt.Sprintf("android run: failed to get serial number: %s", serialNumberErr.Error()))
			return serialNumberErr
		}

		if waitErr := waitUntilBecomeReady(serialNumber); waitErr != nil {
			logger.Debug(fmt.Sprintf("android run: failed to wait until become ready: %s", waitErr.Error()))
			return waitErr
		}

		mainIntent, intentErr := findMainIntent(serialNumber, packageName)
		if intentErr != nil {
			logger.Debug(fmt.Sprintf("android run: failed to find a main intent: %s", intentErr.Error()))
			return intentErr
		}

		if err := startActivity(serialNumber, mainIntent, adb.IntentExtras(intentExtras)); err != nil {
			return err
		}

		if err := watch(serialNumber, lifetime); err != nil {
			return err
		}

		return nil
	}
}
