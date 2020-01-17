package authstatus

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/formatter"
	"github.com/dena/devfarm/cmd/core/cli/subcmd"
	"github.com/dena/devfarm/cmd/core/platforms/all"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "checks authentication status for all platforms",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := cli.TakeOnlyVerboseAndDryRunOpts(args)
	if optsErr != nil {
		_, _ = fmt.Fprint(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	bag := cli.ComposeBag(procInout, opts.Verbose, opts.DryRun)
	ps := all.NewPlatforms(bag)

	authStatusTable := ps.CheckAllAuthStatus()

	result := FormatAuthStatusTable(authStatusTable)

	if _, err := fmt.Fprint(procInout.Stdout, formatter.PrettyTSV(result)); err != nil {
		return cli.ExitAbnormal
	}
	return cli.ExitNormal
}
