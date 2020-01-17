package run

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms/awsdevicefarm/remoteagent"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "launches an iOS/Android app and exit if crashed during the lifetime",
	Command: command,
}

func command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := cli.TakeOnlyVerboseAndDryRunOpts(args)
	if optsErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	bag := cli.ComposeBag(procInout, opts.Verbose, opts.DryRun)
	runApp := remoteagent.NewRunner(bag)

	if err := runApp(); err != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, err.Error())
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}
