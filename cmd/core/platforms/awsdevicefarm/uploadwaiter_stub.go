package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/testutil"
)

func stubStatefulUploadGetter() (*devicefarm.Upload, *error, devicefarm.UploadGetter) {
	upload := devicefarm.AnyUpload()
	err := testutil.AnyError

	return &upload, &err, func(devicefarm.UploadARN) (devicefarm.Upload, error) {
		return upload, err
	}
}

func anySuccessfulUploadWaiter() uploadWaiter {
	return stubImmediatelyBackUploadWaiter(nil)
}

func stubImmediatelyBackUploadWaiter(err error) uploadWaiter {
	return func(devicefarm.UploadARN) error {
		return err
	}
}
