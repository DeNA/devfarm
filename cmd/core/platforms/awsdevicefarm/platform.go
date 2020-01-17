package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/exec/awscli"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
	"time"
)

var ID platforms.ID = "aws-device-farm"

type awsDeviceFarm struct {
	authStatusChecker   platforms.AuthStatusChecker
	deviceLister        platforms.DeviceLister
	planValidator       platforms.PlanValidator
	instanceLister      platforms.InstanceLister
	instanceGroupLister platforms.InstanceGroupLister
	allInstancesLister  platforms.AllInstanceLister
	iosForever          platforms.IOSForever
	iosRunner           platforms.IOSRunner
	androidForever      platforms.AndroidForever
	androidRunner       platforms.AndroidRunner
	deviceFinder        platforms.DeviceFinder
	appHalter           platforms.Halt
}

var pollingInterval = time.Second
var timeout = time.Minute * 30
var uploadRetryCount = 5
var xxxAndroidTestErrorRetryCount = 5

var _ platforms.PlatformFactory = NewPlatform

func NewPlatform(bag platforms.Bag) platforms.Platform {
	logger := bag.GetLogger()
	uploader := bag.GetUploader()
	fileOpener := bag.GetFileOpener()
	executor := bag.GetExecutor()
	executableFinder := bag.GetFinder()
	hash := platforms.NewCRC32Hasher()

	awsCmd := awscli.NewExecutor(executor)
	deviceFarmCmd := devicefarm.NewExecutor(awsCmd)

	reserveAndUploadIfNotExists := newReserveAndUploaderIfNotExists(
		logger,
		newUploadFinder(logger, devicefarm.NewUploadLister(deviceFarmCmd)),
		newReserveAndUploader(
			logger,
			newUploadReserving(logger, devicefarm.NewUploadCreator(deviceFarmCmd)),
			newRetryingUploader(logger, uploader, uploadRetryCount),
		),
		devicefarm.NewUploadDeleter(deviceFarmCmd),
	)
	runResultWaiter := NewRunResultWaiter(logger, devicefarm.NewRunGetter(deviceFarmCmd), pollingInterval, timeout)
	remoteAgentLauncher := NewRemoteAgentLauncher(
		logger,
		newProjectCreatorCached(newProjectCreatorIfNotExists(
			logger,
			newProjectARNFinder(logger, devicefarm.NewProjectLister(deviceFarmCmd)),
			newProjectCreator(logger, devicefarm.NewProjectCreator(deviceFarmCmd)),
		)),
		newDeviceARNFinderCached(newDeviceARNFinder(logger, devicefarm.NewDeviceLister(deviceFarmCmd))),
		newDevicePoolCreatorCached(newDevicePoolCreatorIfNotExists(
			logger,
			newDevicePoolARNFinder(logger, devicefarm.NewDevicePoolLister(deviceFarmCmd)),
			newDevicePoolCreator(logger, devicefarm.NewDevicePoolCreator(deviceFarmCmd)),
		)),
		newAppUploaderCached(newAppUploader(logger, fileOpener, hash, reserveAndUploadIfNotExists)),
		newTestPackageUploaderCached(newTestPackageUploader(logger, newTestPackageGen(), hash, reserveAndUploadIfNotExists)),
		newTestSpecUploader(logger, hash, reserveAndUploadIfNotExists),
		newRunScheduler(logger, devicefarm.NewRunScheduler(deviceFarmCmd)),
		newUploadWaiter(
			logger,
			exec.NewWaiter(),
			devicefarm.NewUploadGetter(deviceFarmCmd),
			pollingInterval,
			timeout,
		),
	)

	return awsDeviceFarm{
		authStatusChecker: newAuthStatusChecker(
			awscli.NewInstallStatusGetter(executableFinder),
			awscli.NewVersionGetter(awsCmd),
			awscli.NewConfigStatusGetter(awsCmd),
			devicefarm.NewAuthorizationStatusChecker(deviceFarmCmd),
		),
		deviceLister:  newDeviceEntryLister(devicefarm.NewDeviceLister(deviceFarmCmd)),
		planValidator: newPlanValidator(),
		instanceLister: newInstanceLister(
			newProjectARNFinder(logger, devicefarm.NewProjectLister(deviceFarmCmd)),
			newInstanceCollector(
				devicefarm.NewRunLister(deviceFarmCmd),
				devicefarm.NewJobLister(deviceFarmCmd),
			),
		),
		instanceGroupLister: newInstanceGroupLister(devicefarm.NewProjectLister(deviceFarmCmd)),
		allInstancesLister: newAllInstanceLister(
			devicefarm.NewProjectLister(deviceFarmCmd),
			newInstanceCollector(
				devicefarm.NewRunLister(deviceFarmCmd),
				devicefarm.NewJobLister(deviceFarmCmd),
			),
		),
		iosForever:     newIOSForever(remoteAgentLauncher),
		iosRunner:      newIOSRunner(remoteAgentLauncher, runResultWaiter),
		androidForever: newAndroidForever(remoteAgentLauncher),
		androidRunner: newAndroidRunnerWithRetry(
			logger,
			remoteAgentLauncher,
			runResultWaiter,
			xxxAndroidTestErrorRetryCount,
		),
		deviceFinder: newDeviceFinder(devicefarm.NewDeviceLister(deviceFarmCmd)),
		appHalter: newAppHalt(
			newProjectARNFinder(logger, devicefarm.NewProjectLister(deviceFarmCmd)),
			devicefarm.NewRunLister(deviceFarmCmd),
			devicefarm.NewRunStopper(deviceFarmCmd),
		),
	}
}

var _ platforms.Platform = awsDeviceFarm{}

func (awsDeviceFarm) ID() platforms.ID {
	return ID
}

func (a awsDeviceFarm) AuthStatusChecker() platforms.AuthStatusChecker {
	return a.authStatusChecker
}

func (a awsDeviceFarm) DeviceLister() platforms.DeviceLister {
	return a.deviceLister
}

func (a awsDeviceFarm) InstanceLister() platforms.InstanceLister {
	return a.instanceLister
}

func (a awsDeviceFarm) InstanceGroupLister() platforms.InstanceGroupLister {
	return a.instanceGroupLister
}

func (a awsDeviceFarm) AllInstanceLister() platforms.AllInstanceLister {
	return a.allInstancesLister
}

func (a awsDeviceFarm) Runner() platforms.Runner {
	return platforms.NewUnoptimizedRunner(a.iosRunner, a.androidRunner)
}

func (a awsDeviceFarm) Forever() platforms.Forever {
	return platforms.NewUnoptimizedForever(a.iosForever, a.androidForever)
}

func (a awsDeviceFarm) IOSForever() platforms.IOSForever {
	return a.iosForever
}

func (a awsDeviceFarm) IOSRunner() platforms.IOSRunner {
	return a.iosRunner
}

func (a awsDeviceFarm) AndroidForever() platforms.AndroidForever {
	return a.androidForever
}

func (a awsDeviceFarm) AndroidRunner() platforms.AndroidRunner {
	return a.androidRunner
}

func (a awsDeviceFarm) DeviceFinder() platforms.DeviceFinder {
	return a.deviceFinder
}

func (a awsDeviceFarm) Halt() platforms.Halt {
	return a.appHalter
}

func (a awsDeviceFarm) PlanValidator() platforms.PlanValidator {
	return a.planValidator
}
