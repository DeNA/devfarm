package platforms

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
)

type Bag struct {
	logger              logging.SeverityLogger
	executor            exec.Executor
	interactiveExecutor exec.InteractiveExecutor
	finder              exec.ExecutableFinder
	uploader            exec.Uploader
	fileOpener          exec.FileOpener
	envGetter           exec.EnvGetter
}

func (bag Bag) GetLogger() logging.SeverityLogger {
	return bag.logger
}

func (bag Bag) GetExecutor() exec.Executor {
	return bag.executor
}

func (bag Bag) GetInteractiveExecutor() exec.InteractiveExecutor {
	return bag.interactiveExecutor
}

func (bag Bag) GetFinder() exec.ExecutableFinder {
	return bag.finder
}

func (bag Bag) GetUploader() exec.Uploader {
	return bag.uploader
}

func (bag Bag) GetFileOpener() exec.FileOpener {
	return bag.fileOpener
}

func (bag Bag) GetEnvGetter() exec.EnvGetter {
	return bag.envGetter
}

func NewBag(
	logger logging.SeverityLogger,
	executor exec.Executor,
	interactiveExecutor exec.InteractiveExecutor,
	finder exec.ExecutableFinder,
	uploader exec.Uploader,
	fileOpener exec.FileOpener,
	envGetter exec.EnvGetter,
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
