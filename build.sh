#!/bin/bash

# Create bin directory if it doesn't exist
mkdir -p bin

# Build for Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o bin/license-generator.exe ./src
zip -j bin/license-generator-windows-amd64.zip bin/license-generator.exe
rm bin/license-generator.exe

# Build for macOS (64-bit)
GOOS=darwin GOARCH=amd64 go build -o bin/license-generator ./src
tar -czf bin/license-generator-darwin-amd64.tar.gz -C bin license-generator
rm bin/license-generator

# Build for Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -o bin/license-generator ./src
tar -czf bin/license-generator-linux-amd64.tar.gz -C bin license-generator
rm bin/license-generator

echo "Compressed binaries are in the bin directory."