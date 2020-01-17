package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
)

func anySuccessfulProjectCreator() projectCreator {
	return stubProjectCreator("arn:aws:devicefarm:ANY_PROJECT_ARN", nil)
}

func stubProjectCreator(projectARN devicefarm.ProjectARN, err error) projectCreator {
	return func(platforms.InstanceGroupName) (devicefarm.ProjectARN, error) {
		return projectARN, err
	}
}
