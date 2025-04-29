# Makefile for the code-decoder project

# Variables
BINARY_NAME=code-decoder
CMD_PATH=./cmd/code-decoder
OUTPUT_DIR=./bin
VERSION ?= $(shell git describe --tags --always --dirty || echo "v0.0.0-dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

# Default target: Build for the current OS/ARCH
.PHONY: build
build:
	@echo "Building $(BINARY_NAME) for $(shell go env GOOS)/$(shell go env GOARCH)..."
	@go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)"

# Cross-compilation targets
.PHONY: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-windows-amd64 build-windows-arm64

build-darwin-amd64:
	@echo "Building $(BINARY_NAME) for darwin/amd64..."
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64 $(CMD_PATH)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64"

build-darwin-arm64:
	@echo "Building $(BINARY_NAME) for darwin/arm64..."
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64 $(CMD_PATH)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64"

build-linux-amd64:
	@echo "Building $(BINARY_NAME) for linux/amd64..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64 $(CMD_PATH)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64"

build-linux-arm64:
	@echo "Building $(BINARY_NAME) for linux/arm64..."
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-arm64 $(CMD_PATH)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-linux-arm64"

build-windows-amd64:
	@echo "Building $(BINARY_NAME) for windows/amd64..."
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-windows-amd64.exe $(CMD_PATH)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-windows-amd64.exe"

build-windows-arm64:
	@echo "Building $(BINARY_NAME) for windows/arm64..."
	@GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-windows-arm64.exe $(CMD_PATH)
	@echo "Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)-windows-arm64.exe"

# Build all target
.PHONY: all
all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-windows-amd64 build-windows-arm64
	@echo "All builds complete in $(OUTPUT_DIR)/"

# Install target
.PHONY: install
install:
	@echo "Installing $(BINARY_NAME)..."
	@go install $(LDFLAGS) $(CMD_PATH)
	@echo "$(BINARY_NAME) installed to $(shell go env GOPATH)/bin"

# Clean target
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Clean complete."

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests complete."

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@echo "Coverage report generated."

# Run a specific test matching pattern
.PHONY: test-one
test-one:
	@if [ -z "$(TEST)" ]; then \
		echo "Please specify a test pattern with TEST=pattern"; \
		exit 1; \
	fi
	@echo "Running tests matching pattern: $(TEST)"
	@go test ./... -run $(TEST) -v
	@echo "Tests complete."

# Help target
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build                 Build $(BINARY_NAME) for the current OS/ARCH (default)"
	@echo "  build-darwin-amd64  Build for macOS (Intel)"
	@echo "  build-darwin-arm64  Build for macOS (Apple Silicon)"
	@echo "  build-linux-amd64   Build for Linux (amd64)"
	@echo "  build-linux-arm64   Build for Linux (arm64)"
	@echo "  build-windows-amd64 Build for Windows (amd64)"
	@echo "  build-windows-arm64 Build for Windows (arm64)"
	@echo "  all                   Build for all supported platforms"
	@echo "  install               Install $(BINARY_NAME) using go install"
	@echo "  test                  Run all tests"
	@echo "  test-coverage         Run tests with coverage report"
	@echo "  test-one              Run tests matching a pattern (use TEST=pattern)"
	@echo "  clean                 Remove build artifacts from $(OUTPUT_DIR)/"
	@echo "  help                  Show this help message"

# Set default goal
.DEFAULT_GOAL := help
