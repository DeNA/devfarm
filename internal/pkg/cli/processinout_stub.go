package cli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyProcInout() ProcessInout {
	return ProcessInout{
		Stdin:  &testutil.ErrorReadCloserStub{},
		Stdout: &testutil.NullWriteCloser{},
		Stderr: &testutil.NullWriteCloser{},
		GetEnv: executor.AnyEnvGetter(),
	}
}
