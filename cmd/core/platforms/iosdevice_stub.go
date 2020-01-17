package platforms

func AnyIOSDevice() IOSDevice {
	return NewIOSDevice("any ios device", "ANY.IOS.VERSION")
}
