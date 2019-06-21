# NOTE: Only this file will be executed using "source" (others will be executed on other processes).
#       So you can define environment variables to share across workflow scripts.

throw() {
	local message="$*"
	echo -e "$message" 1>&2
	false
}

has() {
	local cmd="$1"
	type "$cmd" >/dev/null 2>&1
}

devfarmagent() {
	"$DEVICEFARM_TEST_PACKAGE_PATH/devfarmagent/devfarmagent.bash" "$@"
}

# NOTE: This is for manual execution. Not used in devfarmagent.
ios-deploy() {
	"$DEVFARM_AGENT_IOS_DEPLOY_BIN" "$@"
}


export DEVFARM_AGENT_UNARCHIVED_IOS_APP_PATH="$HOME/SUT.app"
export DEVFARM_AGENT_IOS_DEPLOY_ROOT="$DEVICEFARM_TEST_PACKAGE_PATH/ios-deploy-agent"
export DEVFARM_AGENT_IOS_DEPLOY_BIN="$DEVFARM_AGENT_IOS_DEPLOY_ROOT/node_modules/.bin/ios-deploy"
