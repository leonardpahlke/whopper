#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go mod tidy
go mod verify
echo "DONE verify go mod!"