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
    "kubectl version --client"
)

declare -i i=1
for CMD in "${cmds[@]}"; do
    CMD_OUT=$CMD
    echo -ne "Check installation $i/${#cmds[@]}\r"
    i+=1
done
echo
echo "Tools installed ✅"
echo

# Check if all environment variables are set
envs=(
    "GOPATH"
    "PULUMI_ACCESS_TOKEN"
    "DEFAULT_GOOGLE_PROJECT"
)

declare -i i=1
for CMD in "${envs[@]}"; do
    [ -z "$CMD" ] && exit 1
    echo -ne "Check if environment is set variable $i/${#envs[@]}\r"
    i+=1
done
echo
echo "Environment variables set ✅"

echo
echo "All checks ok ✅"
echo "note: make sure to install 'protoc-gen-go' as well"