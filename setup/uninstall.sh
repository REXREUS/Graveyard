#!/bin/bash
#
# Graveyard Uninstaller Script - Enterprise Grade
# Removes the Graveyard application binary, configuration, and logs completely.
#
# Execution Prerequisites: The user must have permissions to run 'sudo' if the binary exists in a system directory.

# --- Strict Shell Options for Security and Reliability ---
# -e: Exit immediately if a command exits with a non-zero status.
# -u: Treat unset variables as an error.
# -o pipefail: The pipeline's exit status is the rightmost non-zero exit status, or zero if all commands exit successfully.
set -euo pipefail

# --- Configuration Variables ---
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="graveyard"
CONFIG_DIR="$HOME/.config/graveyard"
LOG_FILE="$HOME/graveyard.log"

# --- Main Program Start ---
echo "=========================================="
echo "💀 Starting Graveyard Uninstallation Process 💀"
echo "=========================================="

# --- Function Definitions ---

# Safely remove the binary, checking permissions before escalating with sudo.
remove_binary() {
    local target_path="$INSTALL_DIR/$BINARY_NAME"
    echo -n "Checking for binary at $target_path... "

    if [[ -f "$target_path" ]]; then
        echo "Binary found."
        
        # Check if user has direct write access to the installation directory
        if [[ -w "$INSTALL_DIR" ]]; then
            echo "Removing binary without sudo..."
            rm -v "$target_path"
        else
            echo "Requires sudo permission to remove binary from system directory."
            # Use sudo to remove the binary, ensuring user is prompted for password only when needed
            if sudo rm -f "$target_path"; then
                echo "✓ Binary removed successfully (using sudo)."
            else
                # This branch should be reached if sudo fails or user cancels
                echo "❌ ERROR: Failed to remove binary. Check permissions or user cancellation." >&2
                return 1
            fi
        fi
    else
        echo "Binary not found. Skipping removal."
    fi
}

# Handles the removal of configuration and log files after user confirmation.
cleanup_user_data() {
    echo ""
    # Use -r for raw input, essential for security and correct handling
    read -r -p "Do you want to remove configuration ($CONFIG_DIR) and log files ($LOG_FILE)? (Y/n): " REPLY

    # Default to NO if REPLY is empty or check for explicit 'y' or 'Y'
    if [[ "$REPLY" =~ ^[Yy]$ ]]; then
        echo "--- Proceeding with user data cleanup ---"

        # Remove config directory
        if [[ -d "$CONFIG_DIR" ]]; then
            echo -n "Removing config directory ($CONFIG_DIR)... "
            if rm -rf "$CONFIG_DIR"; then
                echo "✓ Config directory removed."
            else
                echo "❌ Failed to remove config directory. Manual cleanup may be required." >&2
            fi
        else
            echo "Config directory not found. Skipping."
        fi

        # Remove log file
        if [[ -f "$LOG_FILE" ]]; then
            echo -n "Removing log file ($LOG_FILE)... "
            if rm "$LOG_FILE"; then
                echo "✓ Log file removed."
            else
                echo "❌ Failed to remove log file. Manual cleanup may be required." >&2
            fi
        else
            echo "Log file not found. Skipping."
        fi

        echo ""
        echo "=========================================="
        echo "✅ Graveyard completely removed from your system."
        echo "=========================================="
    else
        echo ""
        echo "=========================================="
        echo "⚠️ Graveyard binary removal completed (config and logs kept)."
        echo "Config location: $CONFIG_DIR"
        echo "Log location: $LOG_FILE"
        echo "=========================================="
    fi
}

# --- Main Execution Flow ---
remove_binary
cleanup_user_data

exit 0
