package awsdevicefarm

import (
	"bytes"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

type testPackageUpload struct {
	arn devicefarm.UploadARN
}

type testPackageUploader func(projectARN devicefarm.ProjectARN) (testPackageUpload, error)

func newTestPackageUploader(
	logger logging.SeverityLogger,
	generateTestPackage testPackageGen,
	hash platforms.Hasher32,
	reserveAndUploadIfNotExists reserveAndUploaderIfNotExists,
) testPackageUploader {
	return func(projectARN devicefarm.ProjectARN) (testPackageUpload, error) {
		embedded := newTestPackageEmbeddedData()

		testPackage := &bytes.Buffer{}

		logger.Info("generating AWS Device Farm test package")
		if err := generateTestPackage(embedded, testPackage); err != nil {
			logger.Error(fmt.Sprintf("failure to generate the test package: %s", err.Error()))
			return testPackageUpload{}, err
		}

		hashValue, hashErr := hash(bytes.NewReader(testPackage.Bytes()))
		if hashErr != nil {
			logger.Error(fmt.Sprintf("failed to get a hash from the test package: %s", hashErr.Error()))
			return testPackageUpload{}, hashErr
		}

		filename := fmt.Sprintf("%stest-package-%08x.zip", devfarmUploadNamePrefix, hashValue)
		uploadProp := newUploadProperty(devicefarm.UploadFileName(filename), devicefarm.UploadTypeIsAppiumNodeTestPackage)

		uploadARN, uploadErr := reserveAndUploadIfNotExists(projectARN, uploadProp, testPackage, int64(testPackage.Len()))
		if uploadErr != nil {
			return testPackageUpload{}, uploadErr
		}

		return testPackageUpload{arn: uploadARN}, nil
	}
}

func newTestPackageUploaderCached(uploadTestPackage testPackageUploader) testPackageUploader {
	var mu sync.Mutex
	cache := make(map[devicefarm.ProjectARN]testPackageUpload)

	return func(projectARN devicefarm.ProjectARN) (testPackageUpload, error) {
		mu.Lock()
		defer mu.Unlock()

		if cached, ok := cache[projectARN]; ok {
			return cached, nil
		}

		testPackage, err := uploadTestPackage(projectARN)
		if err != nil {
			return testPackageUpload{}, err
		}

		cache[projectARN] = testPackage
		return testPackage, nil
	}
}
