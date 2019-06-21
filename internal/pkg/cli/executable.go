package cli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

func NewExecutor(logger logging.SeverityLogger, dryRun bool) executor.Executor {
	return executor.NewExecutor(logger, dryRun)
}

func NewExecutableFinder(logger logging.SeverityLogger, dryRun bool) executor.ExecutableFinder {
	return executor.NewExecutableFinder(logger, dryRun)
}
