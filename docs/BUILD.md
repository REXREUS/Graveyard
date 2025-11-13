# Build Instructions

## Prerequisites

- **Go**: Version 1.21 or higher.
- **Make**: Optional, for using the provided `Makefile` shortcuts.
- **Git**: For cloning the repository.

## Quick Build

For a standard build on your current platform:

```bash
# 1. Clone the repository
git clone https://github.com/yourusername/process-monitor-cli.git
cd process-monitor-cli

# 2. Download dependencies
make install

# 3. Build the application
make build
```

The executable will be located at `bin/graveyard` (or `bin\graveyard.exe` on Windows).

## Build Targets (via Makefile)

The `Makefile` provides several convenient targets:

### Development Build
Creates a standard build with debugging symbols.

```bash
make build
```

### Optimized Build
Creates a smaller binary by stripping debug information. Recommended for distribution.

```bash
make build-optimized
```

### Cross-Platform Builds
Builds for Windows, Linux (amd64, arm64), and macOS (Intel, Apple Silicon).

```bash
make build-all
```

This creates the following binaries in the `bin/` directory:
- `graveyard.exe` (Windows amd64)
- `graveyard` (Linux amd64)
- `graveyard-arm` (Linux arm64)
- `graveyard-darwin` (macOS amd64)
- `graveyard-darwin-arm` (macOS arm64)

### Clean Build Artifacts
Removes the `bin/` directory and all compiled executables.

```bash
make clean
```

## Manual Compilation

If you are not using `make`, you can compile the application manually.

### Standard Build
```bash
go build -o bin/graveyard cmd/graveyard/main.go
```

### Optimized Build
```bash
go build -ldflags="-s -w" -o bin/graveyard cmd/graveyard/main.go
```

### Manual Cross-Compilation```bash
# Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o bin/graveyard.exe cmd/graveyard/main.go

# Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -o bin/graveyard cmd/graveyard/main.go

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o bin/graveyard-darwin-arm cmd/graveyard/main.go
```

## Dependency Management

### Download Dependencies
This command downloads the necessary Go modules. The `Makefile` target `install` runs this.

```bash
go mod download```

### Tidy Dependencies
To ensure `go.mod` and `go.sum` are accurate and up-to-date:

```bash
go mod tidy
```

## Development

For a faster development cycle, you can run the application directly without compiling an executable first.

```bash
# Run directly
go run cmd/graveyard/main.go

# Or use the Makefile shortcut
make run
```

## Build Troubleshooting

### "go: command not found"
Ensure Go is installed correctly and that its `bin` directory is in your system's `PATH`. Verify with `go version`.

### "cannot find package"
Your Go modules may be missing or out of sync. Run `go mod download` or `go mod tidy`.

### Cross-compilation fails
Some dependencies may require CGO. If you encounter issues, try disabling it:
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/graveyard cmd/graveyard/main.go
```

## See Also

- [Installation Guide](INSTALL.md) - Detailed installation instructions.
- [Architecture](ARCHITECTURE.md) - System design and components.
- [Project Summary](PROJECT_SUMMARY.md) - High-level overview.
