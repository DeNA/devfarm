package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

func AnyBag() Bag {
	return NewBag(
		logging.NullSeverityLogger(),
		executor.AnyFailedExecutor,
		executor.AnyFailedInteractiveExecutor,
		executor.AnyFailedExecutableFinder,
		executor.AnyFailedUploader(),
		executor.AnyFailedFileOpener(),
		executor.AnyEnvGetter(),
	)
}
