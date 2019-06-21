package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/cli/cmds/devfarmagent/awsdevicefarm/args"
	"github.com/dena/devfarm/internal/pkg/cli/cmds/devfarmagent/awsdevicefarm/forever"
	"github.com/dena/devfarm/internal/pkg/cli/cmds/devfarmagent/awsdevicefarm/run"
	"github.com/dena/devfarm/internal/pkg/cli/subcmd"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "agent works on AWS Device Farm",
	Command: subcmd.NewSubCommand(commandTable, subcmd.HelpFallbackCommand(commandTable)),
}

var commandTable subcmd.CommandTable = map[string]subcmd.SubCommandDef{
	"args":    args.CommandDef,
	"forever": forever.CommandDef,
	"run":     run.CommandDef,
}
