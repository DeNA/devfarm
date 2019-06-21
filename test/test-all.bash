#!/bin/bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "$0")/.."; pwd)"

(cd "$THIS_DIR"
	golangci-lint run

	go test ./...

	./test/aws-device-farm/test-workflow.bash
	./test/test-usages.bash
)
