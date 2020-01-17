package awscli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/logging"
)

type Bag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() exec.Executor
	GetFinder() exec.ExecutableFinder
}
