package cli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"io"
	"os"
)

type ProcessInout struct {
	Stdin  io.Reader
	Stdout io.WriteCloser
	Stderr io.WriteCloser
	GetEnv executor.EnvGetter
}

func GetProcessInout() ProcessInout {
	return ProcessInout{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		GetEnv: executor.NewEnvGetter(),
	}
}
