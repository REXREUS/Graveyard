#!/bin/bash

# Graveyard Installer Script for Linux/Mac
# This script downloads and installs Graveyard from GitHub releases

set -e

REPO="rexreus/Graveyard"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="graveyard"

echo "Installing Graveyard..."

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
    linux)
        OS_NAME="linux"
        ;;
    darwin)
        OS_NAME="darwin"
        ;;
    *)
        echo "Unsupported operating system: $OS"
        exit 1
        ;;
esac

case $ARCH in
    x86_64|amd64)
        ARCH_NAME="amd64"
        if [ "$OS_NAME" = "darwin" ]; then
            BINARY_NAME="graveyard-darwin"
        fi
        ;;
    arm64|aarch64)
        ARCH_NAME="arm64"
        if [ "$OS_NAME" = "linux" ]; then
            BINARY_NAME="graveyard-arm"
        elif [ "$OS_NAME" = "darwin" ]; then
            BINARY_NAME="graveyard-darwin-arm"
        fi
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

echo "Detected: $OS_NAME $ARCH_NAME"

# Get latest release version
echo "Fetching latest release..."
LATEST_VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo "Failed to fetch latest version"
    exit 1
fi

echo "Latest version: $LATEST_VERSION"

# Download URL
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/$BINARY_NAME"

echo "Downloading from: $DOWNLOAD_URL"

# Download binary
TMP_FILE="/tmp/$BINARY_NAME"
curl -L -o "$TMP_FILE" "$DOWNLOAD_URL"

if [ ! -f "$TMP_FILE" ]; then
    echo "Download failed"
    exit 1
fi

# Make executable
chmod +x "$TMP_FILE"

# Install (may require sudo)
echo "Installing to $INSTALL_DIR..."
if [ -w "$INSTALL_DIR" ]; then
    mv "$TMP_FILE" "$INSTALL_DIR/graveyard"
else
    sudo mv "$TMP_FILE" "$INSTALL_DIR/graveyard"
fi

echo ""
echo "✓ Graveyard installed successfully!"
echo "Run 'graveyard' to start the application"
