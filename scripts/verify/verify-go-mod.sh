#!/usr/bin/env bash
# TODO: Not used can be removed

set -o errexit
set -o nounset
set -o pipefail

go mod tidy
go mod verify
echo "DONE verify go mod!"