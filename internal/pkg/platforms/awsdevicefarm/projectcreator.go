package awsdevicefarm

import (
	"fmt"
	"time"

	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

type projectCreator func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, error)

func newProjectCreator(logger logging.SeverityLogger, createProject devicefarm.ProjectCreator) projectCreator {
	return func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, error) {
		projectName := devicefarm.FromInstanceGroupName(groupName)

		logger.Info("creating an AWS Device Farm project...")
		project, err := createProject(projectName, 10*time.Minute)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to create AWS Device Farm project: %s", err.Error()))
			return "", err
		}

		logger.Info("the AWS Device Farm project was successfully created")
		logger.Debug(fmt.Sprintf("created project ARN: %q", project.ARN))
		return project.ARN, nil
	}
}

type projectCreatorIfNotExists func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, error)

func newProjectCreatorIfNotExists(logger logging.SeverityLogger, findProject projectARNFinder, createProject projectCreator) projectCreatorIfNotExists {
	return func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, error) {
		logger.Info("searching to skip creating AWS Device Farm projects")
		projectARN, findErr := findProject(groupName)

		if findErr != nil {
			if findErr.NotFound != nil {
				return createProject(groupName)
			}
			return "", findErr
		}

		logger.Info("skipped to create an AWS Device Farm project (because the project already exists)")
		return projectARN, nil
	}
}
