package platforms

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"github.com/dena/devfarm/internal/pkg/logging"
)

func AnyBag() Bag {
	return NewBag(
		logging.NullSeverityLogger(),
		exec.AnyFailedExecutor,
		exec.AnyFailedInteractiveExecutor,
		exec.AnyFailedExecutableFinder,
		exec.AnyFailedUploader(),
		exec.AnyFailedFileOpener(),
		exec.AnyEnvGetter(),
	)
}
