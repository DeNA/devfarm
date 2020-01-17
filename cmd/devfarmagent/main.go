package main

import (
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/cmds/devfarm/version"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/cmds/devfarmagent/awsdevicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/subcmd"
	"os"
)

var cmdTable subcmd.CommandTable = map[string]subcmd.SubCommandDef{
	"aws-device-farm": awsdevicefarm.CommandDef,
	"version":         version.CommandDef,
}

var MainCommand = subcmd.NewSubCommand(cmdTable, subcmd.HelpFallbackCommand(cmdTable))

func main() {
	procInout := cli.GetProcessInout()
	exitStatus := MainCommand(os.Args[1:], procInout)
	os.Exit(int(exitStatus))
}
