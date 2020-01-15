package cli

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func ComposeBag(procInout ProcessInout, verbose bool, dryRun bool) platforms.Bag {
	logger := NewLogger(verbose, procInout.Stderr)
	return platforms.NewBag(
		logger,
		exec.NewExecutor(logger, dryRun),
		exec.NewInteractiveExecutor(logger, dryRun),
		exec.NewExecutableFinder(logger, dryRun),
		exec.NewUploader(logger, dryRun),
		exec.NewFileOpener(logger, dryRun),
		exec.NewEnvGetter(),
	)
}
