package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"math"
	"time"
)

type testSpecEmbeddedData interface {
	androidAppID() (platforms.AndroidAppID, bool)
	args() TransportableArgs
	lifetime() time.Duration
	remoteAgentSubCmd() remoteAgentSubCmd
}

type remoteAgentSubCmd string

const (
	remoteAgentSubCmdIsRun     remoteAgentSubCmd = "run"
	remoteAgentSubCmdIsForever remoteAgentSubCmd = "forever"
)

type customTestEnvSpec string

func generateCustomTestEnvSpec(embeddedData testSpecEmbeddedData) (customTestEnvSpec, error) {
	appArgsEncoded, appArgsErr := EncodeAppArgs(embeddedData.args())
	if appArgsErr != nil {
		return "", appArgsErr
	}

	androidAppIDOrEmpty, _ := embeddedData.androidAppID()
	lifetimeSec := int(math.Ceil(embeddedData.lifetime().Seconds()))
	agentSubCmd := embeddedData.remoteAgentSubCmd()

	// This is "Default TestSpec for Android Appium Node.js".
	yaml := fmt.Sprintf(`
version: 0.1

# This flag enables your test to run using Device Farm's Amazon Linux 2 test host. For more information,
# please see https://docs.aws.amazon.com/devicefarm/latest/developerguide/amazon-linux-2.html
android_test_host: amazon_linux_2

# Phases are collection of commands that get executed on Device Farm.
phases:
  # The install phase includes commands that install dependencies that your tests use.
  # Default dependencies for testing frameworks supported on Device Farm are already installed.
  install:
    commands:
      - echo "Test Environments"
      - uname -a
      - printenv
      - ls -la "$DEVICEFARM_TEST_PACKAGE_PATH"

      - echo "Export environment vars that shared across phases"
      - export DEVFARM_AGENT_ANDROID_APP_ID_IF_ANDROID=%q
      - export DEVFARM_AGENT_APP_ARGS_ENCODED=%q
      - export DEVFARM_AGENT_LIFETIME_SEC=%d
      - export DEVFARM_AGENT_SUB_CMD=%s

      - source $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/0-shared.bash

      - echo "Run $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/1-install.bash"
      - (source $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/1-install.bash)


  # The pre-test phase includes commands that setup your test environment.
  pre_test:
    commands:
      - echo "Run $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/2-pretest.bash"
      - (source $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/2-pretest.bash)


  # The test phase includes commands that run your test suite execution.
  test:
    commands:
      - echo "Run $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/3-test.bash"
      - (source $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/3-test.bash)


  # The post test phase includes are commands that are run after your tests are executed.
  post_test:
    commands:
      - echo "Run $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/4-posttest.bash"
      - (source $DEVICEFARM_TEST_PACKAGE_PATH/aws-device-farm/workflows/4-posttest.bash)

# The artifacts phase lets you specify the location where your tests logs, device logs will be stored.
# And also let you specify the location of your test logs and artifacts which you want to be collected by Device Farm.
# These logs and artifacts will be available through ListArtifacts API in Device Farm.
artifacts:
  # By default, Device Farm will collect your artifacts from following directories
  - $DEVICEFARM_LOG_DIR
`, androidAppIDOrEmpty, appArgsEncoded, lifetimeSec, agentSubCmd)

	return customTestEnvSpec(yaml), nil
}
