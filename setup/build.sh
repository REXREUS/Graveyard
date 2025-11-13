#!/bin/bash

# Build script for Graveyard - Linux/Mac
# This script builds the project for all platforms

set -e

echo "Building Graveyard for all platforms..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Build for Windows
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o ../bin/graveyard.exe ../cmd/graveyard/main.go

# Build for Linux
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o ../bin/graveyard ../cmd/graveyard/main.go

echo "Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -o ../bin/graveyard-arm ../cmd/graveyard/main.go

# Build for macOS
echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o ../bin/graveyard-darwin ../cmd/graveyard/main.go

echo "Building for macOS (arm64/M1)..."
GOOS=darwin GOARCH=arm64 go build -o ../bin/graveyard-darwin-arm ../cmd/graveyard/main.go

echo ""
echo "Build complete! Binaries are in the bin/ directory:"
ls -lh ../bin/
