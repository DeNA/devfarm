package remoteagent

import (
	"errors"
	"github.com/dena/devfarm/cmd/core/exec/adb"
	"github.com/dena/devfarm/cmd/core/logging"
	"testing"
	"time"
)

func TestAndroidWatcherCrash(t *testing.T) {
	// Timeline:
	//   1. app launched via adb.
	//   2. crashWatcher finished with true (means app get crashed) but adb is still alive.
	//   3. watch() returns androidCrashed.
	var serial adb.SerialNumber = "DUMMY"
	lifetime := time.Second

	spyLogger := logging.SpySeverityLogger()
	amMonitor := adb.FakeAmMonitor(errors.New("adb killed"))
	crashCh := make(chan bool, 1)
	crashCh <- true
	crashWatcher := adb.StubAmMonitorCrashWatcher(crashCh)

	watch := newAndroidWatcher(spyLogger, amMonitor, crashWatcher)

	if err := watch(serial, lifetime); err != androidCrashed {
		t.Logf("\n%s", spyLogger.Logs.String())
		t.Errorf("want androidCrashed, but %v", err)
		return
	}
}

func TestAndroidWatcherAlive(t *testing.T) {
	// Timeline:
	//   1. app launched via adb.
	//   2. adb get killed because lifetime exceeded.
	//   3. crashWatcher finished with false (means app had never been crashed) because stdout was closed.
	//   4. watch() returns nil.
	var serial adb.SerialNumber = "DUMMY"
	lifetime := 100 * time.Millisecond

	spyLogger := logging.SpySeverityLogger()
	amMonitor := adb.FakeAmMonitor(errors.New("adb killed"))
	crashWatcher := adb.FakeAmMonitorCrashWatcher(false)

	watch := newAndroidWatcher(spyLogger, amMonitor, crashWatcher)

	if err := watch(serial, lifetime); err != nil {
		t.Logf("\n%s", spyLogger.Logs.String())
		t.Errorf("want nil, but %v", err)
		return
	}
}

func TestAndroidWatcherAppNormalExit(t *testing.T) {
	// Timeline:
	//   1. app launched via adb.
	//   2. adb exit because app exit.
	//   3. crashWatcher finished with false (means not crashed) because stdout was closed.
	//   4. watch() returns nil.
	var serial adb.SerialNumber = "DUMMY"
	lifetime := 100 * time.Millisecond

	spyLogger := logging.SpySeverityLogger()
	amMonitor := adb.FakeAmMonitor(nil)
	crashWatcher := adb.FakeAmMonitorCrashWatcher(false)

	watch := newAndroidWatcher(spyLogger, amMonitor, crashWatcher)

	if err := watch(serial, lifetime); err != nil {
		t.Logf("\n%s", spyLogger.Logs.String())
		t.Errorf("want nil, but %v", err)
		return
	}
}

func TestAndroidWatcherAdbErrorBeforeTimeout(t *testing.T) {
	// Timeline:
	//   1. try to launch app via adb.
	//   2. adb exit abnormally soon.
	//   3. crashWatcher finished with false (means not crashed) because stdout was closed.
	//   4. watch() returns adb error.
	var serial adb.SerialNumber = "DUMMY"
	lifetime := time.Second

	spyLogger := logging.SpySeverityLogger()
	adbError := errors.New("command adb not found")
	amMonitorErrCh := make(chan error, 1)
	amMonitorErrCh <- adbError
	amMonitor := adb.StubAmMonitor(amMonitorErrCh)
	crashWatcher := adb.FakeAmMonitorCrashWatcher(false)

	watch := newAndroidWatcher(spyLogger, amMonitor, crashWatcher)

	if err := watch(serial, lifetime); err != adbError {
		t.Logf("\n%s", spyLogger.Logs.String())
		t.Errorf("want nil, but %v", err)
		return
	}
}
