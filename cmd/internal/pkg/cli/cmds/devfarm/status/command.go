package status

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/formatter"
	"github.com/dena/devfarm/cmd/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms/all"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "prints instances status",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := takeOptions(args)
	if optsErr != nil {
		_, _ = fmt.Fprint(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	bag := cli.ComposeBag(procInout, opts.verbose, opts.dryRun)
	ps := all.NewPlatforms(bag)

	var instanceTable map[platforms.ID]all.InstancesOrError
	if opts.shouldGetFromAllGroups {
		instanceTable = ps.ListAllInstances()
	} else {
		instanceTable = ps.ListInstances(opts.instanceGroupName)
	}

	textTable := format(all.PlatformInstanceEntryFromTable(instanceTable))

	if _, err := fmt.Fprint(procInout.Stdout, formatter.PrettyTSV(textTable)); err != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, err.Error())
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}

func format(entries []all.PlatformInstancesListEntry) [][]string {
	result := make([][]string, len(entries)+1)

	header := []string{"platform", "device", "os", "state", "note"}
	result[0] = header

	for i, platformEntry := range entries {
		var note string

		if platformEntry.Entry.Error != nil {
			note = platformEntry.Entry.Error.Error()
		} else {
			note = ""
		}

		result[i+1] = []string{
			string(platformEntry.PlatformID),
			platformEntry.Entry.Instance.Device.Name(),
			string(platformEntry.Entry.Instance.Device.OSName),
			string(platformEntry.Entry.Instance.State),
			note,
		}
		continue
	}

	return result
}

type options struct {
	verbose                bool
	dryRun                 bool
	instanceGroupName      platforms.InstanceGroupName
	shouldGetFromAllGroups bool
}

func takeOptions(args []string) (*options, *cli.ErrorAndUsage) {
	flags, usageBuf := cli.NewFlagSet([]string{})

	verbose := cli.DefineVerboseOpts(flags)
	dryRun := cli.DefineDryRunOpts(flags)
	unsafeInstanceGroupName := flags.String("instance-group", "", "instance group name to filter (optional)")

	if err := cli.Parse(args, flags, usageBuf); err != nil {
		return nil, err
	}

	shouldGetFromAllGroups := len(*unsafeInstanceGroupName) < 1
	if shouldGetFromAllGroups {
		return &options{
			verbose:                *verbose,
			dryRun:                 *dryRun,
			instanceGroupName:      "",
			shouldGetFromAllGroups: true,
		}, nil
	}

	instanceGroupName, err := platforms.NewInstanceGroupName(*unsafeInstanceGroupName)
	if err != nil {
		flags.Usage()
		return nil, cli.NewErrorAndUsage(
			fmt.Sprintf("--instance-group: %s", err.Error()),
			usageBuf.String(),
		)
	}

	return &options{
		verbose:                *verbose,
		dryRun:                 *dryRun,
		instanceGroupName:      instanceGroupName,
		shouldGetFromAllGroups: false,
	}, nil
}
