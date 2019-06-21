#!/bin/bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "$0")/../.."; pwd)"

(cd "$THIS_DIR"
	./assets/aws-device-farm/workflows/dry-run
)
