package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
)

type uploadFinder func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty) (devicefarm.Upload, *uploadFinderError)

type uploadFinderError struct {
	notFound    error
	unspecified error
}

func (e uploadFinderError) Error() string {
	if e.notFound != nil {
		return e.notFound.Error()
	}
	return e.unspecified.Error()
}

func newUploadFinder(logger logging.SeverityLogger, listUploads devicefarm.UploadLister) uploadFinder {
	return func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty) (devicefarm.Upload, *uploadFinderError) {
		logger.Info("listing AWS Device Farm uploads")
		uploads, uploadErr := listUploads(projectARN)
		if uploadErr != nil {
			logger.Error(fmt.Sprintf("failed to list AWS Device Farm uploads: %s", uploadErr.Error()))
			return devicefarm.Upload{}, &uploadFinderError{unspecified: uploadErr}
		}

		for _, upload := range uploads {
			if upload.Name == uploadProp.fileName && upload.Type == uploadProp.uploadType {
				logger.Info("found AWS Device Farm upload")
				logger.Debug(fmt.Sprintf("uploaded: %q at %q", upload.Name, upload.ARN))
				return upload, nil
			}
		}

		logger.Info("AWS Device Farm upload not found")
		notFound := fmt.Errorf("no such upload: (%q, %q)", uploadProp.fileName, uploadProp.uploadType)
		return devicefarm.Upload{}, &uploadFinderError{notFound: notFound}
	}
}
