package devicefarm

import (
	"reflect"
	"testing"
)

func TestNewUploadLister(t *testing.T) {
	execute := StubExecutor([]byte(listUploadsJSONExample), []byte{}, nil)
	listUploads := NewUploadLister(execute)

	got, err := listUploads("arn:aws:devicefarm:ANY_PROJECT")

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if !reflect.DeepEqual(got, listUploadsExample) {
		t.Errorf("got (%v, nil), want (%v, nil)", got, listUploadsExample)
		return
	}
}
