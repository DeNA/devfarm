package version

import (
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestCommand(t *testing.T) {
	stdoutSpy := testutil.NewWriteCloserSpy(nil)
	procInout := cli.AnyProcInout()
	procInout.Stdout = stdoutSpy

	exitStatus := Command([]string{}, procInout)

	if exitStatus != cli.ExitNormal {
		t.Errorf("Expected successfully exit, but abnormal exit")
	}

	stdout := stdoutSpy.Captured.String()

	expectedStdout := "0.0.0\n"
	if stdout != expectedStdout {
		t.Errorf("versionCommand write %q, but wanted %q", stdout, expectedStdout)
	}
}
