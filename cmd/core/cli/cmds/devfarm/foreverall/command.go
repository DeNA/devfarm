package foreverall

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/formatter"
	"github.com/dena/devfarm/cmd/core/cli/planfile"
	"github.com/dena/devfarm/cmd/core/cli/subcmd"
	"github.com/dena/devfarm/cmd/core/platforms/all"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "launches multiple iOS/Android apps and restarts automatically if crashed during the lifetime",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := takeOptions(args)
	if optsErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	bag := cli.ComposeBag(procInout, opts.verbose, opts.dryRun)

	planFile, planFileErr := planfile.Read(opts.planFile, bag.GetFileOpener())
	if planFileErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, fmt.Sprintf("invalid plan file:\n%s", planFileErr.Error()))
		return cli.ExitAbnormal
	}

	ps := all.NewPlatforms(bag)
	table, foreverErr := ps.Forever(planFile.Plans())

	successMsg := "launching"
	if _, err := fmt.Fprint(procInout.Stdout, formatter.PrettyTSV(table.TextTable(successMsg))); err != nil {
		return cli.ExitAbnormal
	}

	if foreverErr != nil {
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}
