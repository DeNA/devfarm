package adb

import (
	"bufio"
	"errors"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"strings"
)

type AmMonitorCrashWatcher struct {
	logger logging.SeverityLogger
	stdout io.Reader
}

func NewAmMonitorCrashWatcher(logger logging.SeverityLogger, stdout io.Reader) *AmMonitorCrashWatcher {
	return &AmMonitorCrashWatcher{logger: logger, stdout: stdout}
}

func (h AmMonitorCrashWatcher) Watch() error {
	scanner := bufio.NewScanner(h.stdout)

	for scanner.Scan() {
		line := scanner.Text()
		// > Monitoring activity manager...  available commands:
		// > (q)uit: finish monitoring
		// > ** ERROR: PROCESS CRASHED
		// > processName: com.example.app
		isCrashed := strings.Contains(line, "ERROR: PROCESS CRASHED")
		if isCrashed {
			msg := "am monitor: process crashed"
			h.logger.Debug(msg)
			return errors.New(msg)
		}
	}

	h.logger.Debug("am monitor: process has never crashed")
	return nil
}
