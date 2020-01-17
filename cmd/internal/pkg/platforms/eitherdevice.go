package platforms

import "fmt"

type EitherDevice struct {
	OSName  OSName
	IOS     IOSDevice
	Android AndroidDevice
}

func NewUnavailableEitherDevice() EitherDevice {
	return EitherDevice{OSName: OSIsUnavailable}
}

func (d EitherDevice) Desc() string {
	return fmt.Sprintf("%s %s", d.Name(), d.OS())
}

func (d EitherDevice) Name() string {
	switch d.OSName {
	case OSIsIOS:
		return string(d.IOS.DeviceName)
	case OSIsAndroid:
		return string(d.Android.DeviceName)
	default:
		return "unavailable"
	}
}

func (d EitherDevice) OS() string {
	switch d.OSName {
	case OSIsIOS:
		return fmt.Sprintf("ios %s", d.IOS.OSVersion)
	case OSIsAndroid:
		return fmt.Sprintf("android %s", d.Android.OSVersion)
	default:
		return "unavailable"
	}
}

func (d EitherDevice) Less(another EitherDevice) bool {
	if d.OSName != another.OSName {
		return d.OSName < another.OSName
	}

	if d.IOS != another.IOS {
		return d.IOS.Less(another.IOS)
	}

	return d.Android.Less(another.Android)
}
