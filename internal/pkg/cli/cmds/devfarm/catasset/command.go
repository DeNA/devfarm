package catasset

import (
	"errors"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/assets"
	"github.com/dena/devfarm/internal/pkg/cli"
	"github.com/dena/devfarm/internal/pkg/cli/subcmd"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "prints bundled asset content",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	nonFlagged, optsErr := cli.OnlyHelpOpts(args)
	if optsErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	assetID, assetErr := validateArgs(nonFlagged)
	if assetErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, assetErr.Error())
		return cli.ExitAbnormal
	}

	if _, err := procInout.Stdout.Write(assets.Read(assetID)); err != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, err.Error())
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}

func validateArgs(args []string) (assets.AssetID, error) {
	if len(args) < 1 {
		return "", errors.New("must specify at least an Asset ID")
	}

	if len(args) > 1 {
		return "", errors.New("too many arguments")
	}

	unsafeAssetID := args[0]

	assetID, assetErr := assets.ValidateID(unsafeAssetID)
	if assetErr != nil {
		return "", fmt.Errorf("%s\ntry to execute `devfarm ls-assets` to see available asset IDs.", assetErr.Error())
	}

	return assetID, nil
}
