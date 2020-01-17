package halt

import (
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestCommand(t *testing.T) {
	stdoutSpy := testutil.NewWriteCloserSpy(nil)
	stderrSpy := testutil.NewWriteCloserSpy(nil)

	args := []string{"--instance-group", "ANY_GROUP", "--dry-run"}
	procInout := cli.AnyProcInout()
	procInout.Stdout = stdoutSpy
	procInout.Stderr = stderrSpy

	// FIXME: Should success if --dry-run
	_ = Command(args, procInout)
}
