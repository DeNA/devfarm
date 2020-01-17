package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/platforms"
)

func anySuccessfulAppUploader() appUploader {
	return stubAppUploader(appUpload{arn: "arn:aws:devicefarm:ANY_APP_UPLOAD"}, nil)
}

func stubAppUploader(upload appUpload, err error) appUploader {
	return func(appPath ipaOrApkPathOnLocal, osName platforms.OSName, projectARN devicefarm.ProjectARN) (appUpload, error) {
		return upload, err
	}
}
