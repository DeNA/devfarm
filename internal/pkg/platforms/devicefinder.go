package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type DeviceFinderBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
}

type DeviceFinder func(EitherDevice, DeviceFinderBag) (bool, error)
