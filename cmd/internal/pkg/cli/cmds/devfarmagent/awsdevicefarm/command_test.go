package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"testing"
)

func TestCommand(t *testing.T) {
	procInout := cli.AnyProcInout()
	args := []string{"--help"}

	got := CommandDef.Command(args, procInout)

	if got != cli.ExitNormal {
		t.Errorf("got %v, want %v", got, cli.ExitNormal)
	}
}
