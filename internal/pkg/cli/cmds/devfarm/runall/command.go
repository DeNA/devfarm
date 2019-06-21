package runall

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/cli"
	"github.com/dena/devfarm/internal/pkg/cli/formatter"
	"github.com/dena/devfarm/internal/pkg/cli/planfile"
	"github.com/dena/devfarm/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/internal/pkg/platforms/all"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "runs all iOS/Android apps described in planfile",
	Command: command,
}

func command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
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

	table, runErr := all.RunAll(planFile.Plans(), bag)
	if table == nil {
		panic("result table of RunAll must be not nil")
	}

	successMsg := "running"
	if _, err := fmt.Fprint(procInout.Stdout, formatter.PrettyTSV(table.TextTable(successMsg))); err != nil {
		return cli.ExitAbnormal
	}

	if runErr != nil {
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}
