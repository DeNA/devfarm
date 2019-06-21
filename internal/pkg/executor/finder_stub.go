package executor

func StubFinder(err error) ExecutableFinder {
	return func(_ string) error {
		return err
	}
}
