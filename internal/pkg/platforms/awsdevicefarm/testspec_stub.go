package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"time"
)

func anySuccessfulSpecUploader() testSpecUploader {
	return stubSpecUploader(testSpecUpload{arn: "arn:aws:devicefarm:ANY_TEST_SPEC_UPLOAD"}, nil)
}

func stubSpecUploader(upload testSpecUpload, err error) testSpecUploader {
	return func(devicefarm.ProjectARN, customTestEnvSpec) (testSpecUpload, error) {
		return upload, err
	}
}

type anyTestSpecEmbeddedData struct{}

var _ testSpecEmbeddedData = anyTestSpecEmbeddedData{}

func (s anyTestSpecEmbeddedData) androidAppID() (platforms.AndroidAppID, bool) {
	return "", false
}

func (s anyTestSpecEmbeddedData) args() TransportableArgs {
	return AnyTransportableArgs()
}

func (s anyTestSpecEmbeddedData) lifetime() time.Duration {
	return 0
}

func (s anyTestSpecEmbeddedData) remoteAgentSubCmd() remoteAgentSubCmd {
	return remoteAgentSubCmdIsRun
}

func anyTestSpec() customTestEnvSpec {
	return customTestEnvSpec(`["any", "test", "spec", "yaml"]`)
}
