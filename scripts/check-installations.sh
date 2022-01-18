#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# This script is used to check if all required tools have been installed

echo "Check git version"
git --version

echo "Check pulumi version"
pulumi version

echo "Check kubectl version"
kubectl version

echo "Check node version"
node -v

echo "Check gcloud version"
gcloud -v

echo "Check go version"
go version

echo "Check task version"
task --version