package encode

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/subcmd"
	"github.com/dena/devfarm/cmd/core/platforms/awsdevicefarm"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "encode app args to manually run forever-app",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	appArgs, argsErr := awsdevicefarm.EncodeAppArgs(args)
	if argsErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, argsErr.Error())
		return cli.ExitAbnormal
	}
	if _, err := fmt.Fprintln(procInout.Stdout, appArgs); err != nil {
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}
