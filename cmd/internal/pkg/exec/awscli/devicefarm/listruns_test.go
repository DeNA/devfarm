package devicefarm

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewRunLister(t *testing.T) {
	execute := StubExecutor([]byte(listRunsJSONExample), []byte{}, nil)
	listRuns := NewRunLister(execute)

	got, err := listRuns("arn:aws:devicefarm:ANY_PROJECT")

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if !reflect.DeepEqual(got, listRunsExample) {
		t.Error(cmp.Diff(listRunsExample, got))
		return
	}
}
