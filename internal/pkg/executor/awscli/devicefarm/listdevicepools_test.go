package devicefarm

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewDevicePoolLister(t *testing.T) {
	execute := StubExecutor([]byte(listDevicePoolsJSONExample), []byte{}, nil)
	listDevicePools := NewDevicePoolLister(execute)

	got, err := listDevicePools("arn:aws:devicefarm:ANY_PROJECT")

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if !reflect.DeepEqual(got, listDevicePoolsExample) {
		t.Error(cmp.Diff(listDevicePoolsExample, got))
		return
	}
}
