# Makefile for the wapp-socket project

.PHONY: build run-cli run-daemon test test-coverage lint clean deps install help all

GO_CMD=go
APP_NAME=wapp-socket
CLI_PATH=./cmd/whats-cli
DAEMON_PATH=./cmd/whatsd
BIN_DIR=./bin

# Version and build info
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "unknown")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# ====================================================================================
# HELP
# ====================================================================================

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Development:"
	@echo "  build           Build all binaries"
	@echo "  run-cli         Build and run the CLI"
	@echo "  run-daemon      Build and run the daemon"
	@echo "  deps            Download and tidy dependencies"
	@echo ""
	@echo "Testing:"
	@echo "  test            Run all tests"
	@echo "  test-coverage   Run tests with coverage report"
	@echo "  test-domain     Run domain tests with coverage"
	@echo ""
	@echo "Quality:"
	@echo "  lint            Run linter and static analysis"
	@echo "  fmt             Format code"
	@echo ""
	@echo "Utilities:"
	@echo "  clean           Clean build artifacts"
	@echo "  install         Install tools"
	@echo "  all             Build, test, and lint"
	@echo "  help            Show this help message"

# ====================================================================================
# DEVELOPMENT
# ====================================================================================

all: deps fmt lint test build

build:
	@echo "Building binaries..."
	@mkdir -p $(BIN_DIR)
	$(GO_CMD) build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME)-cli $(CLI_PATH)
	$(GO_CMD) build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME)-daemon $(DAEMON_PATH)
	@echo "✅ Binaries built successfully"

run-cli:
	@echo "Running CLI..."
	$(GO_CMD) run $(CLI_PATH) connect

run-daemon:
	@echo "Running daemon..."
	$(GO_CMD) run $(DAEMON_PATH)

deps:
	@echo "Downloading dependencies..."
	$(GO_CMD) mod download
	$(GO_CMD) mod tidy
	@echo "✅ Dependencies updated"

# ====================================================================================
# TESTING
# ====================================================================================

test:
	@echo "Running tests..."
	$(GO_CMD) test ./...
	@echo "✅ All tests passed"

test-coverage:
	@echo "Running tests with coverage..."
	$(GO_CMD) test -cover ./...
	@echo ""
	@echo "Detailed coverage for critical components:"
	@echo "Domain (target: 100%):"
	@$(GO_CMD) test -cover ./internal/domain
	@echo "Application (target: 80%):"
	@$(GO_CMD) test -cover ./internal/app
	@echo "Logger (target: 70%):"
	@$(GO_CMD) test -cover ./internal/adapter/log/slog

test-domain:
	@echo "Running domain tests with detailed coverage..."
	$(GO_CMD) test -v -cover ./internal/domain

# ====================================================================================
# QUALITY
# ====================================================================================

lint:
	@echo "Running static analysis..."
	$(GO_CMD) vet ./...
	$(GO_CMD) mod verify
	@echo "✅ Static analysis passed"

fmt:
	@echo "Formatting code..."
	$(GO_CMD) fmt ./...
	@echo "✅ Code formatted"

# ====================================================================================
# UTILITIES
# ====================================================================================

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BIN_DIR)
	rm -f coverage.out
	@echo "✅ Clean completed"

install:
	@echo "Installing development tools..."
	@echo "Note: Install golangci-lint manually: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
	@echo "✅ Check tool installation"

# ====================================================================================
# PROJECT SPECIFIC
# ====================================================================================

demo: build
	@echo "Running demo sequence..."
	@echo "1. Starting daemon in background..."
	@$(BIN_DIR)/$(APP_NAME)-daemon &
	@DAEMON_PID=$$!; \
	sleep 2; \
	echo "2. Running CLI commands..."; \
	$(BIN_DIR)/$(APP_NAME)-cli connect || true; \
	$(BIN_DIR)/$(APP_NAME)-cli send --to "demo@s.whatsapp.net" --text "Hello from make!" || true; \
	echo "3. Stopping daemon..."; \
	kill $$DAEMON_PID 2>/dev/null || true

check-standards:
	@echo "Checking project standards compliance..."
	@echo "🔍 Checking build..."
	@$(GO_CMD) build ./... && echo "✅ Build: PASS" || echo "❌ Build: FAIL"
	@echo "🔍 Checking domain coverage..."
	@$(GO_CMD) test -cover ./internal/domain | grep -q "100.0%" && echo "✅ Domain coverage: PASS (100%)" || echo "❌ Domain coverage: FAIL"
	@echo "🔍 Checking app coverage..."
	@$(GO_CMD) test -cover ./internal/app | grep -q -E "([8-9][0-9]|100)\..*%" && echo "✅ App coverage: PASS (>80%)" || echo "❌ App coverage: FAIL"
	@echo "🔍 Checking code format..."
	@test -z "$$($(GO_CMD) fmt ./...)" && echo "✅ Format: PASS" || echo "❌ Format: FAIL"
