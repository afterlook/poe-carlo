#!/usr/bin/bash

set -e  # Exit on error
set -u  # Exit on undefined variable

echo "Running all tests with race detection and coverage..."
echo "=================================================="

# Run all tests including tools tests
go test ./... -race -coverprofile=coverage.out -covermode=atomic -v

echo ""
echo "=================================================="
echo "All tests passed successfully!"
echo ""
echo "Coverage report generated: coverage.out"
echo "View coverage: go tool cover -html=coverage.out"
