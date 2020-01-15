package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type uploadReserving func(arn devicefarm.ProjectARN, uploadProp uploadProperty) (devicefarm.Upload, error)

func newUploadReserving(logger logging.SeverityLogger, createUpload devicefarm.UploadCreator) uploadReserving {
	return func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty) (devicefarm.Upload, error) {
		logger.Info(fmt.Sprintf("reserving AWS Device Farm upload: %q", uploadProp.fileName))

		upload, err := createUpload(projectARN, uploadProp.fileName, uploadProp.uploadType)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to reserve upload: %s", err.Error()))
			return devicefarm.Upload{}, err
		}

		logger.Info(fmt.Sprintf("successfully reserved upload: %q", uploadProp.fileName))
		return upload, nil
	}
}
