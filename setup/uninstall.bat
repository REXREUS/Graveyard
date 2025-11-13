@echo off
REM Graveyard Uninstaller Script for Windows
REM This script removes Graveyard from your system

setlocal enabledelayedexpansion

set INSTALL_DIR=%LOCALAPPDATA%\Graveyard
set BINARY_NAME=graveyard.exe
set CONFIG_DIR=%APPDATA%\graveyard
set LOG_FILE=%USERPROFILE%\graveyard.log

echo Uninstalling Graveyard...

REM Remove binary
if exist "%INSTALL_DIR%\%BINARY_NAME%" (
    echo Removing binary from %INSTALL_DIR%...
    del "%INSTALL_DIR%\%BINARY_NAME%"
    echo Binary removed
) else (
    echo Binary not found at %INSTALL_DIR%\%BINARY_NAME%
)

REM Remove install directory if empty
if exist "%INSTALL_DIR%" (
    rmdir "%INSTALL_DIR%" 2>nul
)

REM Ask about config and log files
echo.
set /p REMOVE_DATA="Do you want to remove configuration and log files? (y/N): "

if /i "%REMOVE_DATA%"=="y" (
    REM Remove config directory
    if exist "%CONFIG_DIR%" (
        echo Removing config directory...
        rmdir /s /q "%CONFIG_DIR%"
        echo Config directory removed
    )
    
    REM Remove log file
    if exist "%LOG_FILE%" (
        echo Removing log file...
        del "%LOG_FILE%"
        echo Log file removed
    )
    
    echo.
    echo Graveyard completely removed from your system
) else (
    echo.
    echo Graveyard binary removed ^(config and logs kept^)
    echo Config location: %CONFIG_DIR%
    echo Log location: %LOG_FILE%
)

echo.
echo NOTE: You may need to manually remove %INSTALL_DIR% from your PATH
echo       if you added it during installation

pause
