package awscli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/testutil"
)

func StubExecutor(stdout []byte, stderr []byte, err error) Executor {
	return func(args ...string) (exec.Result, error) {
		return exec.NewResult(stdout, stderr), err
	}
}

func AnyExecutor() Executor {
	return AnyFailedExecutor()
}

func AnySuccessfulExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), nil)
}

func AnyFailedExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), testutil.AnyError)
}
