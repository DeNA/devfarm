package cli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func ComposeBag(procInout ProcessInout, verbose bool, dryRun bool) platforms.Bag {
	logger := NewLogger(verbose, procInout.Stderr)
	return platforms.NewBag(
		logger,
		executor.NewExecutor(logger, dryRun),
		executor.NewInteractiveExecutor(logger, dryRun),
		executor.NewExecutableFinder(logger, dryRun),
		executor.NewUploader(logger, dryRun),
		executor.NewFileOpener(logger, dryRun),
		executor.NewEnvGetter(),
	)
}
