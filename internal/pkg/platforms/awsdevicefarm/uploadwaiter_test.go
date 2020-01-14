package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"testing"
	"time"
)

func TestNewUploadWaiter(t *testing.T) {
	t.Run("UploadStatus becomes: initialized -> processing -> succeeded", func(t *testing.T) {
		logger := logging.SpySeverityLogger()
		cond, wait := exec.StubWaiter(nil)
		upload, uploadErr, uploadGetter := stubStatefulUploadGetter()
		*upload = devicefarm.AnyUpload()
		*uploadErr = nil

		waitUpload := newUploadWaiter(logger, wait, uploadGetter, 0, time.Second)
		err := waitUpload("arn:aws:devicefarm:ANY_UPLOAD")

		if err != nil {
			t.Errorf("got %v, want nil", err)
			t.Log(logger.Logs.String())
			return
		}

		upload.Status = devicefarm.UploadStatusIsInitialized
		shouldWait1, err1 := (*cond)()

		if !shouldWait1 || err1 != nil {
			t.Errorf("got (%t, %v), want (true, nil)", shouldWait1, err1)
			t.Log(logger.Logs.String())
			return
		}

		upload.Status = devicefarm.UploadStatusIsProcessing
		shouldWait2, err2 := (*cond)()

		if !shouldWait2 || err2 != nil {
			t.Errorf("got (%t, %v), want (true, nil)", shouldWait2, err2)
			return
		}

		upload.Status = devicefarm.UploadStatusIsSucceeded
		shouldWait3, err3 := (*cond)()

		if shouldWait3 || err3 != nil {
			t.Errorf("got (%t, %v), want (false, nil)", shouldWait3, err3)
			t.Log(logger.Logs.String())
			return
		}
	})

	t.Run("UploadStatus becomes: initialized -> processing -> failed", func(t *testing.T) {
		logger := logging.SpySeverityLogger()
		cond, wait := exec.StubWaiter(nil)
		upload, uploadErr, uploadGetter := stubStatefulUploadGetter()
		*upload = devicefarm.AnyUpload()
		*uploadErr = nil

		waitUpload := newUploadWaiter(logger, wait, uploadGetter, 0, time.Second)
		err := waitUpload("arn:aws:devicefarm:ANY_UPLOAD")

		if err != nil {
			t.Errorf("got %v, want nil", err)
			t.Log(logger.Logs.String())
			return
		}

		upload.Status = devicefarm.UploadStatusIsInitialized
		shouldWait1, err1 := (*cond)()

		if !shouldWait1 || err1 != nil {
			t.Errorf("got (%t, %v), want (true, nil)", shouldWait1, err1)
			return
		}

		upload.Status = devicefarm.UploadStatusIsProcessing
		shouldWait2, err2 := (*cond)()

		if !shouldWait2 || err2 != nil {
			t.Errorf("got (%t, %v), want (true, nil)", shouldWait2, err2)
			return
		}

		upload.Status = devicefarm.UploadStatusIsFailed
		shouldWait3, err3 := (*cond)()

		if shouldWait3 || err3 == nil {
			t.Errorf("got (%t, nil), want (false, error)", shouldWait3)
			return
		}
	})
}
