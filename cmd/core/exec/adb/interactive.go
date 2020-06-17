package adb

import (
	"context"
	"github.com/dena/devfarm/cmd/core/exec"
	"io"
)

type InteractiveExecutor func(ctx context.Context, stdin io.ReadCloser, stdout io.Writer, stderr io.Writer, args ...string) error

func NewInteractiveExecutor(find exec.ExecutableFinder, executor exec.InteractiveExecutor) InteractiveExecutor {
	return func(ctx context.Context, stdin io.ReadCloser, stdout io.Writer, stderr io.Writer, args ...string) error {
		if err := find("adb"); err != nil {
			return &ExecutorError{NoSuchCommand: err}
		}

		// > global options:
		// >    -a         listen on all network interfaces, not just localhost
		// >    -d         use USB device (error if multiple devices connected)
		// >    -e         use TCP/IP device (error if multiple TCP/IP devices available)
		// >    -s SERIAL  use device with given serial (overrides $ANDROID_SERIAL)
		// >    -t ID      use device with given transport id
		// >    -H         name of adb server host [default=localhost]
		// >    -P         port of adb server [default=5037]
		// >    -L SOCKET  listen on given socket for adb server [default=tcp:localhost:5037]
		req := exec.NewInteractiveRequest(stdin, stdout, stderr, "adb", args)
		if err := executor.Execute(ctx, req); err != nil {
			return &ExecutorError{UnexpectedExitStatus: err}
		}

		return nil
	}
}
