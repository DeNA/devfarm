package lsassets

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/assets"
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/subcmd"
	"sort"
	"strings"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "lists bundled assets",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	nonFlagged, optsErr := cli.OnlyHelpOpts(args)
	if optsErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	if len(nonFlagged) > 0 {
		_, _ = fmt.Fprintln(procInout.Stderr, "cannot take any arguments")
		return cli.ExitAbnormal
	}

	lines := make([]string, len(assets.AllAssets))

	for i, assetID := range assets.AllAssets {
		lines[i] = string(assetID)
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})

	if _, err := fmt.Fprintln(procInout.Stdout, strings.Join(lines, "\n")); err != nil {
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}
