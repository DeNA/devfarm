package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func newAppHalt(findProjectARN projectARNFinder, listRuns devicefarm.RunLister, stopRun devicefarm.RunStopper) platforms.Halt {
	return func(groupName platforms.InstanceGroupName) (platforms.Results, error) {
		results := platforms.NewResults()

		projectARN, projectARNErr := findProjectARN(groupName)
		if projectARNErr != nil {
			results.AddErrorOrNils(projectARNErr)
			return *results, results.Err()
		}

		runs, runsErr := listRuns(projectARN)
		if runsErr != nil {
			results.AddErrorOrNils(runsErr)
			return *results, results.Err()
		}

		runARNs := make([]devicefarm.RunARN, len(runs))
		for i, run := range runs {
			runARNs[i] = run.ARN
		}

		// NOTE: Stop synchronously to prevent rate limit exceeded.
		for _, run := range runs {
			switch run.Status {
			case devicefarm.RunStatusIsCompleted, devicefarm.RunStatusIsStopping:
				continue

			default:
				err := stopRun(run.ARN)

				if err != nil {
					results.AddErrorOrNils(fmt.Errorf("%s: %s", err.Error(), run.ARN))
				} else {
					results.AddSuccesses(1)
				}
			}
		}

		return *results, results.Err()
	}
}
