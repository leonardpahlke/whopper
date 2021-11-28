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

##@ Update golang dependencies

.PHONY: update-deps-go
update-deps-go: GO111MODULE=on
update-deps-go:
	go get -u -t ./...
	go mod tidy
	go mod verify

##@ Verify go application with a docker container
virt-verify: verify-build virt-run

.PHONY: verify-build virt-jump virt-run
virt-build:
	docker build -t acs-cl2-shipper-container -f ./build/Dockerfile.test .	

virt-run:	
	docker run acs-cl2-shipper-container:latest

virt-jump:
	docker run --rm -it --entrypoint bash acs-cl2-shipper-container