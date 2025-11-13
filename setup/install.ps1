<#
.SYNOPSIS
    Graveyard Installer Script for Windows (PowerShell)
.DESCRIPTION
    This script downloads and installs the latest release of Graveyard 
    from GitHub and adds it to the user's PATH.
#>

# --- Configuration ---
$Repo = "rexreus/Graveyard"
$InstallDir = "$env:LOCALAPPDATA\Graveyard"
$BinaryName = "graveyard.exe"
# ---------------------

Write-Host "Installing Graveyard..."

# 1. Create install directory if it doesn't exist
if (-not (Test-Path -Path $InstallDir)) {
    Write-Host "Creating install directory at: $InstallDir"
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
} else {
    Write-Host "Install directory already exists."
}

# 2. Get latest release version from GitHub API
Write-Host "Fetching latest release version..."
try {
    $ApiUrl = "https://api.github.com/repos/$Repo/releases/latest"
    $LatestRelease = Invoke-RestMethod -Uri $ApiUrl -ErrorAction Stop
    $LatestVersion = $LatestRelease.tag_name
}
catch {
    Write-Error "Failed to fetch latest version: $_"
    exit 1
}

if ([string]::IsNullOrEmpty($LatestVersion)) {
    Write-Error "Failed to parse latest version. Aborting."
    exit 1
}

Write-Host "Latest version: $LatestVersion"

# 3. Download the binary
$DownloadUrl = "https://github.com/$Repo/releases/download/$LatestVersion/$BinaryName"
$OutFile = Join-Path -Path $InstallDir -ChildPath $BinaryName

Write-Host "Downloading from: $DownloadUrl"
try {
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $OutFile -ErrorAction Stop
    Write-Host "Download complete."
}
catch {
    Write-Error "Download failed: $_"
    exit 1
}

# 4. Add to user's PATH if not already there
Write-Host "Checking User PATH..."
try {
    # Get the persistent User PATH from the registry
    $UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
    
    # Check if our install directory is already in the path
    $PathEntries = $UserPath.Split(';') | Where-Object { $_ -ne "" }
    
    if ($PathEntries -notcontains $InstallDir) {
        Write-Host "Adding $InstallDir to User PATH..."
        
        # Create the new path string
        $NewUserPath = $UserPath
        if (-not $NewUserPath.EndsWith(';')) {
            $NewUserPath += ';'
        }
        $NewUserPath += $InstallDir
        
        # Set the persistent User PATH
        [Environment]::SetEnvironmentVariable("Path", $NewUserPath, "User")
        
        # Also update the current session's PATH
        $env:Path += ";$InstallDir"
        
        Write-Host "`nNOTE: PATH has been updated."
        Write-Host "Please restart your terminal or run 'refreshenv' for changes to take full effect."
    } else {
        Write-Host "Install directory is already in your User PATH."
    }
}
catch {
    Write-Warning "Could not automatically update PATH: $_"
    Write-Warning "Please add the following directory to your PATH manually:"
    Write-Warning $InstallDir
}

Write-Host ""
Write-Host "Installation complete!"
Write-Host "Graveyard installed to: $InstallDir"
Write-Host ""
Write-Host "Run 'graveyard' to start the application."
Write-Host "(You may need to restart your terminal first)"
