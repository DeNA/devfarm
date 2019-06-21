package decode

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/cli"
	"github.com/dena/devfarm/internal/pkg/cli/subcmd"
	"github.com/dena/devfarm/internal/pkg/platforms/awsdevicefarm"
)

var CommandDef = subcmd.SubCommandDef{
	Desc:    "decode app args to debug",
	Command: Command,
}

func Command(args []string, procInout cli.ProcessInout) cli.ExitStatus {
	if len(args) < 1 {
		_, _ = fmt.Fprintln(procInout.Stderr, "must specify an encoded app args")
		return cli.ExitAbnormal
	}

	if len(args) > 1 {
		_, _ = fmt.Fprintln(procInout.Stderr, "too many arguments")
		return cli.ExitAbnormal
	}

	encoded := args[0]

	decoded, decodeErr := awsdevicefarm.DecodeAppArgs(encoded)
	if decodeErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, decodeErr.Error())
		return cli.ExitAbnormal
	}

	jsonBytes, jsonErr := json.Marshal([]string(decoded))
	if jsonErr != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, jsonErr.Error())
		return cli.ExitAbnormal
	}

	if _, err := fmt.Fprintln(procInout.Stdout, string(jsonBytes)); err != nil {
		_, _ = fmt.Fprintln(procInout.Stderr, err.Error())
		return cli.ExitAbnormal
	}

	return cli.ExitNormal
}
