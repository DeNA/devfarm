package initialize

import (
	"errors"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/cli"
	"github.com/dena/devfarm/internal/pkg/cli/planfile"
	"github.com/dena/devfarm/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/internal/pkg/executor"
	"os"
	"path"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "Generate configuration files for launch-all",
	Command: command,
}

func command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := takeOpts(args)
	if optsErr != nil {
		_, _ = fmt.Fprintf(procInout.Stderr, "%s\n", optsErr.Error())
		return cli.ExitAbnormal
	}

	logger := cli.NewLogger(opts.verbose, procInout.Stderr)
	opener := executor.NewFileOpener(logger, opts.dryRun)

	template := *planfile.NewTemplate()
	if err := planfile.Write(template, opts.filePath, opener); err != nil {
		_, _ = fmt.Fprintf(procInout.Stderr, "%s\n", err.Error())
		return cli.ExitAbnormal
	}

	if _, err := fmt.Fprintf(procInout.Stdout, "Configuration generated at %q\n", opts.filePath); err != nil {
		_, _ = fmt.Fprintf(procInout.Stderr, "%s\n", err.Error())
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}

type options struct {
	verbose  bool
	dryRun   bool
	filePath planfile.FilePath
}

func takeOpts(args []string) (*options, error) {
	flags, usageBuf := cli.NewFlagSet([]string{"path"})
	verbose := cli.DefineVerboseOpts(flags)
	dryRun := cli.DefineDryRunOpts(flags)

	if err := cli.Parse(args, flags, usageBuf); err != nil {
		return nil, err
	}

	restArgs := flags.Args()
	filePath, filePathErr := validateFilePath(restArgs)
	if filePathErr != nil {
		return nil, filePathErr
	}

	return &options{
		verbose:  *verbose,
		dryRun:   *dryRun,
		filePath: filePath,
	}, nil
}

func validateFilePath(restArgs []string) (planfile.FilePath, error) {
	if len(restArgs) < 1 {
		cwd, wdErr := os.Getwd()
		if wdErr != nil {
			return "", wdErr
		}
		return planfile.FilePath(path.Join(cwd, "planfile.yml")), nil
	}

	if len(restArgs) > 1 {
		return "", errors.New("too many arguments (only 1 argument required)")
	}

	return planfile.FilePath(restArgs[0]), nil
}
