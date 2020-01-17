package awscli

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
)

type Bag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() exec.Executor
	GetFinder() exec.ExecutableFinder
}
