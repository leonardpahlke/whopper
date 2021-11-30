#!/usr/bin/env bash

echo "\ncompile protobuf file\n"
protoc --go_out=pkg/ --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative api/whopper.proto