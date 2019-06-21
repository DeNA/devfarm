package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type AllInstanceListerBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
}

type AllInstanceLister func(bag AllInstanceListerBag) ([]InstanceOrError, error)
