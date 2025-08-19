#!/bin/bash
set -e

echo "Running go vet..."
go vet ./...

# echo "Running staticcheck..."
# staticcheck ./...
# Uncomment the above lines when staticcheck is added to the CI/dev environment.

echo "Linting complete."
