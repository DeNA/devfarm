package awscli

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
)

func StubExecutor(stdout []byte, stderr []byte, err error) Executor {
	return func(args ...string) (exec.Result, error) {
		return exec.NewResult(stdout, stderr), err
	}
}

func AnyExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), testutil.AnyError)
}

func AnySuccessfulExecutor() Executor {
	return StubExecutor([]byte("STDOUT"), []byte("STDERR"), nil)
}
