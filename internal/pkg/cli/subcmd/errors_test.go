package subcmd

import (
	"github.com/dena/devfarm/internal/pkg/cli"
	"reflect"
	"testing"
)

func TestAvailableCommandNames(t *testing.T) {
	command1 := SubCommandDef{
		Desc:    "command1",
		Command: cli.AnyCommand(),
	}
	command2 := SubCommandDef{
		Desc:    "command2",
		Command: cli.AnyCommand(),
	}

	table := map[string]SubCommandDef{
		"command1": command1,
		"command2": command2,
	}

	got := availableCommandNames(table)
	expected := []string{"command1", "command2"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("availableCommandNames(%v) == %v, but wanted %v", table, got, expected)
	}
}
