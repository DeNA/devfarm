package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"io"
	"os"
	"sync"
)

type appUpload struct {
	arn devicefarm.UploadARN
}

type ipaOrApkPathOnLocal string

type appUploader func(appPath ipaOrApkPathOnLocal, osName platforms.OSName, projectARN devicefarm.ProjectARN) (appUpload, error)

func newAppUploader(logger logging.SeverityLogger, openFile exec.FileOpener, hash platforms.Hasher32, reserveAndUploadIfNotExists reserveAndUploaderIfNotExists) appUploader {
	return func(appPath ipaOrApkPathOnLocal, osName platforms.OSName, projectARN devicefarm.ProjectARN) (appUpload, error) {
		logger.Info(fmt.Sprintf("validating the app to upload to AWS Device Farm: %q", appPath))

		appPathString := string(appPath)

		appName, extname, appNameErr := validateAppName(appPath)
		if appNameErr != nil {
			logger.Error(fmt.Sprintf("invalid app name: %s", appNameErr.Error()))
			return appUpload{}, appNameErr
		}

		var uploadType devicefarm.UploadType
		switch osName {
		case platforms.OSIsIOS:
			uploadType = devicefarm.UploadTypeIsIOSApp
		case platforms.OSIsAndroid:
			uploadType = devicefarm.UploadTypeIsAndroidApp
		default:
			return appUpload{}, fmt.Errorf("not supported OS: %s", osName)
		}

		file, openErr := openFile(appPathString, os.O_RDONLY, 0)
		if openErr != nil {
			logger.Debug(fmt.Sprintf("failed to open %q: %s", appPathString, openErr.Error()))
			return appUpload{}, openErr
		}
		defer file.Close()

		fi, statErr := file.Stat()
		if statErr != nil {
			return appUpload{}, statErr
		}

		hashValue, hashErr := hash(file)
		if hashErr != nil {
			logger.Debug(fmt.Sprintf("cannot get hash from %q: %s", appPathString, hashErr.Error()))
			return appUpload{}, hashErr
		}

		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return appUpload{}, err
		}

		// NOTE: CRC32 may not enough to identify the app revision...
		//       But we should keep enough available name length for users.
		uploadName := fmt.Sprintf("%s%s-%08x%s", devfarmUploadNamePrefix, appName, hashValue, extname)
		uploadProp := newUploadProperty(devicefarm.UploadFileName(uploadName), uploadType)

		uploadARN, uploadErr := reserveAndUploadIfNotExists(projectARN, uploadProp, file, fi.Size())
		if uploadErr != nil {
			return appUpload{}, uploadErr
		}

		return appUpload{uploadARN}, nil
	}
}

func newAppUploaderCached(uploadApp appUploader) appUploader {
	var mu sync.Mutex
	cache := make(map[appPathAndOSNameAndProjectARN]appUpload)

	return func(appPath ipaOrApkPathOnLocal, osName platforms.OSName, projectARN devicefarm.ProjectARN) (appUpload, error) {
		mu.Lock()
		defer mu.Unlock()

		key := appPathAndOSNameAndProjectARN{
			appPath:    appPath,
			osName:     osName,
			projectARN: projectARN,
		}

		if cached, ok := cache[key]; ok {
			return cached, nil
		}

		uploadARN, err := uploadApp(appPath, osName, projectARN)
		if err != nil {
			return appUpload{}, err
		}

		cache[key] = uploadARN
		return uploadARN, nil
	}
}

type appPathAndOSNameAndProjectARN struct {
	appPath    ipaOrApkPathOnLocal
	osName     platforms.OSName
	projectARN devicefarm.ProjectARN
}
