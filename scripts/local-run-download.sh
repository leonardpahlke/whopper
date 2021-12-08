#!/usr/bin/env bash

echo "\nstart dapr\n"
dapr init

echo "\nrun downloader server\n"
dapr run --app-id downloader --app-protocol grpc --app-port 50051 --config ./configs/downloader.yaml -- go run ./cmd/server/downloader/main.go