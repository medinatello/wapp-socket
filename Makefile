# Makefile for the wapp-socket project

.PHONY: build run-cli run-daemon test lint help

GO_CMD=go
APP_NAME=wapp-socket
CLI_PATH=./cmd/whats-cli
DAEMON_PATH=./cmd/whatsd

# ====================================================================================
# HELP
# ====================================================================================

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build         Build all binaries"
	@echo "  run-cli       Build and run the CLI"
	@echo "  run-daemon    Build and run the daemon"
	@echo "  test          Run all tests"
	@echo "  lint          Run linter"
	@echo "  help          Show this help message"


# ====================================================================================
# DEVELOPMENT
# ====================================================================================

build:
	@echo "Building binaries..."
	$(GO_CMD) build -o bin/$(APP_NAME)-cli $(CLI_PATH)
	$(GO_CMD) build -o bin/$(APP_NAME)-daemon $(DAEMON_PATH)

run-cli:
	@echo "Running CLI..."
	$(GO_CMD) run $(CLI_PATH) -- --seed 12345

run-daemon:
	@echo "Running daemon..."
	$(GO_CMD) run $(DAEMON_PATH)

test:
	@echo "Running tests..."
	$(GO_CMD) test ./...

lint:
	@echo "Running linter..."
	$(GO_CMD) vet ./...
	# staticcheck ./... # Uncomment when ready
