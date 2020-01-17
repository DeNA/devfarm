package devicefarm

func StubDeviceLister(devices []Device, err error) DeviceLister {
	return func() ([]Device, error) {
		return devices, err
	}
}
