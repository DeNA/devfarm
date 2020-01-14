package devicefarm

import (
	"testing"
)

func TestNewDevicesLister(t *testing.T) {
	execute := StubExecutor([]byte(listDevicesResponseJSONExample), []byte{}, nil)

	listDevices := NewDeviceLister(execute)

	got, err := listDevices()

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		t.Log(got)
		return
	}
}
