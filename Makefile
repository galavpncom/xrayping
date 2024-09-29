GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_MOD_DOWNLOAD=$(GO_CMD) mod download

BINARY_NAME=xrayping
BUILD_DIR=./build/bin

.PHONY: all
all: help

.PHONY: build
build:
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
