#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# This script is used to check if all required tools have been installed
cmds=(
    "git --version"
    "pulumi version"
    "node -v"
    "gcloud -v"
    "go version"
    "task --version"
    "protoc --version"
)

declare -i i=1

for CMD in "${cmds[@]}"; do
    CMD_OUT=$CMD
    echo -ne "Check installation $i/${#cmds[@]}\r"
    i+=1
done

echo
echo "All checks ok"
echo "note: make sure to install 'protoc-gen-go' as well"