#!/usr/bin/env bash

echo "\nstart dapr\n"
dapr init

echo "\nstart downloader container inside dapr\n"
dapr run --app-protocol grpc --app-id downloader --config "./configs/downloader.yaml" --app-port 50051

# echo "\nverify with client\n"