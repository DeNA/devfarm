package awscli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/logging"
)

type BagStub struct {
	Logger   logging.SeverityLogger
	Executor exec.Executor
	Finder   exec.ExecutableFinder
}

func AnyBag() BagStub {
	return BagStub{
		Logger:   logging.NullSeverityLogger(),
		Executor: exec.AnyFailedExecutor,
		Finder:   exec.AnyFailedExecutableFinder,
	}
}

var _ Bag = BagStub{}

func (stub BagStub) GetLogger() logging.SeverityLogger {
	return stub.Logger
}

func (stub BagStub) GetExecutor() exec.Executor {
	return stub.Executor
}

func (stub BagStub) GetFinder() exec.ExecutableFinder {
	return stub.Finder
}
