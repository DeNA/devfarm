package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"time"
)

type uploadWaiter func(arn devicefarm.UploadARN) error

func newUploadWaiter(logger logging.SeverityLogger, wait executor.Waiter, getUpload devicefarm.UploadGetter, interval time.Duration, timeout time.Duration) uploadWaiter {
	return func(uploadARN devicefarm.UploadARN) error {
		cond := func() (bool, error) {
			upload, err := getUpload(uploadARN)
			if err != nil {
				return false, err
			}

			switch upload.Status {
			case devicefarm.UploadStatusIsInitialized, devicefarm.UploadStatusIsProcessing:
				// NOTE: Continue to wait.
				return true, nil

			case devicefarm.UploadStatusIsSucceeded:
				logger.Info("upload seems ready")
				return false, nil

			case devicefarm.UploadStatusIsFailed:
				logger.Error(fmt.Sprintf("upload seems failed: %s", upload.Metadata))
				return false, fmt.Errorf("upload seems invalid: %s", upload)

			default:
				return false, fmt.Errorf("unknown upload status: %v", upload)
			}
		}

		logger.Info("waiting upload ready")
		return wait(cond, "upload status", interval, timeout)
	}
}
