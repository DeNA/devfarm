package cli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/logging"
)

func NewExecutor(logger logging.SeverityLogger, dryRun bool) exec.Executor {
	return exec.NewExecutor(logger, dryRun)
}

func NewExecutableFinder(logger logging.SeverityLogger, dryRun bool) exec.ExecutableFinder {
	return exec.NewExecutableFinder(logger, dryRun)
}
