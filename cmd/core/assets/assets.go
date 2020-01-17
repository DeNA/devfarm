//go:generate ../../../../scripts/build-devfarmagent

package assets

import (
	"fmt"
)

func Read(id AssetID) []byte {
	data, err := read(id)
	if err != nil {
		panic(fmt.Errorf("cannot read asset: %q\n%s", id, err.Error()))
	}
	return data
}

func read(id AssetID) ([]byte, error) {
	return Asset(string(id))
}

type AssetID string

const (
	DevfarmAgentBash                  AssetID = "assets/devfarmagent/devfarmagent.bash"
	DevfarmAgentLinuxAMD64            AssetID = "assets/devfarmagent/linux-amd64/devfarmagent"
	DevfarmAgentDarwinAMD64           AssetID = "assets/devfarmagent/darwin-amd64/devfarmagent"
	IOSDeployAgentPackageJSON         AssetID = "assets/ios-deploy-agent/package.json"
	IOSDeployAgentPackageLockJSON     AssetID = "assets/ios-deploy-agent/package-lock.json"
	AWSDeviceFarmWorkflowShared       AssetID = "assets/aws-device-farm/workflows/0-shared.bash"
	AWSDeviceFarmWorkflowInstallStep  AssetID = "assets/aws-device-farm/workflows/1-install.bash"
	AWSDeviceFarmWorkflowPreTestStep  AssetID = "assets/aws-device-farm/workflows/2-pretest.bash"
	AWSDeviceFarmWorkflowTestStep     AssetID = "assets/aws-device-farm/workflows/3-test.bash"
	AWSDeviceFarmWorkflowPostTestStep AssetID = "assets/aws-device-farm/workflows/4-posttest.bash"
)

var AllAssets = []AssetID{
	DevfarmAgentBash,
	DevfarmAgentLinuxAMD64,
	DevfarmAgentDarwinAMD64,
	IOSDeployAgentPackageJSON,
	IOSDeployAgentPackageLockJSON,
	AWSDeviceFarmWorkflowShared,
	AWSDeviceFarmWorkflowInstallStep,
	AWSDeviceFarmWorkflowPreTestStep,
	AWSDeviceFarmWorkflowTestStep,
	AWSDeviceFarmWorkflowPostTestStep,
}

func ValidateID(unsafeAssetID string) (AssetID, error) {
	for _, assetID := range AllAssets {
		if unsafeAssetID == string(assetID) {
			return assetID, nil
		}
	}
	return "", fmt.Errorf("no such an asset: %q", unsafeAssetID)
}
