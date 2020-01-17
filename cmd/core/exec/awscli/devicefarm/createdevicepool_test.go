package devicefarm

import (
	"reflect"
	"testing"
)

func TestNewDevicePoolCreator(t *testing.T) {
	execute := StubExecutor([]byte(createDevicePoolJSONExample), []byte{}, nil)
	createDevicePool := NewDevicePoolCreator(execute)

	got, err := createDevicePool(
		"arn:aws:devicefarm:ANY_PROJECT",
		"name",
		"desc",
		[]DevicePoolRule{
			NewDeviceARNBasedDevicePoolRule("arn:aws:devicefarm:ANY_DEVICE"),
		},
	)

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if !reflect.DeepEqual(got, devicePoolExample) {
		t.Errorf("got (%v, nil), want (%v, nil)", got, initializedUploadExample)
		return
	}
}
