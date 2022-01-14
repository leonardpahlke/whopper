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

.PHONY: verify-eslint
verify-eslint:
	npm run format
	npm run lint -- --fix

##@ Tests & Verify application

.PHONY: test-go-unit
test-go-unit:
	./scripts/verify-test-go.sh

##@ Project setup and installation

.PHONY: install-verify
install-verify:
	./scripts/install-verify.sh

.PHONY: install verify-installation
install:
	./scripts/install.sh

.PHONY: compile-grpc
compile-grpc:
	protoc -I=. -I=$(GOPATH)/googleapis/ --go_out=pkg/ --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative api/whopper.proto

.PHONY: init-setup-pulumi
init-setup-pulumi:
	echo "This setup has only to be run once per environment (prd, dev) after that an automation should handle updates!"
	echo "Enter pulumi access token which will be set as secret"
	read accessToken
	echo "Enter git repository URL like 'https://github.com/xyzuser/repo'"
	read projectRepo
	echo "Enter commit SHA which should be deployed like 'b19759220f25476605620fdfffeface39a630246'"
	read commitSHA
	stackName="whopper-infra"
	echo "stack name set to $stackName"
	echo "Enter which deployment environment like 'dev, prd'"
	read deployEnv
	echo "Set configuratgion..."
	pulumi config set --secret pulumiAccessToken $accessToken
	pulumi config set stackProjectRepo $projectRepo
	pulumi config set stackCommit $commitSHA
	pulumi config set stackName $stackName/dev
	echo "config set, deploy cluster"
	pulumi up
	echo "check if operator is running"
	kubectl get pods -o wide -l name=pulumi-kubernetes-operator

##@ Start dapr apk locally

.PHONY: start-hub
start-hub:
	echo "run hub"
	dapr run --app-id hub --app-protocol grpc --app-port 50051 --config ./configs/dapr-config.yaml -- go run ./cmd/server/hub/main.go

.PHONY: start-discoverer
start-discoverer:
	echo "run discoverer"
	dapr run --app-id discoverer --app-protocol grpc --app-port 50052 --config ./configs/dapr-config.yaml -- go run ./cmd/server/discoverer/main.go

.PHONY: start-downloader
start-downloader:
	echo "run downloader"
	dapr run --app-id downloader --app-protocol grpc --app-port 50053 --config ./configs/dapr-config.yaml -- go run ./cmd/server/downloader/main.go

.PHONY: start-parser
start-parser:
	echo "run parser"
	dapr run --app-id parser --app-protocol grpc --app-port 50054 --config ./configs/dapr-config.yaml -- go run ./cmd/server/parser/main.go

.PHONY: start-translator
start-translator:
	echo "run translator"
	dapr run --app-id translator --app-protocol grpc --app-port 50055 --config ./configs/dapr-config.yaml -- go run ./cmd/server/downloader/main.go

.PHONY: start-analyzer
start-analyzer:
	echo "run analyzer"
	dapr run --app-id analyzer --app-protocol grpc --app-port 50056 --config ./configs/dapr-config.yaml -- go run ./cmd/server/analyzer/main.go


##@ Deploy project to local cluster

.PHONY: start-local-cluter
start-local-cluter: kind create cluster


.PHONY: deploy-local
deploy-local:
	echo "deploy whopper infrastructure to local system"
	pulumi config set kconfig="$(kubectl config view --raw=true)" --secret -s local
	pulumi up -s local -y

.PHONY: destroy-local
destroy-local: pulumi destroy -s local -y

##@ Deploy project to GKE cluster

.PHONY: deploy-prd
deploy-local: pulumi up -s prd -y

.PHONY: destroy-prd
destroy-local: pulumi destroy -s prd -y


##@ Update golang dependencies

.PHONY: update-deps-go
update-deps-go: GO111MODULE=on
update-deps-go:
	go get -u -t ./...
	go mod tidy
	go mod verify