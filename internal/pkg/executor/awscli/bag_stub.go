package awscli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type BagStub struct {
	Logger   logging.SeverityLogger
	Executor executor.Executor
	Finder   executor.ExecutableFinder
}

func AnyBag() BagStub {
	return BagStub{
		Logger:   logging.NullSeverityLogger(),
		Executor: executor.AnyFailedExecutor,
		Finder:   executor.AnyFailedExecutableFinder,
	}
}

var _ Bag = BagStub{}

func (stub BagStub) GetLogger() logging.SeverityLogger {
	return stub.Logger
}

func (stub BagStub) GetExecutor() executor.Executor {
	return stub.Executor
}

func (stub BagStub) GetFinder() executor.ExecutableFinder {
	return stub.Finder
}
