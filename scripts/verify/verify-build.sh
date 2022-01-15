#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PLATFORMS=(
    linux/amd64
    linux/386
    linux/arm
    linux/arm64
    linux/ppc64le
    linux/s390x
    windows/amd64
    windows/386
    freebsd/amd64
    darwin/amd64
)

for PLATFORM in "${PLATFORMS[@]}"; do
    OS="${PLATFORM%/*}"
    ARCH=$(basename "$PLATFORM")

    echo "Building project for $PLATFORM"
    GOARCH="$ARCH" GOOS="$OS" go build ./...
done