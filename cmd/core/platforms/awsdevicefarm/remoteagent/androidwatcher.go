package remoteagent

import (
	"context"
	"fmt"
	"github.com/dena/devfarm/cmd/core/contextio"
	"github.com/dena/devfarm/cmd/core/exec/adb"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

type androidWatcher func(serialNumber adb.SerialNumber, lifetime time.Duration) error

func newAndroidWatcher(logger logging.SeverityLogger, amMonitor adb.ActivityMonitor) androidWatcher {
	return func(serialNumber adb.SerialNumber, lifetime time.Duration) error {
		logger.Debug(fmt.Sprintf("android watcher: lifetime is %0.f sec", lifetime.Seconds()))

		stdinReader := ioutil.NopCloser(strings.NewReader(""))
		stdoutReader, stdoutWriter := io.Pipe()

		errCh := make(chan error, 1 /* to prevent goroutine leak */)
		ctx, cancelCmd := context.WithTimeout(context.Background(), lifetime+time.Minute)
		defer cancelCmd()

		go func() {
			// NOTE: Make stdoutReader be cancelable.
			reader := contextio.NewReader(stdoutReader, stdoutWriter, ctx)
			crashWatcher := adb.NewAmMonitorCrashWatcher(logger, reader)
			errCh <- crashWatcher.Watch()
		}()

		logger.Debug("android watcher: starting am monitor")
		if err := amMonitor(ctx, serialNumber, stdinReader, stdoutWriter, nil); err != nil {
			logger.Debug(fmt.Sprintf("android watcher: am monitor returned: %s", err.Error()))

			// NOTE: Timeout should be treated as success, because the app has been alive until timeout exceeded.
			if ctxErr := ctx.Err(); ctxErr != nil {
				logger.Debug(fmt.Sprintf("android watcher: context was canceled: %s", ctx.Err()))
				return nil
			}

			return err
		}

		// XXX: Check whether the Android app was crashed or not.
		return <-errCh
	}
}
