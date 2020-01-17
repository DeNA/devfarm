package main

import (
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/authstatus"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/catasset"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/foreverall"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/foreverandroid"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/foreverios"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/halt"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/initialize"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/listdevices"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/lsassets"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/runall"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/runandroid"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/runios"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/status"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/validate"
	"github.com/dena/devfarm/cmd/core/cli/cmds/devfarm/version"
	"github.com/dena/devfarm/cmd/core/cli/subcmd"
	"os"
)

var cmdTable subcmd.CommandTable = map[string]subcmd.SubCommandDef{
	"auth-status":     authstatus.CommandDef,
	"cat-asset":       catasset.CommandDef,
	"forever-all":     foreverall.CommandDef,
	"forever-android": foreverandroid.CommandDef,
	"forever-ios":     foreverios.CommandDef,
	"halt":            halt.CommandDef,
	"init":            initialize.CommandDef,
	"list-devices":    listdevices.CommandDef,
	"ls-assets":       lsassets.CommandDef,
	"run-all":         runall.CommandDef,
	"run-android":     runandroid.CommandDef,
	"run-ios":         runios.CommandDef,
	"status":          status.CommandDef,
	"validate":        validate.CommandDef,
	"version":         version.CommandDef,
}

var MainCommand = subcmd.NewSubCommand(cmdTable, subcmd.HelpFallbackCommand(cmdTable))

func main() {
	procInout := cli.GetProcessInout()
	exitStatus := MainCommand(os.Args[1:], procInout)
	os.Exit(int(exitStatus))
}
