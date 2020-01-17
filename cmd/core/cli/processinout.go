package cli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"io"
	"os"
)

type ProcessInout struct {
	Stdin  io.Reader
	Stdout io.WriteCloser
	Stderr io.WriteCloser
	GetEnv exec.EnvGetter
}

func GetProcessInout() ProcessInout {
	return ProcessInout{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		GetEnv: exec.NewEnvGetter(),
	}
}
