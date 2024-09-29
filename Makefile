GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_FMT=$(GO_CMD) fmt
GO_CLEAN=$(GO_CMD) clean
GO_MOD_DOWNLOAD=$(GO_CMD) mod download

BINARY_NAME=xrayping
BUILD_DIR=./build/bin

# Platforms
LINUX_64=linux/amd64
LINUX_32=linux/386
LINUX_ARM32_V7A=linux/arm
LINUX_ARM64_V8A=linux/arm64

.PHONY: all
all: help

.PHONY: build
build:
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) main.go

.PHONY: build-linux-64
build-linux-64:
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go

.PHONY: build-linux-32
build-linux-32:
	GOOS=linux GOARCH=386 $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-386 main.go

.PHONY: build-linux-arm32-v7a
build-linux-arm32-v7a:
	GOOS=linux GOARCH=arm GOARM=7 $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm32-v7a main.go

.PHONY: build-linux-arm64-v8a
build-linux-arm64-v8a:
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64-v8a main.go

.PHONY: build-linux
build-linux: build-linux-64 build-linux-32 build-linux-arm32-v7a build-linux-arm64-v8a

.PHONY: format
format:
	$(GO_FMT) ./...

.PHONY: clean
clean:
	rm -f $(BUILD_DIR)/*

.PHONY: help
help:
	@echo "Usage:"
	@echo "  make build                  - Build for the current platform"
	@echo "  make build-linux-64         - Build for Linux amd64 (64-bit)"
	@echo "  make build-linux-32         - Build for Linux 386 (32-bit)"
	@echo "  make build-linux-arm32-v7a  - Build for Linux ARM 32-bit (ARMv7-a)"
	@echo "  make build-linux-arm64-v8a  - Build for Linux ARM 64-bit (ARMv8-a)"
	@echo "  make build-linux            - Build for all Linux platforms"
	@echo "  make format                 - Format Go source code"
	@echo "  make clean                  - Clean the build directory"
	@echo "  make help                   - Show this help message"
