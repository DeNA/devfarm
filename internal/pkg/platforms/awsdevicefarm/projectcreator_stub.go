package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func anySuccessfulProjectCreatorSkipIfExists() projectCreatorIfNotExists {
	return stubProjectCreatorSkipIfExists("arn:aws:devicefarm:ANY_PROJECT_ARN", nil)
}

func stubProjectCreatorSkipIfExists(projectARN devicefarm.ProjectARN, err error) projectCreatorIfNotExists {
	return func(platforms.InstanceGroupName) (devicefarm.ProjectARN, error) {
		return projectARN, err
	}
}
