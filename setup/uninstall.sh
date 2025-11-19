#!/bin/bash
#
# Graveyard Uninstaller Script - Enterprise Ready
# Removes the Graveyard application binary, configuration, and logs.
#
# Prerequisites: Must be run with appropriate permissions for /usr/local/bin access.

# Set strict shell options for security and error handling
set -euo pipefail

# --- Configuration Variables ---
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="graveyard"
CONFIG_DIR="$HOME/.config/graveyard"
LOG_FILE="$HOME/graveyard.log"

echo "=========================================="
echo "💀 Starting Graveyard Uninstallation 💀"
echo "=========================================="

# --- Function Definitions ---

# Function to safely remove the binary, using sudo only if necessary
remove_binary() {
    local target_path="$INSTALL_DIR/$BINARY_NAME"
    echo -n "Checking binary at $target_path... "

    if [[ -f "$target_path" ]]; then
        if [[ -w "$INSTALL_DIR" ]]; then
            # User has write access, no sudo needed
            rm -v "$target_path" 2>/dev/null
        else
            # Sudo is required for cleanup in a system directory
            echo "Requires root permissions to remove."
            # Use 'rm -f' with sudo to suppress non-fatal warnings
            if sudo rm -f "$target_path"; then
                echo "✓ Binary removed successfully (using sudo)."
            else
                echo "❌ Failed to remove binary (sudo failed)." >&2
                return 1
            fi
        fi
        echo "✓ Binary removed."
    else
        echo "Binary not found. Skipping."
    fi
}

# Function to remove configuration and log files after user confirmation
cleanup_user_data() {
    echo ""
    read -r -p "Do you want to remove configuration ($CONFIG_DIR) and log files ($LOG_FILE)? (y/N): " REPLY

    if [[ "$REPLY" =~ ^[Yy]$ ]]; then
        echo "Proceeding with user data cleanup..."

        # Remove config directory
        if [[ -d "$CONFIG_DIR" ]]; then
            echo -n "Removing config directory... "
            if rm -rf "$CONFIG_DIR"; then
                echo "✓ Config directory removed."
            else
                echo "❌ Failed to remove config directory." >&2
            fi
        else
            echo "Config directory not found. Skipping."
        fi

        # Remove log file
        if [[ -f "$LOG_FILE" ]]; then
            echo -n "Removing log file... "
            if rm "$LOG_FILE"; then
                echo "✓ Log file removed."
            else
                echo "❌ Failed to remove log file." >&2
            fi
        else
            echo "Log file not found. Skipping."
        fi

        echo ""
        echo "=========================================="
        echo "✅ Graveyard completely removed from your system"
        echo "=========================================="
    else
        echo ""
        echo "=========================================="
        echo "⚠️ Graveyard binary removed (config and logs kept)"
        echo "Config location: $CONFIG_DIR"
        echo "Log location: $LOG_FILE"
        echo "=========================================="
    fi
}

# --- Main Execution ---

remove_binary
cleanup_user_data

exit 0
