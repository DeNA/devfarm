#!/bin/bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "$0")"; pwd)"

throw () {
	local message="$*"
	echo -e "$message" 1>&2
	false
}

os_arch="$(uname -m) $(uname)"
case "$os_arch" in
	"x86_64 Linux")
		"${THIS_DIR}/linux-amd64/devfarmagent" "$@" ;; 

	"x86_64 Darwin")
		"${THIS_DIR}/darwin-amd64/devfarmagent" "$@" ;; 

	*)
		throw "this os and arch is not supported: $os_arch"
esac
