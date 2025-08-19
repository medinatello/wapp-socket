#!/bin/bash
set -e

# This script will be used to generate mocks for our interfaces in Sprint 2.
# It requires the 'mockery' tool to be installed.
# go install github.com/vektra/mockery/v2@latest

echo "Mock generation is placeholder for Sprint 1."
echo "In Sprint 2, this script would run a command like:"
echo "# mockery --all --keeptree --case=underscore"

# Example for a single interface:
# mockery --name=WebSocketDialer --dir=./internal/port/outbound --output=./internal/port/outbound/mocks

exit 0
