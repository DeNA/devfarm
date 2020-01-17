package status

import (
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
	"testing"
)

func TestCommand(t *testing.T) {
	stderrSpy := testutil.NewWriteCloserSpy(nil)
	procInout := cli.AnyProcInout()
	procInout.Stderr = stderrSpy

	args := []string{"--instance-group", "example", "--dry-run"}
	got := Command(args, procInout)

	if got != cli.ExitNormal {
		t.Errorf("Command(%q, procInout) == %v, want %v", args, got, cli.ExitNormal)
		t.Log(stderrSpy.Captured.String())
	}
}
