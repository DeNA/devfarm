package subcmd

import "github.com/dena/devfarm/cmd/core/cli"

func AnySubCommandDef() SubCommandDef {
	return SubCommandDef{
		Desc:    "ANY",
		Command: cli.AnyCommand(),
	}
}
