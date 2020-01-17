package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"strings"
)

type testSpecUpload struct {
	arn devicefarm.UploadARN
}

type testSpecUploader func(projectARN devicefarm.ProjectARN, testSpec customTestEnvSpec) (testSpecUpload, error)

func newTestSpecUploader(logger logging.SeverityLogger, hash platforms.Hasher32, reserveAndUploadIfNotExists reserveAndUploaderIfNotExists) testSpecUploader {
	return func(projectARN devicefarm.ProjectARN, testSpec customTestEnvSpec) (testSpecUpload, error) {
		logger.Info("uploading AWS Device Farm custom test spec")
		hashValue, hashErr := hash(strings.NewReader(string(testSpec)))
		if hashErr != nil {
			logger.Error(fmt.Sprintf("failed to get a hash value from the test spec file: %s", hashErr.Error()))
			return testSpecUpload{}, hashErr
		}

		filename := fmt.Sprintf("%stest-spec-%08x.yml", devfarmUploadNamePrefix, hashValue)

		uploadARN, uploadErr := reserveAndUploadIfNotExists(
			projectARN,
			newUploadProperty(devicefarm.UploadFileName(filename), devicefarm.UploadTypeIsAppiumNodeTestSpec),
			strings.NewReader(string(testSpec)),
			int64(len(testSpec)),
		)
		if uploadErr != nil {
			return testSpecUpload{}, uploadErr
		}

		return testSpecUpload{arn: uploadARN}, nil
	}
}
