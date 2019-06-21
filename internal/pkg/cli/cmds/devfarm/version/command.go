package version

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/cli"
	"github.com/dena/devfarm/internal/pkg/cli/subcmd"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "prints devfarm version",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	if _, err := cli.OnlyHelpOpts(args); err != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, "no options are available")
		return cli.ExitNormal
	}

	if _, err := fmt.Fprintln(procInout.Stdout, "0.0.0"); err != nil {
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}
