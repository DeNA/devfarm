package decode

import (
	"github.com/dena/devfarm/internal/pkg/cli"
	"github.com/dena/devfarm/internal/pkg/testutil"
	"testing"
)

func TestCommand(t *testing.T) {
	stdout := testutil.NewWriteCloserSpy(nil)
	stderr := testutil.NewWriteCloserSpy(nil)

	args := []string{"W10="}
	procInout := cli.AnyProcInout()
	procInout.Stdout = stdout
	procInout.Stderr = stderr

	got := Command(args, procInout)

	if got != cli.ExitNormal {
		t.Errorf("got %v, want %v", got, cli.ExitNormal)
		t.Log(stderr.Captured.String())
		return
	}

	if stdout.Captured.Len() < 1 {
		t.Errorf("got %q, want any not empty string", stdout.Captured.String())
		return
	}
}
