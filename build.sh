#!/bin/bash

# Create bin directory if it doesn't exist
mkdir -p bin

# Build for Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o bin/license-generator.exe .

# Build for macOS (64-bit)
GOOS=darwin GOARCH=amd64 go build -o bin/license-generator-darwin .

# Build for Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -o bin/license-generator .

echo "Binaries are in the bin directory."
