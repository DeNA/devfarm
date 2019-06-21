package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type AndroidRunnerBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
	GetUploader() executor.Uploader
	GetFileOpener() executor.FileOpener
}

type AndroidRunner func(AndroidPlan, AndroidRunnerBag) error
