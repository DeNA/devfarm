package iosdeploy

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/logging"
	"io"
	"os/exec"
	"strings"
)

type ExitCode int

// > Error codes we report on different failures, so scripts can distinguish between user app exit
// > codes and our exit codes. For non app errors we use codes in reserved 128-255 range.
// >
// >     const int exitcode_timeout = 252;
// >     const int exitcode_error = 253;
// >     const int exitcode_app_crash = 254;
// >
// SEE: https://github.com/ios-control/ios-deploy/blob/1d1602092a39b48b0e0b7ce36dbd301af5839ab3/src/ios-deploy/ios-deploy.m#L111-L115
const (
	ExitBecauseTimeout       ExitCode = 252
	ExitBecauseInternalError ExitCode = 253
	ExitBecauseAppWasCrashed ExitCode = 254
)

func (e ExitCode) CausingByApp() bool {
	switch e {
	case ExitBecauseTimeout, ExitBecauseInternalError:
		return false

	case ExitBecauseAppWasCrashed:
		return true

	default:
		return 128 > int(e) || int(e) > 255
	}
}

type Executor func(ctx context.Context, args ...string) error

func NewExecutor(logger logging.SeverityLogger, iosDeployBin string) Executor {
	return func(ctx context.Context, args ...string) error {
		devNull := &bytes.Buffer{}
		reader, writer := io.Pipe()

		cmdCtx, cancelCmd := context.WithCancel(ctx)
		defer cancelCmd()
		cmd := exec.CommandContext(cmdCtx, iosDeployBin, args...)
		cmd.Stdin = devNull
		cmd.Stdout = writer
		cmd.Stderr = writer

		executed, formatErr := formatCmdArgs(iosDeployBin, args)
		if formatErr != nil {
			return formatErr
		}

		logger.Debug(fmt.Sprintf("ios-deploy: start: %s", executed))
		if err := cmd.Start(); err != nil {
			logger.Debug(fmt.Sprintf("ios-deploy: failed to start ios-deploy: %s", err.Error()))
			return err
		}

		logCtx, cancelLog := context.WithCancel(ctx)
		watchLog := newLogWatcher(logger)
		go watchLog(logCtx, reader)
		defer cancelLog()

		if err := cmd.Wait(); err != nil {
			exitCode := ExitCode(cmd.ProcessState.ExitCode())
			if !exitCode.CausingByApp() {
				logger.Debug(fmt.Sprintf("ios-deploy: ios-deploy exited abnormally: %s", err.Error()))
				return err
			}

			logger.Debug(fmt.Sprintf("ios-deploy: app crashed or exited abnormally: %s", err.Error()))
			return nil
		}

		logger.Debug(fmt.Sprintf("ios-deploy: app exited"))
		return nil
	}
}

func newLogWatcher(logger logging.SeverityLogger) func(context.Context, io.Reader) {
	return func(ctx context.Context, reader io.Reader) {
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			select {
			case <-ctx.Done():
				break
			default:
				line := scanner.Text()
				logger.Debug(fmt.Sprintf("ios-deploy: log: %s", line))
			}
		}
	}
}

func formatCmdArgs(iosDeployBin string, args []string) (string, error) {
	quotedArgs := make([]string, len(args))

	for i, arg := range args {
		quotedArg, jsonErr := json.Marshal(arg)
		if jsonErr != nil {
			return "", jsonErr
		}
		quotedArgs[i] = string(quotedArg)
	}

	return fmt.Sprintf("%q %s", iosDeployBin, strings.Join(quotedArgs, " ")), nil
}
