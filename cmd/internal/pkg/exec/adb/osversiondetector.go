package adb

import (
	"errors"
)

type OSVersion string

type OSVersionDetector func(serialNumber SerialNumber) (OSVersion, error)

func NewOSVersionDetector(getProp PropGetter) OSVersionDetector {
	return func(serialNumber SerialNumber) (OSVersion, error) {
		unsafeOSVersion, err := getProp(serialNumber, ROBuildVersionRelease)
		if err != nil {
			return "", err
		}

		if len(unsafeOSVersion) < 1 {
			return "", errors.New("os version must not be empty")
		}

		return OSVersion(unsafeOSVersion), nil
	}
}
