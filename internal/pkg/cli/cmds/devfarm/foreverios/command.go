package foreverios

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/cli"
	"github.com/dena/devfarm/internal/pkg/cli/formatter"
	"github.com/dena/devfarm/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"github.com/dena/devfarm/internal/pkg/platforms/all"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "launches iOS app and restarts automatically if crashed during the lifetime",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := takeOptions(args)
	if optsErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	bag := cli.ComposeBag(procInout, opts.verbose, opts.dryRun)
	plan := platforms.NewIOSPlan(
		opts.platformID,
		opts.instanceGroupName,
		opts.device,
		opts.ipaPath,
		opts.iosArgs,
		opts.lifetime,
		platforms.LocationHintForCLIArguments,
	)

	foreverErr := all.ForeverIOS(plan, bag)

	if _, err := fmt.Fprint(procInout.Stdout, formatter.PrettyTSV(format(opts.platformID, foreverErr))); err != nil {
		return cli.ExitAbnormal
	}

	if foreverErr != nil {
		return cli.ExitAbnormal
	}
	return cli.ExitNormal
}

func format(platformID platforms.ID, err error) [][]string {
	result := make([][]string, 2)

	header := []string{"platform", "status"}
	result[0] = header

	var status string
	if err != nil {
		status = "error"
	} else {
		status = "launching"
	}

	result[1] = []string{
		string(platformID),
		status,
	}

	return result
}
