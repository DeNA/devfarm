package platforms

type AndroidDeviceName string

type AndroidDevice struct {
	DeviceName AndroidDeviceName `json:"name"`
	OSVersion  AndroidVersion    `json:"android_version"`
}

func NewAndroidDevice(deviceName AndroidDeviceName, osVersion AndroidVersion) AndroidDevice {
	return AndroidDevice{
		DeviceName: deviceName,
		OSVersion:  osVersion,
	}
}

func (d AndroidDevice) Less(another AndroidDevice) bool {
	if d.OSVersion != another.OSVersion {
		return d.OSVersion < another.OSVersion
	}
	return d.DeviceName < another.DeviceName
}
