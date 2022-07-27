.DEFAULT_GOAL:=all
GO ?= GO111MODULE=on go

define goinstall
	mkdir -p $(shell pwd)/bin
	echo "Installing $(1)"
	GOBIN=$(shell pwd)/bin go install $(1)
endef

.PHONY: all
all: build fmt test example cover install

fmt:
	-$(GO) fmt ./...

test: gen-test generate
	$(GO) test -v -race -coverprofile=coverage.out ./...

cover: gen-test test
	$(GO) tool cover -html=coverage.out -o coverage.html

gen-test: build
	$(GO) generate $(PACKAGES)

install:
	$(GO) install

all: build

build:
	echo hello
