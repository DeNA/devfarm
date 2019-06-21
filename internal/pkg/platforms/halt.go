package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type HaltBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
}

type Halt func(groupName InstanceGroupName, bag HaltBag) (Results, error)
