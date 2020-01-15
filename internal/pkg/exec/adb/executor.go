package adb

import (
	"github.com/dena/devfarm/internal/pkg/exec"
)

type Executor func(args ...string) ([]byte, *ExecutorError)

type ExecutorError struct {
	NoSuchCommand        error
	UnexpectedExitStatus error
}

func (e ExecutorError) Error() string {
	if e.NoSuchCommand != nil {
		return e.NoSuchCommand.Error()
	}

	return e.UnexpectedExitStatus.Error()
}

func NewExecutor(find exec.ExecutableFinder, execute exec.Executor) Executor {
	return func(args ...string) ([]byte, *ExecutorError) {
		if err := find("adb"); err != nil {
			return nil, &ExecutorError{NoSuchCommand: err}
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
		result, execErr := execute(exec.NewRequest("adb", args))

		if execErr != nil {
			return nil, &ExecutorError{UnexpectedExitStatus: execErr}
		}

		return result.Stdout, nil
	}
}
