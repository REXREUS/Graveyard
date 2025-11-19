<#
.SYNOPSIS
    Graveyard Uninstaller Script for Windows.
    Removes the application binary, configuration folder, and log file.

.DESCRIPTION
    This script is the PowerShell equivalent of the original Batch script.
    It uses standard PowerShell cmdlets for reliable file system operations.

.NOTES
    Enterprise Ready: Uses strict error handling and full paths.
#>
# --- Strict Error Handling ---
# Ensures the script stops immediately upon a non-terminating error.
$ErrorActionPreference = "Stop"

# --- Configuration Variables ---
# Use built-in PowerShell environment variables for reliability
$INSTALL_DIR = Join-Path $env:LOCALAPPDATA "Graveyard"
$BINARY_NAME = "graveyard.exe"
$CONFIG_DIR  = Join-Path $env:APPDATA "graveyard"
$LOG_FILE    = Join-Path $env:USERPROFILE "graveyard.log"

$BinaryPath = Join-Path $INSTALL_DIR $BINARY_NAME

Write-Host "============================================="
Write-Host "💀 Starting Graveyard Uninstallation Process 💀"
Write-Host "============================================="
Write-Host ""

# --- 1. Remove Binary File ---
Write-Host "--- 1. Removing Binary ---"
if (Test-Path -Path $BinaryPath -PathType Leaf) {
    Write-Host "Removing binary from $INSTALL_DIR..."
    try {
        # Use -Force to suppress prompts, but only targeting the file
        Remove-Item -Path $BinaryPath -Force -Confirm:$false
        Write-Host "✓ Binary removed successfully."
    }
    catch {
        Write-Host "❌ ERROR: Failed to remove binary. Details: $($_.Exception.Message)" -ForegroundColor Red
        # Do not stop the script here, continue to cleanup user data
    }
} else {
    Write-Host "Binary not found at $BinaryPath. Skipping."
}

# --- 2. Remove Install Directory (if empty) ---
Write-Host "--- 2. Removing Installation Directory ---"
if (Test-Path -Path $INSTALL_DIR -PathType Container) {
    try {
        # Attempt to remove the directory. If it fails (not empty), suppress the error.
        Remove-Item -Path $INSTALL_DIR -Recurse -Force -ErrorAction SilentlyContinue
        
        # Check if the directory was actually removed (i.e., it was empty)
        if (-not (Test-Path -Path $INSTALL_DIR)) {
            Write-Host "✓ Installation directory ($INSTALL_DIR) removed (was empty)."
        } else {
            Write-Host "⚠️ Installation directory ($INSTALL_DIR) kept (not empty or permissions issue)." -ForegroundColor Yellow
        }
    }
    catch {
        Write-Host "❌ An unexpected error occurred while checking/removing $INSTALL_DIR." -ForegroundColor Red
    }
} else {
    Write-Host "Installation directory not found. Skipping."
}

# --- 3. Ask about Config and Log Files ---
Write-Host ""
Write-Host "--- 3. User Data Cleanup ---"
$REMOVE_DATA = Read-Host "Do you want to remove configuration and log files? (y/N)"

if ($REMOVE_DATA -ceq "y" -or $REMOVE_DATA -ceq "Y") {
    # --- Remove Config Directory ---
    if (Test-Path -Path $CONFIG_DIR -PathType Container) {
        Write-Host "Removing config directory ($CONFIG_DIR)..."
        Remove-Item -Path $CONFIG_DIR -Recurse -Force -Confirm:$false
        Write-Host "✓ Config directory removed."
    } else {
        Write-Host "Config directory not found. Skipping."
    }
    
    # --- Remove Log File ---
    if (Test-Path -Path $LOG_FILE -PathType Leaf) {
        Write-Host "Removing log file ($LOG_FILE)..."
        Remove-Item -Path $LOG_FILE -Force -Confirm:$false
        Write-Host "✓ Log file removed."
    } else {
        Write-Host "Log file not found. Skipping."
    }
    
    Write-Host ""
    Write-Host "============================================="
    Write-Host "✅ Graveyard completely removed from your system" -ForegroundColor Green
    Write-Host "============================================="
} else {
    Write-Host ""
    Write-Host "============================================="
    Write-Host "⚠️ Graveyard binary removed (config and logs kept)" -ForegroundColor Yellow
    Write-Host "Config location: $CONFIG_DIR"
    Write-Host "Log location: $LOG_FILE"
    Write-Host "============================================="
}

Write-Host ""
Write-Host "NOTE: You may need to manually remove $INSTALL_DIR from your PATH if you added it during installation." -ForegroundColor Cyan

# The 'pause' command is not idiomatic in PowerShell; we use Read-Host for the same effect.
# If the script is run in a separate window, it will stay open until a key is pressed.
if ($Host.Name -eq "ConsoleHost") {
    Write-Host "Press any key to exit..."
    $null = [System.Console]::ReadKey()
}
