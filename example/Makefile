
# GOOS=$(shell go env GOOS)
# GOARCH=$(shell go env GOARCH)

BUILD_VCSTAG ?= $(shell git describe --always --abbrev=16 --match "v[0-9]*" HEAD)
BUILD_TIME ?= $(shell date -u "+%Y-%m-%dT%H:%M:%S+00:00")
BUILD_MODULE = $(shell go list -m)
BUILD_BINARY = $(shell basename `go list -m`)
BUILD_DIRECTORY = $(shell pwd)

.DEFAULT_GOAL := all

.PHONY: all
all: clean all-linux all-darwin all-windows

.PHONY: all-%
all-%:
	GOARCH="amd64" GOOS=$* ${MAKE} binary

.PHONY: clean
clean:
	echo $(BUILD_DIRECTORY)
	go clean
	go mod tidy

.PHONY: binary
binary:
	env \
		CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
		go build -ldflags \
		"-X main.VCSTag=$(BUILD_VCSTAG) -X main.BuildTime=$(BUILD_TIME)" \
		-a -installsuffix nocgo \
		-o ${BUILD_BINARY}-${GOOS}