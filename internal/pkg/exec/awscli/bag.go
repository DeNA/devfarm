package awscli

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type Bag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() exec.Executor
	GetFinder() exec.ExecutableFinder
}
