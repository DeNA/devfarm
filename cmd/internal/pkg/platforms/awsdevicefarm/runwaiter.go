package awsdevicefarm

import (
	"context"
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"time"
)

type RunResultWaiter func(arn devicefarm.RunARN) (devicefarm.RunResult, error)

func NewRunResultWaiter(logger logging.SeverityLogger, getRun devicefarm.RunGetter, pollingInterval time.Duration, timeout time.Duration) RunResultWaiter {
	return func(runARN devicefarm.RunARN) (devicefarm.RunResult, error) {
		logger.Info("waiting until the AWS Device Farm test result became ready")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				logger.Error(fmt.Sprintf("canceled to wait: %s", ctx.Err().Error()))
				return "", ctx.Err()
			default:
				time.Sleep(pollingInterval)

				run, runErr := getRun(runARN)
				if runErr != nil {
					logger.Error(fmt.Sprintf("failed to wait: %s", runErr.Error()))
					return "", runErr
				}

				switch run.Result {
				case devicefarm.RunResultIsPending:
					continue
				default:
					logger.Error("result seems ready")
					return run.Result, nil
				}
			}
		}
	}
}
