#!/usr/bin/env bash

# TODO: generate golang proto compl. to pkg/... from api/...

echo "Compile protobuf file"
protoc --go_out=pkg/ --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative api/whopper.proto