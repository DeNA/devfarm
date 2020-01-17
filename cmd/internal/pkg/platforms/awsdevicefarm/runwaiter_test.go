package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"testing"
	"time"
)

func TestNewRunResultWaiter(t *testing.T) {
	t.Run("passed", func(t *testing.T) {
		logger := logging.SpySeverityLogger()
		runResult := devicefarm.RunResultIsPending

		// FIXME: Hard to understand what doing...
		waitRunResult := NewRunResultWaiter(logger, func(devicefarm.RunARN) (devicefarm.Run, error) {
			return devicefarm.Run{Result: runResult}, nil
		}, time.Millisecond*10, time.Second*100)

		go func() {
			time.Sleep(time.Millisecond * 50)
			runResult = devicefarm.RunResultIsPassed
		}()

		got, err := waitRunResult("arn:aws:devicefarm:ANY_RUN_ARN")
		if err != nil {
			t.Log(logger.Logs.String())
			t.Errorf("got (_, %v), want (_, nil)", err)
			return
		}

		if got != devicefarm.RunResultIsPassed {
			t.Errorf("got (%q, nil), want (%q, nil)", got, devicefarm.RunResultIsPassed)
			t.Log(logger.Logs.String())
			return
		}
	})

	t.Run("failed", func(t *testing.T) {
		logger := logging.SpySeverityLogger()
		runResult := devicefarm.RunResultIsPending

		// FIXME: Hard to understand what doing...
		waitRunResult := NewRunResultWaiter(logger, func(devicefarm.RunARN) (devicefarm.Run, error) {
			return devicefarm.Run{Result: runResult}, nil
		}, time.Millisecond*10, time.Second*100)

		go func() {
			time.Sleep(time.Millisecond * 50)
			runResult = devicefarm.RunResultIsPassed
		}()

		got, err := waitRunResult("arn:aws:devicefarm:ANY_RUN_ARN")
		if err != nil {
			t.Errorf("got (_, %v), want (_, nil)", err)
			t.Log(logger.Logs.String())
			return
		}

		if got != devicefarm.RunResultIsPassed {
			t.Errorf("got (%q, nil), want (%q, nil)", got, devicefarm.RunResultIsPassed)
			t.Log(logger.Logs.String())
			return
		}
	})
}
