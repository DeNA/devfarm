package validate

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/planfile"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "validates plan.yml",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	flags, usageBuf := cli.NewFlagSet([]string{"<plan.yml>"})

	verbose := cli.DefineVerboseOpts(flags)

	if err := cli.Parse(args, flags, usageBuf); err != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, err.Error())
		return cli.ExitAbnormal
	}

	unsafePlanFiles := flags.Args()
	planFilePath, planFilePathErr := validatePlanFilePath(unsafePlanFiles)
	if planFilePathErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, planFilePathErr.Error())
		return cli.ExitAbnormal
	}

	logger := cli.NewLogger(false, procInout.Stderr)
	open := exec.NewFileOpener(logger, false)

	planFile, planFileErr := planfile.Read(planFilePath, open)
	if planFileErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, planFileErr.Error())
		return cli.ExitAbnormal
	}

	if *verbose {
		planFileJSON, jsonErr := json.MarshalIndent(planFile.Plans(), "", "  ")
		if jsonErr != nil {
			_, _ = fmt.Fprintln(procInout.Stderr, jsonErr.Error())
			return cli.ExitAbnormal
		}

		if _, err := fmt.Fprintln(procInout.Stdout, string(planFileJSON)); err != nil {
			return cli.ExitAbnormal
		}
	}

	return cli.ExitNormal
}

func validatePlanFilePath(unsafePlanFiles []string) (planfile.FilePath, error) {
	if len(unsafePlanFiles) < 1 {
		return "", fmt.Errorf("plan file must be specified")
	}

	if len(unsafePlanFiles) > 1 {
		return "", fmt.Errorf("too many arguments")
	}

	return planfile.FilePath(unsafePlanFiles[0]), nil
}
