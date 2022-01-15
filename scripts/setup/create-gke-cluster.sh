#!/usr/bin/env bash

# This script is used to create a GKE service.
# The kubernetes cluster is used to deploy the pulumi operator to which handles additional resource deployments.

set -o errexit
#set -o nounset
set -o pipefail

arg1=$1
arg2=$2
arg3=$3
arg4=$4

echo "Source config file"
. ../../configs/cluster.config

clusterName="${arg1:-whopper-infra-cluster}"
initialNodeCount="${arg2:-2}"
machineType="${arg3:-n1-standard-1}"
region="${arg4:-europe-west3}"

echo "Enable GCP services required to deploy a GKE cluster"
SERVICES=(
    container.googleapis.com
)
for SERVICE in "${SERVICES[@]}"; do
    echo "enable service $SERVICE"gcloud services enable
    gcloud services enable $SERVICE
done

echo
echo "Deploy GKE cluster"
echo "configuration:"
echo "- arg1 clusterName: $clusterName"
echo "- arg2 initialNodeCount: $initialNodeCount"
echo "- arg3 machineType: $machineType"
echo "- arg4 region: $region"
echo

gcloud container clusters create $clusterName \
    --num-nodes=$initialNodeCount \
    --region=$region \
    --machine-type=$machineType \
    --no-enable-ip-alias

echo "run 'cloud container clusters delete $clusterName --region $region' to delete the cluster again"