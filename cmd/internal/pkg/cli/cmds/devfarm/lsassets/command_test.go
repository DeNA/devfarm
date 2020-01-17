package lsassets

import (
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"testing"
)

func TestCommand(t *testing.T) {
	procInout := cli.AnyProcInout()

	got := Command([]string{}, procInout)

	if got != cli.ExitNormal {
		t.Errorf("got %v, want %v", got, cli.ExitNormal)
		return
	}
}
