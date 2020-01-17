package devicefarm

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewRunScheduler(t *testing.T) {
	execute := StubExecutor([]byte(scheduleRunJSONExample), []byte{}, nil)
	scheduleRun := NewRunScheduler(execute)

	got, err := scheduleRun(
		"arn:aws:devicefarm:ANY_PROJECT",
		"arn:aws:devicefarm:ANY_DEVICE_POOL",
		AnyTestProp(),
		AnyExecutionConfiguration(),
		"arn:aws:devicefarm:ANY_APP_UPLOAD",
	)

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if got != pendingRunExample {
		t.Error(cmp.Diff(pendingRunExample, got))
		return
	}
}
