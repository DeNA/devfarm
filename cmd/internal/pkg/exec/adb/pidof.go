package adb

import (
	"bytes"
	"strconv"
)

type PID int

type PIDGetter func(serialNumber SerialNumber, packageName PackageName) (PID, *PIDGetterError)

type PIDGetterError struct {
	NotRunning  error
	Unspecified error
}

func (e PIDGetterError) Error() string {
	if e.NotRunning != nil {
		return e.NotRunning.Error()
	}
	return e.Unspecified.Error()
}

func NewPIDGetter(adbCmd Executor) PIDGetter {
	return func(serialNumber SerialNumber, packageName PackageName) (PID, *PIDGetterError) {
		stdout, execErr := adbCmd("-s", string(serialNumber), "shell", "pidof", string(packageName))
		if execErr != nil {
			if execErr.UnexpectedExitStatus != nil {
				return -1, &PIDGetterError{NotRunning: execErr}
			}
			return -1, &PIDGetterError{Unspecified: execErr}
		}

		unsafePID := string(bytes.TrimSpace(stdout))
		pid, parseErr := strconv.ParseInt(unsafePID, 10, 32)
		if parseErr != nil {
			return -1, &PIDGetterError{NotRunning: parseErr}
		}

		return PID(pid), nil
	}
}
