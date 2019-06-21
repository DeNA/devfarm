package awscli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
)

type Executor func(args ...string) (executor.Result, error)

func NewExecutor(bag Bag) Executor {
	return func(args ...string) (executor.Result, error) {
		request := executor.NewRequest("aws", args)
		execute := bag.GetExecutor()
		return execute(request)
	}
}
