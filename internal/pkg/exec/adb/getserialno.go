package adb

import (
	"bytes"
	"errors"
)

type SerialNumber string

type SerialNumberGetter func() (SerialNumber, error)

func NewSerialNumberGetter(adbCmd Executor) SerialNumberGetter {
	return func() (SerialNumber, error) {
		stdout, err := adbCmd("get-serialno")
		if err != nil {
			return "", err
		}

		unsafeSerialNumber := string(bytes.TrimSpace(stdout))
		if len(unsafeSerialNumber) < 1 {
			return "", errors.New("serial number must not be empty")
		}

		return SerialNumber(unsafeSerialNumber), nil
	}
}
