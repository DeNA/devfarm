package adb

import (
	"context"
	"io"
)

type ActivityMonitor func(ctx context.Context, serialNumber SerialNumber, stdin io.ReadCloser, stdout io.Writer, stderr io.Writer) error

func NewActivityMonitor(adbCmd InteractiveExecutor) ActivityMonitor {
	return func(ctx context.Context, serialNumber SerialNumber, stdin io.ReadCloser, stdout io.Writer, stderr io.Writer) error {
		// > monitor [options] Start monitoring for crashes or ANRs.
		// > Options are:
		// > --gdb: Start gdbserv on the given port at crash/ANR.
		// https://developer.android.com/studio/command-line/adb.html#am
		adbArgs := []string{
			"-s", string(serialNumber),
			"shell",
			"am",
			"monitor",
		}

		if err := adbCmd(ctx, stdin, stdout, stderr, adbArgs...); err != nil {
			return err
		}
		return nil
	}
}
