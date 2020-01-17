package platforms

type IOSDeviceName string

type IOSDevice struct {
	DeviceName IOSDeviceName `json:"name"`
	OSVersion  IOSVersion    `json:"ios_version"`
}

func NewIOSDevice(deviceName IOSDeviceName, osVersion IOSVersion) IOSDevice {
	return IOSDevice{
		DeviceName: deviceName,
		OSVersion:  osVersion,
	}
}

func (d IOSDevice) Less(another IOSDevice) bool {
	if d.OSVersion != another.OSVersion {
		return d.OSVersion < another.OSVersion
	}
	return d.DeviceName < another.DeviceName
}
