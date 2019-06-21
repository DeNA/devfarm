package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

type allInstanceLister func() ([]platforms.InstanceOrError, error)

func newAllInstanceLister(listProjects devicefarm.ProjectLister, collectInstances instanceCollector) allInstanceLister {
	return func() ([]platforms.InstanceOrError, error) {
		projects, projectsErr := listProjects()
		if projectsErr != nil {
			return nil, projectsErr
		}

		allInstances := make([]platforms.InstanceOrError, 0)

		for _, project := range projects {
			instances, instancesErr := collectInstances(project.ARN)
			if instancesErr != nil {
				return nil, instancesErr
			}

			allInstances = append(allInstances, instances...)
		}

		return allInstances, nil
	}
}
