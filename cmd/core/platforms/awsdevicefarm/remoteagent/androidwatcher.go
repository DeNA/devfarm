package remoteagent

import (
	"context"
	"errors"
	"fmt"
	"github.com/dena/devfarm/cmd/core/contextio"
	"github.com/dena/devfarm/cmd/core/exec/adb"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"time"
)

type androidWatcher func(serialNumber adb.SerialNumber, lifetime time.Duration) error

func newAndroidWatcher(
	logger logging.SeverityLogger,
	amMonitor adb.ActivityMonitor,
	crashWatcher adb.AmMonitorCrashWatcher,
) androidWatcher {
	return func(serialNumber adb.SerialNumber, lifetime time.Duration) error {
		logger.Debug(fmt.Sprintf("android watcher: lifetime is %0.f sec", lifetime.Seconds()))

		stdoutReader, stdoutWriter := io.Pipe()

		crashCh := make(chan bool, 1 /* to prevent goroutines leak */)
		ctx, cancelCmd := context.WithTimeout(context.Background(), lifetime)
		defer cancelCmd()

		go func() {
			// NOTE: Make stdoutReader be cancelable.
			reader := contextio.NewReader(stdoutReader, stdoutWriter, ctx)
			crashCh <- crashWatcher(reader)
		}()

		logger.Debug("android watcher: starting am monitor")
		if err := amMonitor(ctx, serialNumber, nil, stdoutWriter, nil); err != nil {
			logger.Debug(fmt.Sprintf("android watcher: am monitor returned: %s", err.Error()))

			// NOTE: Timeout should be treated as success, because the app has been alive until timeout exceeded.
			if ctxErr := ctx.Err(); ctxErr != nil {
				logger.Debug(fmt.Sprintf("android watcher: context was canceled: %s", ctx.Err()))
				// NOTE: Check whether the Android app was crashed or not.
				if <-crashCh {
					return androidCrashed
				}
				return nil
			}

			// NOTE: Early exit should be treated as failed, because it might be adb command error.
			return err
		}

		// NOTE: App exit normally.
		return nil
	}
}

var androidCrashed = errors.New("android watcher: process crashed")
