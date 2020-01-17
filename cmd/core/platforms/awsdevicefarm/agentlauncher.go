package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/dena/devfarm/cmd/core/platforms"
)

type RemoteAgentIntermediates struct {
	projectARN             devicefarm.ProjectARN
	projectARNOk           bool
	deviceARN              devicefarm.DeviceARN
	deviceARNOk            bool
	devicePoolARN          devicefarm.DevicePoolARN
	devicePoolARNOk        bool
	testSpecUploadARN      devicefarm.UploadARN
	testSpecUploadARNOk    bool
	testPackageUploadARN   devicefarm.UploadARN
	testPackageUploadARNOk bool
	appUploadARN           devicefarm.UploadARN
	appUploadARNOk         bool
	runARN                 devicefarm.RunARN
	runARNOk               bool
}

type RemoteAgentLauncher func(platforms.InstanceGroupName, remoteAgentLauncherOpts) (RemoteAgentIntermediates, error)

func NewRemoteAgentLauncher(
	logger logging.SeverityLogger,
	createProject projectCreator,
	findDeviceARN deviceARNFinder,
	createDevicePool devicePoolCreator,
	uploadApp appUploader,
	uploadTestPackage testPackageUploader,
	uploadTestSpec testSpecUploader,
	scheduleRun runScheduler,
	waitUntilUploadIsCompleted uploadWaiter,
) RemoteAgentLauncher {
	return func(groupName platforms.InstanceGroupName, opts remoteAgentLauncherOpts) (intermediates RemoteAgentIntermediates, err error) {
		projectARN, projectErr := createProject(groupName)
		if projectErr != nil {
			err = projectErr
			return
		}
		intermediates.projectARN = projectARN
		intermediates.projectARNOk = true

		iosOrAndroid := opts.iosOrAndroidDevice()
		deviceARN, deviceErr := findDeviceARN(iosOrAndroid)
		if deviceErr != nil {
			err = deviceErr
			return
		}
		intermediates.deviceARN = deviceARN
		intermediates.deviceARNOk = true

		devicePoolARN, devicePoolErr := createDevicePool(projectARN, deviceARN)
		if devicePoolErr != nil {
			err = devicePoolErr
			return
		}
		intermediates.devicePoolARN = devicePoolARN
		intermediates.devicePoolARNOk = true

		logger.Info("generating AWS Device Farm custom test spec")
		testSpec, testSpecErr := generateCustomTestEnvSpec(opts)
		if testSpecErr != nil {
			logger.Error(fmt.Sprintf("failed to generate the AWS Device Farm custom test spec: %s", testSpecErr.Error()))
			err = testSpecErr
			return
		}
		logger.Info("AWS Device Farm custom test spec was successfully generated")

		testSpecUploaded, specErr := uploadTestSpec(projectARN, testSpec)
		if specErr != nil {
			err = specErr
			return
		}
		intermediates.testSpecUploadARN = testSpecUploaded.arn
		intermediates.testSpecUploadARNOk = true

		testPackageUploaded, pkgErr := uploadTestPackage(projectARN)
		if pkgErr != nil {
			err = pkgErr
			return
		}
		intermediates.testPackageUploadARN = testPackageUploaded.arn
		intermediates.testPackageUploadARNOk = true

		appUploaded, appErr := uploadApp(opts.ipaOrApkPath(), iosOrAndroid.OSName, projectARN)
		if appErr != nil {
			err = appErr
			return
		}
		intermediates.appUploadARN = appUploaded.arn
		intermediates.appUploadARNOk = true

		if waitErr := waitUntilUploadIsCompleted(testSpecUploaded.arn); waitErr != nil {
			err = waitErr
			return
		}

		if waitErr := waitUntilUploadIsCompleted(testPackageUploaded.arn); waitErr != nil {
			err = waitErr
			return
		}

		if waitErr := waitUntilUploadIsCompleted(appUploaded.arn); waitErr != nil {
			err = waitErr
			return
		}

		runARN, runErr := scheduleRun(
			iosOrAndroid.OSName,
			projectARN,
			devicePoolARN,
			appUploaded,
			testSpecUploaded,
			testPackageUploaded,
		)
		if runErr != nil {
			err = runErr
			return
		}
		intermediates.runARN = runARN
		intermediates.runARNOk = true

		return
	}
}
