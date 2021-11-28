
TOPDIR ?= $(shell pwd)
BUILD_VCSTAG ?= $(shell git log -n 1 --pretty=format:"%H")
BUILD_TIME ?= $(shell date -u "+%Y-%m-%dT%H:%M:%S+00:00")
BUILD_MODULE = $(shell go list -m)
BUILD_BINARY = $(shell basename `go list -m`)
BUILD_DIRECTORY = ${TOPDIR}/deploy/${BUILD_VCSTAG}

.DEFAULT_GOAL := all

.PHONY: all
all: setup report clean all-linux all-darwin all-windows

.PHONY: setup
setup:
	@mkdir -p ${BUILD_DIRECTORY}

.PHONY: report
report:
	@echo BUILD_VCSTAG=${BUILD_VCSTAG}
	@echo BUILD_TIME=${BUILD_TIME}
	@echo BUILD_MODULE=${BUILD_MODULE}
	@echo BUILD_BINARY=${BUILD_BINARY}
	@echo TOPDIR=${TOPDIR}
	@echo BUILD_DIRECTORY=${BUILD_DIRECTORY}

clean:
.PHONY: clean
	echo $(BUILD_DIRECTORY)
	go clean
	go mod tidy

.PHONY: all-windows all-%
all-windows:
	GOARCH="amd64" GOOS=windows BUILD_EXTENSION=".exe" ${MAKE} binary
all-%:
	GOARCH="amd64" GOOS=$* BUILD_EXTENSION="" ${MAKE} binary

.PHONY: binary
binary:
	@mkdir -p ${BUILD_DIRECTORY}/${GOOS}
	env \
		CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
		go build -ldflags \
		"-X main.VCSTag=$(BUILD_VCSTAG) -X main.BuildTime=$(BUILD_TIME)" \
		-a -installsuffix nocgo \
		-o "$(BUILD_DIRECTORY)/$(GOOS)/$(BUILD_BINARY)$(BUILD_EXTENSION)"