package listdevices

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/formatter"
	"github.com/dena/devfarm/cmd/core/cli/subcmd"
	"github.com/dena/devfarm/cmd/core/platforms/all"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "list all available devices",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	opts, optsErr := cli.TakeOnlyVerboseAndDryRunOpts(args)
	if optsErr != nil {
		_, _ = fmt.Fprint(procInout.Stderr, optsErr.Error())
		return cli.ExitAbnormal
	}

	bag := cli.ComposeBag(procInout, opts.Verbose, opts.DryRun)
	ps := all.NewPlatforms(bag)

	table := ps.ListAllDevices()
	entries := all.DeviceListEntries(table)
	textTable := format(entries)

	if _, err := fmt.Fprint(procInout.Stdout, formatter.PrettyTSV(textTable)); err != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, err.Error())
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}

func format(entries []all.PlatformDeviceListEntry) [][]string {
	result := make([][]string, len(entries)+1)

	header := []string{"platform", "os", "device", "available", "note"}
	result[0] = header

	for i, platformEntry := range entries {
		if platformEntry.Entry.Error != nil {
			result[i+1] = []string{
				string(platformEntry.PlatformID),
				platformEntry.Entry.Device.OS(),
				platformEntry.Entry.Device.Name(),
				"no",
				platformEntry.Entry.Error.Error(),
			}
			continue
		}
		result[i+1] = []string{
			string(platformEntry.PlatformID),
			platformEntry.Entry.Device.OS(),
			platformEntry.Entry.Device.Name(),
			"yes",
			"",
		}
	}

	return result
}
