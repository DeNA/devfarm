package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

type projectARNFinderError struct {
	NotFound    error
	Unspecified error
}

func (e projectARNFinderError) Error() string {
	if e.NotFound != nil {
		return e.NotFound.Error()
	}
	return e.Unspecified.Error()
}

type projectARNFinder func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, *projectARNFinderError)

func newProjectARNFinder(logger logging.SeverityLogger, listProjects devicefarm.ProjectLister) projectARNFinder {
	return func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, *projectARNFinderError) {
		projectNameToSearch := devicefarm.FromInstanceGroupName(groupName)
		logger.Info(fmt.Sprintf("finding the AWS Device Farm project: %q", projectNameToSearch))

		projects, listErr := listProjects()
		if listErr != nil {
			logger.Error(fmt.Sprintf("failed to list all AWS Device Farm projects: %s", listErr.Error()))
			return "", &projectARNFinderError{
				Unspecified: listErr,
			}
		}

		for _, project := range projects {
			if project.Name == projectNameToSearch {
				logger.Info("the AWS Device Farm project was found")
				logger.Debug(fmt.Sprintf("project ARN: %q", project.ARN))
				return project.ARN, nil
			}
		}

		logger.Info("the AWS Device Farm project was not found")
		return "", &projectARNFinderError{
			NotFound: fmt.Errorf("no projects related to %v. other projects are: %v", groupName, projects),
		}
	}
}
