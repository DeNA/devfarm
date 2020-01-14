package devicefarm

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewRunGetter(t *testing.T) {
	execute := StubExecutor([]byte(getRunJSONExample), []byte{}, nil)
	getRun := NewRunGetter(execute)

	got, err := getRun(`arn:aws:devicefarm:ANY_RUN_ARN`)

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if got != getRunExample {
		t.Error(cmp.Diff(getRunExample, got))
		return
	}
}
