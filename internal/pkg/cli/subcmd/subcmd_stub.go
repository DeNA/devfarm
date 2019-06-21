package subcmd

import "github.com/dena/devfarm/internal/pkg/cli"

func AnySubCommandDef() SubCommandDef {
	return SubCommandDef{
		Desc:    "ANY",
		Command: cli.AnyCommand(),
	}
}
