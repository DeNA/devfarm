package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
)

func anySuccessfulTestPackageUploader() testPackageUploader {
	return stubTestPackageUploader(testPackageUpload{arn: "arn:aws:devicefarm:TEST_PACKAGE_UPLOAD"}, nil)
}

func stubTestPackageUploader(upload testPackageUpload, err error) testPackageUploader {
	return func(devicefarm.ProjectARN) (testPackageUpload, error) {
		return upload, err
	}
}

func anyTestPackageEmbeddedData() testPackageEmbeddedData {
	return map[embeddedDataFilePath]embeddedDataFile{
		"any/embedded/data/1": {executable: true, data: []byte("1")},
		"any/embedded/data/2": {executable: true, data: []byte("2")},
		"any/embedded/data/3": {executable: true, data: []byte("3")},
	}
}
