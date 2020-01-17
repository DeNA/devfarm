package awsdevicefarm

import (
	"errors"
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"net/http"
)

type reserveAndUploader func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty, reader io.Reader, size int64) (devicefarm.UploadARN, error)

func newReserveAndUploader(logger logging.SeverityLogger, reserveUpload uploadReserving, upload retryingUploader) reserveAndUploader {
	return func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty, reader io.Reader, size int64) (devicefarm.UploadARN, error) {
		reserved, reservingErr := reserveUpload(projectARN, uploadProp)
		if reservingErr != nil {
			return "", reservingErr
		}

		modifyRequest := func(request *http.Request) {
			request.Method = "PUT"
			request.Header.Add("Content-Type", uploadProp.uploadType.MIMEType())
			// XXX: AWS Device Farm Upload server returns "501 Not Implemented" if Transfer-Encoding: chunked present.
			//      http.Client implicitly adds the header if the TransferEncodings is empty and the ContentLength < 0.
			request.ContentLength = size
		}

		logger.Info(fmt.Sprintf("uploading file to AWS Device Farm: %q", uploadProp.fileName))
		uploadErr := upload(string(reserved.URL), modifyRequest, reader)
		if uploadErr != nil {
			logger.Error(fmt.Sprintf("failed to upload: %s", uploadErr.Error()))
			return "", uploadErr
		}

		logger.Info(fmt.Sprintf("uploaded successfully: %q", uploadProp.fileName))
		logger.Debug(fmt.Sprintf("upload ARN: %q at %q", uploadProp.fileName, reserved.ARN))
		return reserved.ARN, nil
	}
}

type reserveAndUploaderIfNotExists func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty, reader io.Reader, size int64) (devicefarm.UploadARN, error)

func newReserveAndUploaderIfNotExists(
	logger logging.SeverityLogger,
	findUpload uploadFinder,
	reserveAndUpload reserveAndUploader,
	deleteUpload devicefarm.UploadDeleter,
) reserveAndUploaderIfNotExists {
	return func(projectARN devicefarm.ProjectARN, uploadProp uploadProperty, reader io.Reader, size int64) (devicefarm.UploadARN, error) {
		logger.Info("searching AWS Device Farm upload to skip upload")
		found, findErr := findUpload(projectARN, uploadProp)
		if findErr != nil {
			if findErr.notFound != nil {
				return reserveAndUpload(projectARN, uploadProp, reader, size)
			}

			logger.Error(fmt.Sprintf("failed to find AWS Device Farm upload: %s", findErr.Error()))
			return "", findErr
		}

		switch found.Status {
		case devicefarm.UploadStatusIsInitialized:
			// NOTE: The URL of the uncompleted one may be expired. But we cannot know whether the URL is expired ot not.
			//       So, it always assumes the URL is expired to keep simple.
			logger.Info("found uncompleted AWS Device Farm upload")
			logger.Debug(fmt.Sprintf("uncompleted upload ARN: %q", found.ARN))

			logger.Info("deleting uncompleted AWS Device Farm upload (because it may be expired, so possibly never completed)")
			if err := deleteUpload(found.ARN); err != nil {
				logger.Error(fmt.Sprintf("failed to delete the AWS Device Farm upload: %q", found.ARN))
				return "", err
			}
			return reserveAndUpload(projectARN, uploadProp, reader, size)

		case devicefarm.UploadStatusIsProcessing:
			logger.Info("found processing AWS Device Farm upload")
			logger.Debug(fmt.Sprintf("processing upload ARN: %q", found.ARN))
			return found.ARN, nil

		case devicefarm.UploadStatusIsSucceeded:
			logger.Info("found completed AWS Device Farm upload")
			logger.Debug(fmt.Sprintf("succeeded upload ARN: %q", found.ARN))
			return found.ARN, nil

		case devicefarm.UploadStatusIsFailed:
			logger.Info("found failed AWS Device Farm upload")

			// NOTE: The URL of the previous failed one may be expired. But we cannot know whether the URL is expired ot not.
			//       So, it always assumes the URL is expired to keep simple.
			logger.Info("deleting failed AWS Device Farm upload (because it may be just expired)")
			logger.Debug(fmt.Sprintf("failed upload ARN: %q", found.ARN))
			if err := deleteUpload(found.ARN); err != nil {
				logger.Error(fmt.Sprintf("failed to delete the AWS Device Farm upload: %q", found.ARN))
				return "", err
			}
			return reserveAndUpload(projectARN, uploadProp, reader, size)
		}

		err := fmt.Errorf("unknown upload status: %v", found.Status)
		logger.Error(fmt.Sprintf("failed to upload %q to AWS Device Farm: %s", uploadProp.fileName, err.Error()))
		return "", err
	}
}

type retryingUploader func(url string, modifyRequest func(*http.Request), reader io.Reader) error

func newRetryingUploader(logger logging.SeverityLogger, uploader exec.Uploader, maxRetryCount int) retryingUploader {
	return func(url string, modifyRequest func(*http.Request), reader io.Reader) error {
		remainedRetryCount := maxRetryCount

		for remainedRetryCount > 0 {
			if err := uploader(url, modifyRequest, reader); err != nil {
				logger.Debug(fmt.Sprintf("failure to upload (remained retry: %d/%d): %s", remainedRetryCount, maxRetryCount, err.Error()))
				remainedRetryCount--
				continue
			}
			return nil
		}

		msg := fmt.Sprintf("reached to max retry count: %d", maxRetryCount)
		logger.Error(msg)
		return errors.New(msg)
	}
}
