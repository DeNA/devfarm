package exec

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyExecutionResult() Result {
	return NewResult(
		[]byte("STDOUT"),
		[]byte("STDERR"),
	)
}

var AnySuccessfulExecutor = StubExecutor([]byte{}, []byte{}, nil)
var AnyFailedExecutor = StubExecutor([]byte{}, []byte{}, nil)

func StubExecutor(stdout []byte, stderr []byte, err error) Executor {
	return func(Request) (Result, error) {
		return NewResult(stdout, stderr), err
	}
}

var AnySuccessfulExecutableFinder = StubExecutableFinder(nil)
var AnyFailedExecutableFinder = StubExecutableFinder(testutil.AnyError)

func StubExecutableFinder(err error) ExecutableFinder {
	return func(string) error {
		return err
	}
}
