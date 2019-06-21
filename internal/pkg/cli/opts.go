package cli

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

func NewFlagSet(trailingArgNames []string) (flags *flag.FlagSet, usageBuf *bytes.Buffer) {
	usageBuf = &bytes.Buffer{}
	flags = flag.NewFlagSet("", flag.ContinueOnError)
	flags.SetOutput(usageBuf)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage: %s\n", strings.Join(trailingArgNames, " "))
		flags.PrintDefaults()
	}
	return
}

func Parse(args []string, flags *flag.FlagSet, usageBuf *bytes.Buffer) *ErrorAndUsage {
	if err := flags.Parse(args); err != nil {
		return NewErrorAndUsage("", usageBuf.String())
	}
	return nil
}

func DefineVerboseOpts(flags *flag.FlagSet) *bool {
	return flags.Bool("verbose", false, "enables verbose logs")
}

func DefineDryRunOpts(flags *flag.FlagSet) *bool {
	return flags.Bool("dry-run", false, "enables dry-run (WARNING: not stable yet)")
}

func DefineInstanceGroupNameOpts(flags *flag.FlagSet) *string {
	return flags.String("instance-group", "", "instance group name (required)")
}

func OnlyHelpOpts(args []string) ([]string, *ErrorAndUsage) {
	flags, usageBuf := NewFlagSet([]string{"[--help]"})

	// XXX: flag supports --help and it handle --help as an error.
	if err := Parse(args, flags, usageBuf); err != nil {
		return nil, err
	}

	return flags.Args(), nil
}

func TakeOnlyVerboseAndDryRunOpts(args []string) (*struct {
	Verbose bool
	DryRun  bool
}, *ErrorAndUsage) {
	flags, usageBuf := NewFlagSet([]string{})

	verbose := DefineVerboseOpts(flags)
	dryRun := DefineDryRunOpts(flags)

	if err := Parse(args, flags, usageBuf); err != nil {
		return nil, err
	}

	opts := struct {
		Verbose bool
		DryRun  bool
	}{
		Verbose: *verbose,
		DryRun:  *dryRun,
	}

	return &opts, nil
}
