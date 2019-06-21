#!/bin/bash
set -euo pipefail

# NOTE: DO NOT DEFINE ENV VARS TO SHARE DATA ACROSS WORKFLOW STEPS IN THIS DIR. IT MUST NOT WORK.
#       Instead, you should do it on ./0-shared.bash

# NOTE: Implicitly loaded functions from 0-shared.bash

echo "Do nothing."
