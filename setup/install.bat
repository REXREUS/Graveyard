@echo off
REM Graveyard Installer Script for Windows
REM This script downloads and installs Graveyard from GitHub releases

setlocal enabledelayedexpansion

set REPO=rexreus/Graveyard
set INSTALL_DIR=%LOCALAPPDATA%\Graveyard
set BINARY_NAME=graveyard.exe

echo Installing Graveyard...

REM Create install directory
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

REM Get latest release version using PowerShell
echo Fetching latest release...
for /f "delims=" %%i in ('powershell -Command "(Invoke-RestMethod -Uri 'https://api.github.com/repos/%REPO%/releases/latest').tag_name"') do set LATEST_VERSION=%%i

if "%LATEST_VERSION%"=="" (
    echo Failed to fetch latest version
    exit /b 1
)   

echo Latest version: %LATEST_VERSION%

REM Download URL
set DOWNLOAD_URL=https://github.com/%REPO%/releases/download/%LATEST_VERSION%/%BINARY_NAME%

echo Downloading from: %DOWNLOAD_URL%

REM Download binary using PowerShell
echo Downloading...
powershell -Command "Invoke-WebRequest -Uri '%DOWNLOAD_URL%' -OutFile '%INSTALL_DIR%\%BINARY_NAME%'"

if not exist "%INSTALL_DIR%\%BINARY_NAME%" (
    echo Download failed
    exit /b 1
)

REM Add to PATH if not already there
echo Checking PATH...
echo %PATH% | find /i "%INSTALL_DIR%" >nul
if errorlevel 1 (
    echo Adding to PATH...
    setx PATH "%PATH%;%INSTALL_DIR%"
    echo.
    echo NOTE: Please restart your terminal for PATH changes to take effect
)

echo.
echo Installation complete!
echo Graveyard installed to: %INSTALL_DIR%
echo.
echo Run 'graveyard' to start the application
echo ^(You may need to restart your terminal first^)

pause
