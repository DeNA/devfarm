package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func newInstanceLister(
	findProjectARN projectARNFinder,
	collectInstances instanceCollector,
) platforms.InstanceLister {
	return func(groupName platforms.InstanceGroupName) ([]platforms.InstanceOrError, error) {
		projectARN, projectErr := findProjectARN(groupName)
		if projectErr != nil {
			return nil, projectErr
		}

		instances, instancesErr := collectInstances(projectARN)
		if instancesErr != nil {
			return nil, instancesErr
		}

		return instances, nil
	}
}

type instanceCollector func(projectARN devicefarm.ProjectARN) ([]platforms.InstanceOrError, error)

func newInstanceCollector(listRuns devicefarm.RunLister, listJobs devicefarm.JobLister) instanceCollector {
	return func(projectARN devicefarm.ProjectARN) ([]platforms.InstanceOrError, error) {
		runs, runsErr := listRuns(projectARN)
		if runsErr != nil {
			return nil, runsErr
		}

		entries := make([]platforms.InstanceOrError, 0)

		for _, run := range runs {
			if run.Status == devicefarm.RunStatusIsCompleted {
				continue
			}

			jobs, jobsErr := listJobs(run.ARN)
			if jobsErr != nil {
				entry := platforms.NewInstanceListEntry(
					platforms.NewInstance(
						platforms.NewUnavailableEitherDevice(),
						platforms.InstanceStateIsUnknown,
					),
					jobsErr,
				)
				entries = append(entries, entry)
				continue
			}

			for _, job := range jobs {
				instanceState, instanceStateErr := instanceStateFrom(job.Status)

				entry := platforms.NewInstanceListEntry(
					platforms.NewInstance(
						iosOrAndroidDeviceFrom(job.Device),
						instanceState,
					),
					instanceStateErr,
				)
				entries = append(entries, entry)
			}
		}

		return entries, nil
	}
}
