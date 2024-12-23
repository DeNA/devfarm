package run

import (
	"github.com/dena/devfarm/cmd/core/cli"
	"testing"
)

func TestCommand(t *testing.T) {
	procInout := cli.AnyProcInout()
	args := []string{"--help"}

	got := command(args, procInout)

	if got != cli.ExitAbnormal {
		t.Errorf("got %v, want %v", got, cli.ExitAbnormal)
	}
}
