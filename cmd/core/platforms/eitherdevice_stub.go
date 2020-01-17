package platforms

func AnyIOSOrAndroidDevice() EitherDevice {
	return EitherDevice{OSName: OSIsUnavailable}
}
