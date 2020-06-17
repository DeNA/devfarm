package exec

import (
	"context"
	"github.com/dena/devfarm/cmd/core/testutil"
)

var AnySuccessfulInteractiveExecutor = &StubInteractiveExecutor{Err: nil}
var AnyFailedInteractiveExecutor = &StubInteractiveExecutor{Err: testutil.AnyError}

type StubInteractiveExecutor struct {
	Err error
}

var _ InteractiveExecutor = &StubInteractiveExecutor{}

func (s *StubInteractiveExecutor) Execute(_ context.Context, _ InteractiveRequest) error {
	return s.Err
}
