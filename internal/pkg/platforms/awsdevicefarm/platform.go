package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/executor/awscli"
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"time"
)

var ID platforms.ID = "aws-device-farm"

var AWSDeviceFarm platforms.Platform = awsDeviceFarm{}

type awsDeviceFarm struct{}

func (awsDeviceFarm) ID() platforms.ID {
	return ID
}

func (awsDeviceFarm) AuthStatusChecker() platforms.AuthStatusChecker {
	return func(bag platforms.AuthStatusCheckerBag) error {
		awsCmd := awscli.NewExecutor(bag)
		deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

		checkAuthStatus := newAuthStatusChecker(
			awscli.NewInstallStatusGetter(bag.GetFinder()),
			awscli.NewVersionGetter(awsCmd),
			awscli.NewConfigStatusGetter(awsCmd),
			devicefarm.NewAuthorizationStatusChecker(deviceFarmCmd),
		)

		return checkAuthStatus()
	}
}

func (awsDeviceFarm) DeviceLister() platforms.DeviceLister {
	return func(bag platforms.DevicesListerBag) (entries []platforms.DeviceOrError, e error) {
		awsCmd := awscli.NewExecutor(bag)
		deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

		listDevices := newDeviceEntryLister(devicefarm.NewDeviceLister(deviceFarmCmd))

		return listDevices()
	}
}

func (awsDeviceFarm) InstanceLister() platforms.InstanceLister {
	return func(groupName platforms.InstanceGroupName, bag platforms.InstanceListerBag) ([]platforms.InstanceOrError, error) {
		awsCmd := awscli.NewExecutor(bag)
		deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

		listInstances := newInstanceLister(
			newProjectARNFinder(bag.GetLogger(), devicefarm.NewProjectLister(deviceFarmCmd)),
			newInstanceCollector(
				devicefarm.NewRunLister(deviceFarmCmd),
				devicefarm.NewJobLister(deviceFarmCmd),
			),
		)

		return listInstances(groupName)
	}
}

func (awsDeviceFarm) InstanceGroupLister() platforms.InstanceGroupLister {
	return func(bag platforms.InstanceGroupListerBag) ([]platforms.InstanceGroupListEntry, error) {
		awsCmd := awscli.NewExecutor(bag)
		deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

		listInstanceGroups := newInstanceGroupLister(devicefarm.NewProjectLister(deviceFarmCmd))

		return listInstanceGroups()
	}
}

func (awsDeviceFarm) AllInstanceLister() platforms.AllInstanceLister {
	return func(bag platforms.AllInstanceListerBag) ([]platforms.InstanceOrError, error) {
		awsCmd := awscli.NewExecutor(bag)
		deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

		listAllInstances := newAllInstanceLister(
			devicefarm.NewProjectLister(deviceFarmCmd),
			newInstanceCollector(
				devicefarm.NewRunLister(deviceFarmCmd),
				devicefarm.NewJobLister(deviceFarmCmd),
			),
		)

		return listAllInstances()
	}
}

type remoteAgentLauncherBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
	GetUploader() executor.Uploader
	GetFileOpener() executor.FileOpener
}

func (a awsDeviceFarm) IOSForever() platforms.IOSForever {
	return func(plan platforms.IOSPlan, bag platforms.IOSForeverBag) error {
		launchRemoteAgent := a.newRemoteAgentLauncher(bag)
		foreverIOS := newIOSForever(launchRemoteAgent)
		return foreverIOS(plan, bag)
	}
}

func (a awsDeviceFarm) IOSRunner() platforms.IOSRunner {
	return func(plan platforms.IOSPlan, bag platforms.IOSRunnerBag) error {
		launchRemoteAgent := a.newRemoteAgentLauncher(bag)
		waitRunResult := a.newRunResultWaiter(bag)
		runIOS := newIOSRunner(launchRemoteAgent, waitRunResult)
		return runIOS(plan, bag)
	}
}

func (a awsDeviceFarm) AndroidForever() platforms.AndroidForever {
	return func(
		plan platforms.AndroidPlan,
		bag platforms.AndroidForeverBag,
	) error {
		launchRemoteAgent := a.newRemoteAgentLauncher(bag)
		foreverAndroid := newAndroidForever(launchRemoteAgent)
		return foreverAndroid(plan, bag)
	}
}

func (a awsDeviceFarm) AndroidRunner() platforms.AndroidRunner {
	return func(
		plan platforms.AndroidPlan,
		bag platforms.AndroidRunnerBag,
	) error {
		launchRemoteAgent := a.newRemoteAgentLauncher(bag)
		waitRunResult := a.newRunResultWaiter(bag)
		runnerAndroid := newAndroidRunner(launchRemoteAgent, waitRunResult)
		return runnerAndroid(plan, bag)
	}
}

func (awsDeviceFarm) DeviceFinder() platforms.DeviceFinder {
	return func(device platforms.EitherDevice, bag platforms.DeviceFinderBag) (bool, error) {
		awsCmd := awscli.NewExecutor(bag)
		deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

		findDevice := newDeviceFinder(devicefarm.NewDeviceLister(deviceFarmCmd))

		return findDevice(device)
	}
}

func (awsDeviceFarm) Halt() platforms.Halt {
	return func(groupName platforms.InstanceGroupName, bag platforms.HaltBag) (platforms.Results, error) {
		awsCmd := awscli.NewExecutor(bag)
		deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

		haltApp := newAppHalt(
			newProjectARNFinder(bag.GetLogger(), devicefarm.NewProjectLister(deviceFarmCmd)),
			devicefarm.NewRunLister(deviceFarmCmd),
			devicefarm.NewRunStopper(deviceFarmCmd),
		)

		return haltApp(groupName)
	}
}

func (awsDeviceFarm) PlanValidator() platforms.PlanValidator {
	return newPlanValidator()
}

func (awsDeviceFarm) newRunResultWaiter(bag awscli.Bag) runResultWaiter {
	logger := bag.GetLogger()
	awsCmd := awscli.NewExecutor(bag)
	deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

	pollingInterval := time.Second
	timeout := time.Minute * 30

	return newRunResultWaiter(logger, devicefarm.NewRunGetter(deviceFarmCmd), pollingInterval, timeout)
}

func (awsDeviceFarm) newRemoteAgentLauncher(bag remoteAgentLauncherBag) remoteAgentLauncher {
	logger := bag.GetLogger()
	awsCmd := awscli.NewExecutor(bag)
	deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

	pollingInterval := time.Second
	timeout := time.Minute * 30

	reserveAndUploadIfNotExists := newReserveAndUploaderIfNotExists(
		logger,
		newUploadFinder(logger, devicefarm.NewUploadLister(deviceFarmCmd)),
		newReserveAndUploader(
			logger,
			newUploadReserving(logger, devicefarm.NewUploadCreator(deviceFarmCmd)),
			newRetryingUploader(bag.GetLogger(), bag.GetUploader(), 5),
		),
		devicefarm.NewUploadDeleter(deviceFarmCmd),
	)
	hash := platforms.NewCRC32Hasher()

	return newRemoteAgentLauncher(
		logger,
		newProjectCreatorIfNotExists(
			logger,
			newProjectARNFinder(logger, devicefarm.NewProjectLister(deviceFarmCmd)),
			newProjectCreator(logger, devicefarm.NewProjectCreator(deviceFarmCmd)),
		),
		newDeviceARNFinder(logger, devicefarm.NewDeviceLister(deviceFarmCmd)),
		newDevicePoolCreatorIfNotExists(
			logger,
			newDevicePoolARNFinder(logger, devicefarm.NewDevicePoolLister(deviceFarmCmd)),
			newDevicePoolCreator(logger, devicefarm.NewDevicePoolCreator(deviceFarmCmd)),
		),
		newAppUploader(logger, bag.GetFileOpener(), hash, reserveAndUploadIfNotExists),
		newTestPackageUploader(logger, newTestPackageGen(), hash, reserveAndUploadIfNotExists),
		newTestSpecUploader(logger, hash, reserveAndUploadIfNotExists),
		newRunScheduler(logger, devicefarm.NewRunScheduler(deviceFarmCmd)),
		newUploadWaiter(
			logger,
			executor.NewWaiter(),
			devicefarm.NewUploadGetter(deviceFarmCmd),
			pollingInterval,
			timeout,
		),
	)
}
