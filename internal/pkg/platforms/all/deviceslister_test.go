package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"github.com/dena/devfarm/internal/pkg/testutil"
	"reflect"
	"testing"
)

func TestListAllDevices(t *testing.T) {
	var examplePlatform1 platforms.ID = "example1"
	var examplePlatform2 platforms.ID = "example2"
	err := testutil.AnyError

	anyIOSDevice := platforms.AnyIOSOrAndroidDevice()
	anyIOSDevice.IOS.DeviceName = "any ios device"
	iosDeviceOrError := platforms.AnyDeviceOrError()
	iosDeviceOrError.Device = anyIOSDevice
	iosDeviceOrError.Error = nil

	anyAndroidDevice := platforms.AnyIOSOrAndroidDevice()
	anyAndroidDevice.Android.DeviceName = "any android device"
	androidDeviceOrError := platforms.AnyDeviceOrError()
	androidDeviceOrError.Device = anyAndroidDevice
	androidDeviceOrError.Error = nil

	cases := []struct {
		table    DevicesListerTable
		expected DevicesTable
	}{
		{
			table:    DevicesListerTable{},
			expected: DevicesTable{},
		},
		{
			table: DevicesListerTable{
				examplePlatform1: platforms.StubDeviceLister([]platforms.DeviceOrError{}, nil),
			},
			expected: DevicesTable{
				examplePlatform1: {entries: []platforms.DeviceOrError{}},
			},
		},
		{
			table: DevicesListerTable{
				examplePlatform1: platforms.StubDeviceLister([]platforms.DeviceOrError{
					iosDeviceOrError,
				}, nil),
			},
			expected: DevicesTable{
				examplePlatform1: {entries: []platforms.DeviceOrError{iosDeviceOrError}},
			},
		},
		{
			table: DevicesListerTable{
				examplePlatform1: platforms.StubDeviceLister([]platforms.DeviceOrError{iosDeviceOrError, androidDeviceOrError}, nil),
			},
			expected: DevicesTable{
				examplePlatform1: {entries: []platforms.DeviceOrError{iosDeviceOrError, androidDeviceOrError}},
			},
		},
		{
			table: DevicesListerTable{
				examplePlatform1: platforms.StubDeviceLister([]platforms.DeviceOrError{iosDeviceOrError}, nil),
				examplePlatform2: platforms.StubDeviceLister([]platforms.DeviceOrError{androidDeviceOrError}, nil),
			},
			expected: DevicesTable{
				examplePlatform1: {entries: []platforms.DeviceOrError{iosDeviceOrError}},
				examplePlatform2: {entries: []platforms.DeviceOrError{androidDeviceOrError}},
			},
		},
		{
			table: DevicesListerTable{
				examplePlatform1: platforms.StubDeviceLister(nil, err),
				examplePlatform2: platforms.StubDeviceLister([]platforms.DeviceOrError{androidDeviceOrError}, nil),
			},
			expected: DevicesTable{
				examplePlatform1: {platformError: err},
				examplePlatform2: {entries: []platforms.DeviceOrError{androidDeviceOrError}},
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("listAllDevicesOn(%#v, bag)", c.table), func(t *testing.T) {
			table := make(map[platforms.ID]platforms.Platform)
			for platformID, deviceLister := range c.table {
				p := platforms.AnyPlatform()
				p.DeviceListerFunc = deviceLister
				table[platformID] = p
			}
			ps := Platforms{table: table}

			got := ps.ListAllDevices()

			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}
}
