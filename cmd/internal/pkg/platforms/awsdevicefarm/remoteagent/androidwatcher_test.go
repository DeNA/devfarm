package remoteagent

import (
	"bytes"
	"context"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"io"
	"strings"
	"testing"
	"time"
)

func TestAndroidWatcher(t *testing.T) {
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
	spyLogger := logging.SpySeverityLogger()

	stdin := &bytes.Buffer{}
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader := strings.NewReader("")

	handler := adbAmMonitorHandler{
		logger: spyLogger,
		stdin:  stdin,
		stdout: stdoutReader,
		stderr: stderrReader,
	}

	go func() {
		_, _ = io.WriteString(stdoutWriter, startMsg)
		time.Sleep(100 * time.Millisecond)
		_, _ = io.WriteString(stdoutWriter, crashMsg)
	}()

	err := handler.wait(context.Background())

	if err == nil {
		t.Log(spyLogger.Logs.String())
		t.Error("got nil, want error")
		return
	}

	expectedCmd := "q\n" // means (q)uit
	if stdin.String() != expectedCmd {
		t.Log(spyLogger.Logs.String())
		t.Errorf("got %q, want %q", stdin.String(), expectedCmd)
		return
	}
}
