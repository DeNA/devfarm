#!/bin/bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "$0")"; pwd)"

(cd "$THIS_DIR"
	set -x

	./usage/devfarm/auth-status
	./usage/devfarm/cat-asset
	./usage/devfarm/forever-all
	./usage/devfarm/forever-android
	./usage/devfarm/forever-ios
	./usage/devfarm/halt
	./usage/devfarm/init
	./usage/devfarm/list-devices
	./usage/devfarm/ls-assets
	./usage/devfarm/run-all
	./usage/devfarm/run-android
	./usage/devfarm/run-ios
	./usage/devfarm/status
	./usage/devfarm/validate

	./usage/devfarmagent/aws-device-farm/args
	./usage/devfarmagent/aws-device-farm/forever
	./usage/devfarmagent/aws-device-farm/run
)
