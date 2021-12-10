.DEFAULT_GOAL:=help
SHELL:=/usr/bin/env bash

COLOR:=\\033[36m
NOCOLOR:=\\033[0m

##@ Verify application

.PHONY: verify verify-build verify-golangci-lint verify-go-mod test-go-unit

verify: verify-build verify-golangci-lint verify-go-mod test-go-unit

verify-build:
	./scripts/verify-build.sh

verify-go-mod:
	./scripts/verify-go-mod.sh

verify-golangci-lint:
	./scripts/verify-golangci-lint.sh

##@ Tests & Verify application

.PHONY: test-go-unit
test-go-unit:
	./scripts/verify-test-go.sh

##@ Project setup and stallation

.PHONY: verify-installation
verify-installation:
	./scripts/verify-installation.sh

.PHONY: install verify-installation
install:
	./scripts/install.sh

.PHONY: compile-grpc
compile-grpc:
	protoc --go_out=pkg/ --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative api/whopper.proto

##@ Start dapr apk locally

.PHONY: start-downloader
start-downloader:
	echo "run downloader"
	dapr run --app-id downloader --app-protocol grpc --app-port 50051 --config ./configs/downloader.yaml -- go run ./cmd/server/downloader/main.go

.PHONY: start-parser
start-parser:
	echo "run parser"
	echo "TODO: not yet implemented"

.PHONY: start-translator
start-translator:
	echo "run translator"
	echo "TODO: not yet implemented"

.PHONY: start-analyzer
start-analyzer:
	echo "run analyzer"
	echo "TODO: not yet implemented"

.PHONY: start-hub
start-hub:
	echo "run hub"
	echo "TODO: not yet implemented"

.PHONY: start-discoverer
start-discoverer:
	echo "run discoverer"
	dapr run --app-id discoverer --app-protocol grpc --app-port 50055 --config ./configs/discoverer.yaml -- go run ./cmd/server/discoverer/main.go


##@ Update golang dependencies

.PHONY: update-deps-go
update-deps-go: GO111MODULE=on
update-deps-go:
	go get -u -t ./...
	go mod tidy
	go mod verify