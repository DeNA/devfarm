package platforms

func AnyInstance() Instance {
	return NewInstance(
		AnyIOSOrAndroidDevice(),
		InstanceStateIsUnknown,
	)
}
