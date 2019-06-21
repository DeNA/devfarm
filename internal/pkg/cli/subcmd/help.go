package subcmd

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/cli"
)

func HelpFallbackCommand(table CommandTable) cli.Command {
	return func(args []string, procInout cli.ProcessInout) cli.ExitStatus {
		if _, err := cli.OnlyHelpOpts(args); err != nil {
			_, _ = fmt.Fprintln(procInout.Stderr, subCommandUsage(table))
			return cli.ExitNormal
		}

		_, _ = fmt.Fprintln(procInout.Stderr, subCommandUsage(table))
		return cli.ExitAbnormal
	}
}
