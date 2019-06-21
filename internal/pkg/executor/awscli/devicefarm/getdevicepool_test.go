package devicefarm

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewDevicePoolGetter(t *testing.T) {
	execute := StubExecutor([]byte(getDevicePoolExampleJSON), []byte{}, nil)
	getDevicePool := NewDevicePoolGetter(execute)

	got, err := getDevicePool("arn:aws:devicefarm:ANY_DEVICE_POOL")

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err.Error())
		return
	}

	if !reflect.DeepEqual(got, devicePoolExample) {
		t.Error(cmp.Diff(devicePoolExample, got))
		return
	}
}
