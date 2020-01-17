package devicefarm

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
)

func StubExecutor(stdout []byte, stderr []byte, err error) Executor {
	return func(...string) (exec.Result, error) {
		return exec.NewResult(stdout, stderr), err
	}
}

func AnySuccessfulExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), nil)
}
