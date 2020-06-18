package adb

import (
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"testing"
	"time"
)

func TestAndroidWatcherNotCrashed(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	stdoutReader, stdoutWriter := io.Pipe()

	crashWatcher := NewAmMonitorCrashWatcher(spyLogger)

	go func() {
		_, _ = io.WriteString(stdoutWriter, startMsg)
		time.Sleep(100 * time.Millisecond)
		_ = stdoutWriter.Close()
	}()

	if crashed := crashWatcher(stdoutReader); crashed {
		t.Log(spyLogger.Logs.String())
		t.Error("got true, want false")
		return
	}
}

func TestAndroidWatcherCrashed(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	stdoutReader, stdoutWriter := io.Pipe()

	crashWatcher := NewAmMonitorCrashWatcher(spyLogger)

	go func() {
		_, _ = io.WriteString(stdoutWriter, startMsg)
		time.Sleep(100 * time.Millisecond)
		_, _ = io.WriteString(stdoutWriter, crashMsg)
	}()

	if crashed := crashWatcher(stdoutReader); !crashed {
		t.Log(spyLogger.Logs.String())
		t.Error("got false, want true")
		return
	}
}

var startMsg = `Monitoring activity manager...  available commands:
(q)uit: finish monitoring
`
var crashMsg = `** ERROR: PROCESS CRASHED
processName: com.example.apk
processPid: 1234
shortMsg: java.lang.RuntimeException
longMsg: java.lang.RuntimeException
timeMillis: 1568779503316
stack:
java.lang.RuntimeException
...

#

Waiting after crash...  available commands:
(c)ontinue: show crash dialog
(k)ill: immediately kill app
(q)uit: finish monitoring
`
