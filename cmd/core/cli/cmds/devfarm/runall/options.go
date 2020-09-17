package runall

import (
	"errors"
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/planfile"
	"os"
	"path/filepath"
)

type planFilePath string

type options struct {
	verbose  bool
	dryRun   bool
	planFile planfile.FilePath
}

func takeOptions(args []string) (options, *cli.ErrorAndUsage) {
	flags, usageBuf := cli.NewFlagSet([]string{"[options] <plan.yml>"})

	verbose := cli.DefineVerboseOpts(flags)
	dryRun := cli.DefineDryRunOpts(flags)

	if err := cli.Parse(args, flags, usageBuf); err != nil {
		return options{}, err
	}

	unsafePlanFile, planFileErr := validatePlanFiles(flags.Args())
	if planFileErr != nil {
		return options{}, cli.NewErrorAndUsage(
			planFileErr.Error(),
			usageBuf.String(),
		)
	}
	planFile := planfile.FilePath(unsafePlanFile)

	return options{
		verbose:  *verbose,
		dryRun:   *dryRun,
		planFile: planFile,
	}, nil
}

func validatePlanFiles(unsafePlanFiles []string) (planFilePath, error) {
	if len(unsafePlanFiles) < 1 {
		return "", errors.New("plan file must be specified")
	}

	if len(unsafePlanFiles) > 1 {
		return "", errors.New("too many plan files")
	}

	unsafePlanFilePathAbs, err := filepath.Abs(unsafePlanFiles[0])
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(unsafePlanFilePathAbs); err != nil {
		return "", err
	}

	return planFilePath(unsafePlanFilePathAbs), nil
}
