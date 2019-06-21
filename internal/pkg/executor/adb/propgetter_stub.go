package adb

func StubPropGetter(propValue string, err error) PropGetter {
	return func(SerialNumber, PropName) (string, error) {
		return propValue, err
	}
}
