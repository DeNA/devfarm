package platforms

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/logging"
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
