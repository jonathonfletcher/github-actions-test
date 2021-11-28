
# GOOS=$(shell go env GOOS)
# GOARCH=$(shell go env GOARCH)

BUILD_VCSTAG ?= $(shell git log -n 1 --pretty=format:"%H")
BUILD_TIME ?= $(shell date -u "+%Y-%m-%dT%H:%M:%S+00:00")
BUILD_MODULE = $(shell go list -m)
BUILD_BINARY = $(shell basename `go list -m`)
BUILD_DIRECTORY = $(shell pwd)

.DEFAULT_GOAL := all

.PHONY: all
all: report clean all-linux all-darwin all-windows

.PHONY: all-%
all-%:
	GOARCH="amd64" GOOS=$* ${MAKE} binary

.PHONY: report
report:
	@echo BUILD_VCSTAG=${BUILD_VCSTAG}
	@echo BUILD_TIME=${BUILD_TIME}
	@echo BUILD_MODULE=${BUILD_MODULE}
	@echo BUILD_BINARY=${BUILD_BINARY}
	@echo BUILD_DIRECTORY=${BUILD_DIRECTORY}

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