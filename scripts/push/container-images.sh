#!/usr/bin/env bash

VERSION=$1
REPO=$2 # like: quay.io/leonardpahlke

echo
echo "Start building and pushing the version: $VERSION of the golang server containers with 'ko' to $REPO/whopper-..."
echo

servers=(
    analyzer
    discoverer
    hub
    parser
    translator
)

declare -i i=1
for SV in "${servers[@]}"; do
    echo "🕐 $i Build and push container image $SV"
    KO_DOCKER_REPO=$REPO/whopper-$SV ko publish ./cmd/server/$SV --bare -t $VERSION
    i+=1
done

echo "DONE! ✅"
echo