package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"io"
)

func anySuccessfulReserveAndUploaderIfNotExists() reserveAndUploaderIfNotExists {
	return stubReserveAndUploaderIfNotExists("arn:aws:devicefarm:ANY_UPLOAD", nil)
}

func stubReserveAndUploaderIfNotExists(uploadARN devicefarm.UploadARN, err error) reserveAndUploaderIfNotExists {
	return func(devicefarm.ProjectARN, uploadProperty, io.Reader, int64) (devicefarm.UploadARN, error) {
		return uploadARN, err
	}
}

type uploaderCallArgs struct {
	projectARN devicefarm.ProjectARN
	uploadProp uploadProperty
	reader     io.Reader
	size       int64
}

func spyReserveUploaderIfNotExists(inherited reserveAndUploaderIfNotExists) (*[]uploaderCallArgs, reserveAndUploaderIfNotExists) {
	callArgs := make([]uploaderCallArgs, 0)

	return &callArgs, func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty, reader io.Reader, size int64) (devicefarm.UploadARN, error) {
		callArgs = append(callArgs, uploaderCallArgs{projectARN, uploadProp, reader, size})
		return inherited(projectARN, uploadProp, reader, size)
	}
}
