package adb

import (
	"bytes"
)

type PropName string

const (
	ROBuildVersionRelease PropName = "ro.build.version.release"
	InitSVCBootAnim       PropName = "init.svc.bootanim"
)

type PropGetter func(serialNumber SerialNumber, propName PropName) (string, error)

func NewPropGetter(adbCmd Executor) PropGetter {
	return func(serialNumber SerialNumber, propName PropName) (string, error) {
		stdout, err := adbCmd("-s", string(serialNumber), "shell", "getprop", string(propName))
		if err != nil {
			return "", err
		}

		return string(bytes.TrimSpace(stdout)), nil
	}
}
