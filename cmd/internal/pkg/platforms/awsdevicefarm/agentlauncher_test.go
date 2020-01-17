package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"testing"
)

func TestNewRemoteAgentLauncher(t *testing.T) {
	launchAgent := NewRemoteAgentLauncher(
		logging.NullSeverityLogger(),
		anySuccessfulProjectCreator(),
		anySuccessfulDeviceARNFinder(),
		anySuccessfulDevicePoolCreator(),
		anySuccessfulAppUploader(),
		anySuccessfulTestPackageUploader(),
		anySuccessfulSpecUploader(),
		anySuccessfulRunScheduler(),
		anySuccessfulUploadWaiter(),
	)

	opts := anyRemoteAgentLauncherOpts()

	if _, err := launchAgent(platforms.AnyInstanceGroup().Name, opts); err != nil {
		t.Errorf("got %v, want nil", err)
		return
	}
}
