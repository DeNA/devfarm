#!/bin/bash
set -euo pipefail

# NOTE: DO NOT DEFINE ENV VARS TO SHARE DATA ACROSS WORKFLOW STEPS IN THIS DIR. IT MUST NOT WORK.
#       Instead, you should do it on ./0-shared.bash

# NOTE: Implicitly loaded functions from 0-shared.bash


install_ios() {
	echo "Xcode Version"
	xcodebuild -version
	echo

	echo "Selected Xcode"
	local xcode_path="$(xcode-select -p)"
	echo "$xcode_path"
	echo

	echo "Xcode DeviceSupport Status (files used by ios-deploy)"
	(find "$xcode_path/Platforms/iPhoneOS.platform/DeviceSupport" -name 'DeveloperDiskImage.dmg' -print0 | xargs -0 ls -l) || true
	(find "$HOME/Library/Developer/Xcode/iOS DeviceSupport" -name 'Symbols' -print0 | xargs -0 ls -ld) || true
	echo

	echo "Other Xcode Versions"
	find /Applications -name 'Xcode*.app' -type d -maxdepth 1 -mindepth 1
	echo

	echo "Install ios-deploy into $DEVFARM_AGENT_IOS_DEPLOY_ROOT"
	(cd "$DEVFARM_AGENT_IOS_DEPLOY_ROOT"
		npm install
	)

	echo "Verify ios-deploy"
	ios-deploy --version
	echo

	echo "Verify devfarmagent"
	devfarmagent version
	echo
}


install_android() {
	echo "Verify adb"
	adb version

	echo "Verify devfarmagent"
	devfarmagent version
}

        
case "$DEVICEFARM_DEVICE_PLATFORM_NAME" in
	"iOS")
		echo "Install dependencies for iOS"
		install_ios
		;;

	"Android")
		echo "Install dependencies for Android"
		install_android
		;;

	*)
		throw "unsupported platform: $DEVICEFARM_DEVICE_PLATFORM_NAME"
		;;
esac
