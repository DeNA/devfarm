package awscli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func StubExecutor(stdout []byte, stderr []byte, err error) Executor {
	return func(args ...string) (executor.Result, error) {
		return executor.NewResult(stdout, stderr), err
	}
}

func AnyExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), testutil.AnyError)
}

func AnySuccessfulExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), nil)
}
