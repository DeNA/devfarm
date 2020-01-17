package devicefarm

import "testing"

func TestNewRunStopper(t *testing.T) {
	execute := AnySuccessfulExecutor()
	stopRun := NewRunStopper(execute)

	err := stopRun("arn:aws:devicefarm:ANY_RUN")

	if err != nil {
		t.Errorf("got %v, want nil", err)
		return
	}
}
