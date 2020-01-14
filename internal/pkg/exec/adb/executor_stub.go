package adb

func StubExecutor(stdout []byte, err *ExecutorError) Executor {
	return func(...string) ([]byte, *ExecutorError) {
		return stdout, err
	}
}
