package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type IOSForeverBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
	GetUploader() executor.Uploader
	GetFileOpener() executor.FileOpener
}

type IOSForever func(IOSPlan, IOSForeverBag) error
