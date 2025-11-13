@echo off
REM Build script for Graveyard - Windows
REM This script builds the project for all platforms

echo Building Graveyard for all platforms...

REM Create bin directory if it doesn't exist
if not exist bin mkdir bin

REM Build for Windows
echo Building for Windows (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -o bin\graveyard.exe ../cmd/graveyard/main.go 
if errorlevel 1 goto error

REM Build for Linux
echo Building for Linux (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -o bin\graveyard ../cmd/graveyard/main.go 
if errorlevel 1 goto error

echo Building for Linux (arm64)...
set GOOS=linux
set GOARCH=arm64
go build -o ../bin\graveyard-arm ../cmd/graveyard/main.go 
if errorlevel 1 goto error

REM Build for macOS
echo Building for macOS (amd64)...
set GOOS=darwin
set GOARCH=amd64
go build -o ../bin\graveyard-darwin ../cmd/graveyard/main.go 
if errorlevel 1 goto error

echo Building for macOS (arm64/M1)...
set GOOS=darwin
set GOARCH=arm64
go build -o ../bin\graveyard-darwin-arm ../cmd/graveyard/main.go 
if errorlevel 1 goto error

echo.
echo Build complete! Binaries are in the bin\ directory:
dir ../bin/
goto end

:error
echo.
echo Build failed!
exit /b 1

:end
