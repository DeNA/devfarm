package adb

import (
	"bufio"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"strings"
)

type AmMonitorCrashWatcher func(reader io.Reader) (crashed bool)

func NewAmMonitorCrashWatcher(logger logging.SeverityLogger) AmMonitorCrashWatcher {
	return func(stdout io.Reader) (crashed bool) {
		scanner := bufio.NewScanner(stdout)

		for scanner.Scan() {
			line := scanner.Text()
			// > Monitoring activity manager...  available commands:
			// > (q)uit: finish monitoring
			// > ** ERROR: PROCESS CRASHED
			// > processName: com.example.app
			isCrashed := strings.Contains(line, "ERROR: PROCESS CRASHED")
			if isCrashed {
				logger.Debug("am monitor: process crashed")
				return true
			}
		}

		logger.Debug("am monitor: process has never crashed")
		return false
	}
}
