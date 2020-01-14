package cli

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"github.com/dena/devfarm/internal/pkg/logging"
)

func NewExecutor(logger logging.SeverityLogger, dryRun bool) exec.Executor {
	return exec.NewExecutor(logger, dryRun)
}

func NewExecutableFinder(logger logging.SeverityLogger, dryRun bool) exec.ExecutableFinder {
	return exec.NewExecutableFinder(logger, dryRun)
}
