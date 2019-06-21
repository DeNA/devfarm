package authstatus

import (
	"github.com/dena/devfarm/internal/pkg/cli"
	"testing"
)

func TestCommand(t *testing.T) {
	procInout := cli.AnyProcInout()

	exitStatus := Command([]string{"--dry-run"}, procInout)

	if exitStatus != cli.ExitNormal {
		t.Errorf("Expected successfully exit, but abnormal exit")
	}
}
