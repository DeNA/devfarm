package adb

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"time"
)

type WaitUntilBecomeReady func(serialNumber SerialNumber) error

func NewWaitUntilBecomeReady(readyDetector ReadyDetector, waitUntil exec.Waiter) WaitUntilBecomeReady {
	return func(serialNumber SerialNumber) error {
		becomesReady := func() (bool, error) { return readyDetector(serialNumber) }
		return waitUntil(becomesReady, "WaitUntilBecomeReady", time.Second, time.Minute*5)
	}
}

type ReadyDetector func(SerialNumber) (bool, error)

func NewReadyDetector(getProp PropGetter) ReadyDetector {
	return func(serialNumber SerialNumber) (bool, error) {
		bootStat, err := getProp(serialNumber, InitSVCBootAnim)
		if err != nil {
			return false, err
		}

		isReady := bootStat == "stopped"
		return !isReady, nil
	}
}
