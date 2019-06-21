package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type Bag struct {
	logger              logging.SeverityLogger
	executor            executor.Executor
	interactiveExecutor executor.InteractiveExecutor
	finder              executor.ExecutableFinder
	uploader            executor.Uploader
	fileOpener          executor.FileOpener
	envGetter           executor.EnvGetter
}

func (bag Bag) GetLogger() logging.SeverityLogger {
	return bag.logger
}

func (bag Bag) GetExecutor() executor.Executor {
	return bag.executor
}

func (bag Bag) GetInteractiveExecutor() executor.InteractiveExecutor {
	return bag.interactiveExecutor
}

func (bag Bag) GetFinder() executor.ExecutableFinder {
	return bag.finder
}

func (bag Bag) GetUploader() executor.Uploader {
	return bag.uploader
}

func (bag Bag) GetFileOpener() executor.FileOpener {
	return bag.fileOpener
}

func (bag Bag) GetEnvGetter() executor.EnvGetter {
	return bag.envGetter
}

func NewBag(
	logger logging.SeverityLogger,
	executor executor.Executor,
	interactiveExecutor executor.InteractiveExecutor,
	finder executor.ExecutableFinder,
	uploader executor.Uploader,
	fileOpener executor.FileOpener,
	envGetter executor.EnvGetter,
) Bag {
	return Bag{
		logger:              logger,
		executor:            executor,
		interactiveExecutor: interactiveExecutor,
		finder:              finder,
		uploader:            uploader,
		fileOpener:          fileOpener,
		envGetter:           envGetter,
	}
}

var _ IOSForeverBag = Bag{}
var _ AndroidForeverBag = Bag{}
var _ IOSRunnerBag = Bag{}
var _ AndroidRunnerBag = Bag{}
