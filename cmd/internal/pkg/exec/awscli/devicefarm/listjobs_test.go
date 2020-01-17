package devicefarm

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewJobLister(t *testing.T) {
	execute := StubExecutor([]byte(listJobsJSONExample), []byte{}, nil)
	listJobs := NewJobLister(execute)

	got, err := listJobs("arn:aws:devicefarm:ANY_RUN")

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if !reflect.DeepEqual(got, listJobsExample) {
		t.Error(cmp.Diff(listJobsExample, got))
		return
	}
}
