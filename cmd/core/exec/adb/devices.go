package adb

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type DeviceLister func() ([]DeviceEntry, error)

func NewDeviceLister(adbCmd Executor) DeviceLister {
	return func() ([]DeviceEntry, error) {
		// > general commands:
		// >     devices [-l]             list connected devices (-l for long output)
		stdout, err := adbCmd("devices")
		if err != nil {
			return nil, err
		}

		return parseDevices(stdout)
	}
}

type DeviceState string

const (
	// > offline: The device is not connected to adb or is not responding.
	// > device: The device is now connected to the adb server. Note that this state does not imply that the Android system is fully booted and operational because the device connects to adb while the system is still booting. However, after boot-up, this is the normal operational state of an device.
	// > no device: There is no device connected.
	DeviceIsOffline   DeviceState = "offline"
	DeviceIsConnected DeviceState = "device"
	DeviceIsNotFound  DeviceState = "no device"
)

type DeviceEntry struct {
	Name  string
	State DeviceState
}

func parseDevices(stdoutBytes []byte) ([]DeviceEntry, error) {
	scanner := bufio.NewScanner(bytes.NewReader(stdoutBytes))

	devices := make([]DeviceEntry, 0)

	// $ adb devices
	// List of devices attached
	// emulator-5554   device
	for scanner.Scan() {
		line := scanner.Text()

		isEntry := strings.Contains(line, "\t")
		if !isEntry {
			continue
		}

		components := strings.Split(line, "\t")
		if len(components) < 2 {
			return nil, fmt.Errorf("unrecognized line: %q", line)
		}

		deviceName := components[0]
		var deviceState DeviceState

		switch components[1] {
		case "offline":
			deviceState = DeviceIsOffline
		case "device":
			deviceState = DeviceIsConnected
		case "no device":
			deviceState = DeviceIsNotFound
		default:
			return nil, fmt.Errorf("unrecognized device state: %q", components[1])
		}

		devices = append(devices, DeviceEntry{
			Name:  deviceName,
			State: deviceState,
		})
	}

	return devices, nil
}
