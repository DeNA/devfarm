package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"testing"
)

func TestNewRemoteAgentLauncher(t *testing.T) {
	launchAgent := newRemoteAgentLauncher(
		logging.NullSeverityLogger(),
		anySuccessfulProjectCreatorSkipIfExists(),
		anySuccessfulDeviceARNFinder(),
		anySuccessfulDevicePoolCreatorIfNotExists(),
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
