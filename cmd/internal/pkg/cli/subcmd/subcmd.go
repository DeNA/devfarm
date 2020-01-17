package subcmd

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"sort"
	"strings"
)

type SubCommandDef struct {
	Desc    string
	Command cli.Command
}

func NewSubCommand(table CommandTable, defaultCmd cli.Command) cli.Command {
	return func(args []string, procInout cli.ProcessInout) cli.ExitStatus {
		if len(args) < 1 {
			_, _ = fmt.Fprintln(procInout.Stderr, subCommandUsage(table))
			return cli.ExitAbnormal
		}

		if !isSubCommand(args[0]) {
			return defaultCmd(args, procInout)
		}

		subCmdDef, err := GetSubCommandDef(args[0], table)
		if err != nil {
			_, _ = fmt.Fprintln(procInout.Stderr, subCommandUsage(table))
			return cli.ExitAbnormal
		}

		return subCmdDef.Command(args[1:], procInout)
	}
}

func isSubCommand(arg string) bool {
	if len(arg) < 1 {
		return false
	}

	return arg[0] != '-'
}

func longestCommandNameLen(table CommandTable) int {
	longest := -1

	for name := range table {
		nameLen := len(name)
		if nameLen > longest {
			longest = nameLen
		}
	}

	return longest
}

func subCommandUsage(table CommandTable) string {
	commandDescriptions := make([]string, len(table))
	i := 0

	longest := longestCommandNameLen(table)

	for name, command := range table {
		padding := strings.Repeat(" ", longest-len(name))
		commandDescriptions[i] = fmt.Sprintf("    %s%s    %s", name, padding, command.Desc)
		i++
	}

	sort.Strings(commandDescriptions)
	commandPart := strings.Join(commandDescriptions, "\n")

	return fmt.Sprintf(`Error: devfarm [<command>]

%s
`, commandPart)
}
