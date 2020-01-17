package runandroid

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/formatter"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms/all"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "runs Android app",
	Command: command,
}

func command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := takeOptions(args)
	if optsErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	bag := cli.ComposeBag(procInout, opts.verbose, opts.dryRun)
	plan := platforms.NewAndroidPlan(
		opts.platformID,
		opts.instanceGroupName,
		opts.device,
		opts.apkPath,
		opts.appID,
		opts.intentExtras,
		opts.lifetime,
		platforms.LocationHintForCLIArguments,
	)

	ps := all.NewPlatforms(bag)

	runningErr := ps.RunAndroid(plan)
	if _, err := fmt.Fprint(procInout.Stdout, formatter.PrettyTSV(format(opts.platformID, runningErr))); err != nil {
		return cli.ExitAbnormal
	}

	if runningErr != nil {
		return cli.ExitAbnormal
	}
	return cli.ExitNormal
}

func format(platformID platforms.ID, err error) [][]string {
	result := make([][]string, 2)

	header := []string{"platform", "status", "note"}
	result[0] = header

	var status string
	var note string
	if err != nil {
		status = "error"
		note = err.Error()
	} else {
		status = "passed"
		note = ""
	}

	result[1] = []string{
		string(platformID),
		status,
		note,
	}

	return result
}
