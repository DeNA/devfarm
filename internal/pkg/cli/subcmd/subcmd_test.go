package subcmd

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/cli"
	"testing"
)

func TestNewSubCommand(t *testing.T) {
	subCmd1 := AnySubCommandDef()
	subCmd1.Desc = "subCmd1"
	subCmd1.Command = cli.SuccessfulCommand()
	subCmd2 := AnySubCommandDef()
	subCmd2.Desc = "subCmd2"
	subCmd2.Command = cli.FailureCommand()
	table := map[string]SubCommandDef{"subCmd1": subCmd1, "subCmd2": subCmd2}

	fallback := cli.FailureCommand()
	procInout := cli.AnyProcInout()

	cases := []struct {
		args     []string
		expected cli.ExitStatus
	}{
		{
			args:     []string{},
			expected: cli.ExitAbnormal,
		},
		{
			args:     []string{"subCmd1"},
			expected: cli.ExitNormal,
		},
		{
			args:     []string{"subCmd1", "--any-flag"},
			expected: cli.ExitNormal,
		},
		{
			args:     []string{"subCmd2"},
			expected: cli.ExitAbnormal,
		},
		{
			args:     []string{"subCmd2", "--any-flag"},
			expected: cli.ExitAbnormal,
		},
		{
			args:     []string{"somethingWrong"},
			expected: cli.ExitAbnormal,
		},
	}

	command := NewSubCommand(table, fallback)

	for _, c := range cases {
		t.Run(fmt.Sprintf("cmd := NewSubCommand({\"command1\": subcmd1, \"command2\": subcmd2}); cmd(%v, procInout)", c.args), func(t *testing.T) {
			got := command(c.args, procInout)

			if got != c.expected {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}

}

func TestIsSubCommand(t *testing.T) {
	cases := []struct {
		name     string
		expected bool
	}{
		{
			name:     "",
			expected: false,
		},
		{
			name:     "subCmd",
			expected: true,
		},
		{
			name:     "--flag",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("isSubCommand(%q)", c.name), func(t *testing.T) {
			got := isSubCommand(c.name)

			if got != c.expected {
				t.Errorf("got %t, want %t", got, c.expected)
			}
		})
	}
}

func TestGetCommand(t *testing.T) {
	command1 := AnySubCommandDef()
	command1.Desc = "command1"
	command2 := AnySubCommandDef()
	command2.Desc = "command2"
	table := map[string]SubCommandDef{"command1": command1, "command2": command2}

	cases := []struct {
		commandName string
		expected    *SubCommandDef
	}{
		{
			commandName: "command1",
			expected:    &command1,
		},
		{
			commandName: "command2",
			expected:    &command2,
		},
		{
			commandName: "something-wrong",
			expected:    nil,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("GetSubCommandDef(%q)", c.commandName), func(t *testing.T) {
			command, err := GetSubCommandDef(c.commandName, table)

			if c.expected == nil {
				if command != nil {
					t.Errorf("got (%v, nil), want (nil, %v)", command, err)
				}
			} else {
				if command == nil {
					t.Errorf("got (nil, %v), want (cmd, nil)", err)
				}
			}
		})
	}
}
