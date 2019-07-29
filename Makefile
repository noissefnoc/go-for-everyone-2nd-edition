# Meta information
NAME := myproj
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

# Install developer dependencies
## Setup
.PHONY: devel-deps
devel-deps: deps
	GO111MODULE=off go get \
	  golang.org/x/lint/golint \
	  github.com/motemen/gobump/cmd/gobump \
	  github.com/Songmu/make2help/cmd/make2help

# Execute Tests
## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binaries ex. make bin/myproj
bin/%: cmd/%/main.go devel-deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

## build binary
.PHONY: build
build: bin/myproj

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
