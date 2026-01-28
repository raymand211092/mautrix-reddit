#!/bin/bash

set -euo pipefail

# Variables de build
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "unknown")
COMMIT=$(git rev-parse HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S UTC')

# Flags de build
LDFLAGS="-X 'main.Tag=${VERSION}' -X 'main.Commit=${COMMIT}' -X 'main.BuildTime=${BUILD_TIME}'"

echo "Building mautrix-reddit ${VERSION}..."
echo "Commit: ${COMMIT}"
echo "Build time: ${BUILD_TIME}"

# Build
go build -ldflags "${LDFLAGS}" -o mautrix-reddit ./cmd/mautrix-reddit

echo "Build complete! Binary: ./mautrix-reddit"
