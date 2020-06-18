package adb

import (
	"io"
	"io/ioutil"
)

func StubAmMonitorCrashWatcher(crashCh <-chan bool) AmMonitorCrashWatcher {
	return func(_ io.Reader) (crashed bool) {
		return <-crashCh
	}
}

func FakeAmMonitorCrashWatcher(crashed bool) AmMonitorCrashWatcher {
	return func(reader io.Reader) bool {
		_, _ = ioutil.ReadAll(reader)
		return crashed
	}
}
