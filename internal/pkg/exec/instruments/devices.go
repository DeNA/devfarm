package instruments

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec"
	"regexp"
	"strings"
)

type DeviceEntry struct {
	DeviceName  string
	DeviceUDID  string
	OSVersion   string
	IsSimulator bool
}

type DeviceLister func() ([]DeviceEntry, error)

func NewDeviceLister(find exec.ExecutableFinder, execute exec.Executor) DeviceLister {
	return func() ([]DeviceEntry, error) {
		if err := find("instruments"); err != nil {
			return nil, err
		}
		result, execErr := execute(exec.NewRequest("instruments", []string{"-s", "devices"}))
		if execErr != nil {
			return nil, execErr
		}

		return parseRealDevices(result.Stdout)
	}
}

func parseRealDevices(stdoutBytes []byte) ([]DeviceEntry, error) {
	// $ instruments -s devices
	// Known Devices:
	// local-macOS [0CB6923F-EAA8-54E2-B3D2-56A8A5B78803]
	// 9905949119 (12.0) [b1eac0c7cc09b77d902cffdaf912178fd5c5526b]
	// Apple TV (12.2) [76350CD6-C196-418F-9749-A6313DB0CB4E] (Simulator)
	// iPhone Xs Max (12.2) + Apple Watch Series 4 - 44mm (5.2) [940478FB-551B-443C-9B9B-EC15532F00D0] (Simulator)

	scanner := bufio.NewScanner(bytes.NewReader(stdoutBytes))
	entries := make([]DeviceEntry, 0)

	regex := regexp.MustCompile(`(?P<deviceName>.*)\s+\((?P<osVersion>[^(]*)\)\s+\[(?P<deviceUDID>[^\[]*)]\s*(?P<simulatorSuffix>\(Simulator\))?$`)

line:
	for scanner.Scan() {
		line := scanner.Text()

		var deviceName string
		var osVersion string
		var deviceUDID string
		isSimulator := false

		matches := regex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		for i, patternName := range regex.SubexpNames() {
			switch patternName {
			case "deviceName":
				unsafeDeviceName := matches[i]

				hasAppleWatch := strings.Contains(unsafeDeviceName, "Apple Watch")
				if hasAppleWatch {
					// NOTE: iPhone with Apple Watch is out of our scope.
					continue line
				}
				deviceName = unsafeDeviceName
			case "osVersion":
				osVersion = matches[i]
			case "deviceUDID":
				deviceUDID = matches[i]
			case "simulatorSuffix":
				isSimulator = len(matches[i]) > 0
			case "":
				continue
			default:
				return nil, fmt.Errorf("unexpected pattern name: %q in %v", patternName, regex)
			}
		}

		entries = append(entries, DeviceEntry{
			DeviceName:  deviceName,
			OSVersion:   osVersion,
			DeviceUDID:  deviceUDID,
			IsSimulator: isSimulator,
		})
	}

	return entries, nil
}
