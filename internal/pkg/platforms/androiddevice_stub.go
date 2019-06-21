package platforms

func AnyAndroidDevice() AndroidDevice {
	return NewAndroidDevice("any device name", "ANY.ANDROID.VERSION")
}
