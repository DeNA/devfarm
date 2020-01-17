package cli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/testutil"
)

func AnyProcInout() ProcessInout {
	return ProcessInout{
		Stdin:  &testutil.ErrorReadCloserStub{},
		Stdout: &testutil.NullWriteCloser{},
		Stderr: &testutil.NullWriteCloser{},
		GetEnv: exec.AnyEnvGetter(),
	}
}
