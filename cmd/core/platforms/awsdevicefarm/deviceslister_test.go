package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestDevicesLister(t *testing.T) {
	cases := []struct {
		devices  []devicefarm.Device
		expected []platforms.DeviceOrError
	}{
		{
			devices:  []devicefarm.Device{},
			expected: []platforms.DeviceOrError{},
		},
		{
			devices: []devicefarm.Device{
				devicefarm.DeviceIOSExample(),
			},
			expected: []platforms.DeviceOrError{
				platforms.NewDeviceListEntry(
					platforms.EitherDevice{
						OSName: platforms.OSIsIOS,
						IOS:    platforms.NewIOSDevice("apple iphone xs", "12.0"),
					},
					nil,
				),
			},
		},
		{
			devices: []devicefarm.Device{
				devicefarm.DeviceIOSExample(),
				devicefarm.DeviceAndroidExample(),
			},
			expected: []platforms.DeviceOrError{
				platforms.NewDeviceListEntry(
					platforms.EitherDevice{
						OSName: platforms.OSIsIOS,
						IOS:    platforms.NewIOSDevice("apple iphone xs", "12.0"),
					},
					nil,
				),
				platforms.NewDeviceListEntry(
					platforms.EitherDevice{
						OSName:  platforms.OSIsAndroid,
						Android: platforms.NewAndroidDevice("pixel 2 google pixel 2", "8.1.0"),
					},
					nil,
				),
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("listDevices(%v)", c.devices), func(t *testing.T) {
			listDevices := devicefarm.StubDeviceLister(c.devices, nil)
			listDeviceEntries := newDeviceEntryLister(listDevices)

			got, err := listDeviceEntries()

			if err != nil {
				t.Errorf("got (nil, %v), want (%v, nil)", err, c.expected)
				return
			}

			if !reflect.DeepEqual(got, c.expected) {
				t.Error(cmp.Diff(c.expected, got))
			}
		})
	}
}
