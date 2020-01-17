package listdevices

import (
	"github.com/dena/devfarm/cmd/core/cli"
	"testing"
)

func TestCommand(t *testing.T) {
	procInout := cli.AnyProcInout()

	args := []string{"--dry-run"}
	got := Command(args, procInout)

	if got != cli.ExitNormal {
		t.Errorf("Command(%q, procInout) == %v, but wanted %v", args, got, cli.ExitNormal)
	}
}
