package awsdevicefarm

import (
	"fmt"
	"sync"
	"time"

	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/dena/devfarm/cmd/core/platforms"
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

func newProjectCreatorIfNotExists(logger logging.SeverityLogger, findProject projectARNFinder, createProject projectCreator) projectCreator {
	return func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, error) {
		logger.Info(fmt.Sprintf("searching AWS Device Farm project: %q", groupName))

		projectARN, findErr := findProject(groupName)

		if findErr != nil {
			if findErr.NotFound != nil {
				return createProject(groupName)
			}
			return "", findErr
		}

		logger.Info(fmt.Sprintf("skipped creating an AWS Device Farm project (because the project already exists): %q", groupName))
		return projectARN, nil
	}
}

func newProjectCreatorCached(createProject projectCreator) projectCreator {
	var mu sync.Mutex
	cache := make(map[platforms.InstanceGroupName]devicefarm.ProjectARN)

	return func(groupName platforms.InstanceGroupName) (devicefarm.ProjectARN, error) {
		mu.Lock()
		defer mu.Unlock()

		if cached, ok := cache[groupName]; ok {
			return cached, nil
		}

		projectARN, err := createProject(groupName)
		if err != nil {
			return "", err
		}

		cache[groupName] = projectARN
		return projectARN, nil
	}
}
