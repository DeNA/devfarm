package remoteagent

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/adb"
	"github.com/dena/devfarm/internal/pkg/logging"
	"io"
	"strings"
	"time"
)

type androidWatcher func(serialNumber adb.SerialNumber, lifetime time.Duration) error

func newAndroidWatcher(logger logging.SeverityLogger, amMonitor adb.ActivityMonitor) androidWatcher {
	return func(serialNumber adb.SerialNumber, lifetime time.Duration) error {
		logger.Debug(fmt.Sprintf("android watcher: lifetime is %0.f sec", lifetime.Seconds()))

		stdinReader, stdinWriter := io.Pipe()
		stdoutReader, stdoutWriter := io.Pipe()
		stderrReader, stderrWriter := io.Pipe()

		errCh := make(chan error, 1 /* to prevent goroutine leak */)
		cmdCtx, cancelCmd := context.WithTimeout(context.Background(), lifetime+time.Minute /* to ensure exit */)
		defer cancelCmd()
		watchCtx, cancelWatch := context.WithTimeout(cmdCtx, lifetime)
		defer cancelWatch()

		go func() {
			watcher := adbAmMonitorHandler{
				logger: logger,
				stdin:  stdinWriter,
				stdout: stdoutReader,
				stderr: stderrReader,
			}
			err := watcher.wait(watchCtx)
			errCh <- err
		}()

		go func() {
			<-watchCtx.Done()
			logger.Debug(fmt.Sprintf("android watcher: lifetime (%0.f sec) exceeded", lifetime.Seconds()))
			_ = stdoutWriter.Close()
		}()

		logger.Debug("android watcher: starting am monitor")
		if err := amMonitor(cmdCtx, serialNumber, stdinReader, stdoutWriter, stderrWriter); err != nil {
			logger.Debug(fmt.Sprintf("android watcher: am monitor returned: %s", err.Error()))
			select {
			case <-watchCtx.Done():
				logger.Debug(fmt.Sprintf("android watcher: canceled: %s", watchCtx.Err()))
				// NOTE: Timeout should be treated as success, because the app has been alive until timeout exceeded.
			default:
				return err
			}
		}

		// XXX: Check whether the Android app was crashed or not.
		if err := <-errCh; err != nil {
			return err
		}

		return nil
	}
}

type (
	adbAmMonitorHandler struct {
		logger logging.SeverityLogger
		stdin  io.Writer
		stdout io.Reader
		stderr io.Reader
	}
	adbAmMonitorCommand string
)

func (h adbAmMonitorHandler) wait(ctx context.Context) error {
	scanner := bufio.NewScanner(h.stdout)

	defer func() {
		if err := h.sendCommand(cmdIsQuit); err != nil {
			msg := fmt.Sprintf("am monitor: failed to kill app: %q", err.Error())
			h.logger.Error(msg)
			panic(msg) // XXX: Force to exit the runner to exit adb.
		}
	}()

	for scanner.Scan() {
		line := scanner.Text()
		// > Monitoring activity manager...  available commands:
		// > (q)uit: finish monitoring
		// > ** ERROR: PROCESS CRASHED
		// > processName: com.example.app
		isCrashed := strings.Contains(line, "ERROR: PROCESS CRASHED")
		if isCrashed {
			h.logger.Debug("am monitor: process crashed")
			return errors.New("process crashed")
		}
	}

	h.logger.Debug("am monitor: process has never crashed")
	return nil
}

func (h adbAmMonitorHandler) sendCommand(cmd adbAmMonitorCommand) error {
	h.logger.Debug(fmt.Sprintf("am monitor: send command: %q", cmd))
	if err := cmd.WriteTo(h.stdin); err != nil {
		h.logger.Debug(fmt.Sprintf("am monitor: failed to send command: %q", cmd))
		return err
	}
	return nil
}

const (
	// > Waiting after crash...  available commands:
	// > (c)ontinue: show crash dialog
	// > (k)ill: immediately kill app
	// > (q)uit: finish monitoring
	cmdIsQuit adbAmMonitorCommand = "q"
)

func (c adbAmMonitorCommand) WriteTo(writer io.Writer) error {
	_, err := io.WriteString(writer, fmt.Sprintf("%s\n", c))
	if err != nil {
		return err
	}
	return nil
}
