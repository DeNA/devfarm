#!/bin/bash
set -euo pipefail

# NOTE: DO NOT DEFINE ENV VARS TO SHARE DATA ACROSS WORKFLOW STEPS IN THIS DIR. IT MUST NOT WORK.
#       Instead, you should do it on ./0-shared.bash

# NOTE: Implicitly loaded functions from 0-shared.bash


test_ios() {
	echo "Arguments"
	devfarmagent aws-device-farm args decode "$DEVFARM_AGENT_APP_ARGS_ENCODED"

	devfarmagent aws-device-farm "$DEVFARM_AGENT_SUB_CMD" --verbose | tee "${DEVICEFARM_LOG_DIR}/devfarmagent.log"
}


test_android() {
	echo "Arguments"
	devfarmagent aws-device-farm args decode "$DEVFARM_AGENT_APP_ARGS_ENCODED"

	devfarmagent aws-device-farm "$DEVFARM_AGENT_SUB_CMD" --verbose | tee "${DEVICEFARM_LOG_DIR}/devfarmagent.log"
}
        

case "$DEVICEFARM_DEVICE_PLATFORM_NAME" in
	"iOS")
		echo "Launch iOS apps (it takes about 10 minutes)"
		test_ios
		;;

	"Android")
		echo "Launch Android apps (it takes about 10 minutes)"
		test_android
		;;

	*)
		throw "unsupported platform: $DEVICEFARM_DEVICE_PLATFORM_NAME"
		;;
esac
