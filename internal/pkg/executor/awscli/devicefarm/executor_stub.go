package devicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor"
)

func StubExecutor(stdout []byte, stderr []byte, err error) Executor {
	return func(...string) (executor.Result, error) {
		return executor.NewResult(stdout, stderr), err
	}
}

func AnySuccessfulExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), nil)
}
