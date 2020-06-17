package remoteagent

import (
	"github.com/dena/devfarm/cmd/core/exec/adb"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"testing"
	"time"
)

func TestAndroidWatcherNotCrashed(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	stdoutReader, stdoutWriter := io.Pipe()

	crashWatcher := adb.NewAmMonitorCrashWatcher(spyLogger, stdoutReader)

	go func() {
		_, _ = io.WriteString(stdoutWriter, startMsg)
		time.Sleep(100 * time.Millisecond)
		_ = stdoutWriter.Close()
	}()

	if err := crashWatcher.Watch(); err != nil {
		t.Log(spyLogger.Logs.String())
		t.Errorf("got %v, want nil", err)
		return
	}
}

func TestAndroidWatcherCrashed(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	stdoutReader, stdoutWriter := io.Pipe()

	crashWatcher := adb.NewAmMonitorCrashWatcher(spyLogger, stdoutReader)

	go func() {
		_, _ = io.WriteString(stdoutWriter, startMsg)
		time.Sleep(100 * time.Millisecond)
		_, _ = io.WriteString(stdoutWriter, crashMsg)
	}()

	if err := crashWatcher.Watch(); err == nil {
		t.Log(spyLogger.Logs.String())
		t.Error("got nil, want error")
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
