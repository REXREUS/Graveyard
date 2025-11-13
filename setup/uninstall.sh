#!/bin/bash

# Graveyard Uninstaller Script for Linux/Mac
# This script removes Graveyard from your system

set -e

INSTALL_DIR="/usr/local/bin"
BINARY_NAME="graveyard"
CONFIG_DIR="$HOME/.config/graveyard"
LOG_FILE="$HOME/graveyard.log"

echo "Uninstalling Graveyard..."

# Remove binary
if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
    echo "Removing binary from $INSTALL_DIR..."
    if [ -w "$INSTALL_DIR" ]; then
        rm "$INSTALL_DIR/$BINARY_NAME"
    else
        sudo rm "$INSTALL_DIR/$BINARY_NAME"
    fi
    echo "✓ Binary removed"
else
    echo "Binary not found at $INSTALL_DIR/$BINARY_NAME"
fi

# Ask about config and log files
echo ""
read -p "Do you want to remove configuration and log files? (y/N): " -n 1 -r
echo ""

if [[ $REPLY =~ ^[Yy]$ ]]; then
    # Remove config directory
    if [ -d "$CONFIG_DIR" ]; then
        echo "Removing config directory..."
        rm -rf "$CONFIG_DIR"
        echo "✓ Config directory removed"
    fi
    
    # Remove log file
    if [ -f "$LOG_FILE" ]; then
        echo "Removing log file..."
        rm "$LOG_FILE"
        echo "✓ Log file removed"
    fi
    
    echo ""
    echo "✓ Graveyard completely removed from your system"
else
    echo ""
    echo "✓ Graveyard binary removed (config and logs kept)"
    echo "Config location: $CONFIG_DIR"
    echo "Log location: $LOG_FILE"
fi
