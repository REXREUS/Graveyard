# Technology Stack

## Language & Runtime

- **Go 1.21+**: Primary language for the entire application
- **Module**: `github.com/yourusername/process-monitor-cli`

## Core Dependencies

### UI & Terminal
- **tview** (`github.com/rivo/tview`): Terminal UI framework for responsive interface
- **tcell** (`github.com/gdamore/tcell/v2`): Low-level terminal handling

### System Monitoring
- **gopsutil** (`github.com/shirou/gopsutil/v3`): Cross-platform system and process utilities
- Platform-specific commands for GPU monitoring:
  - `nvidia-smi` for NVIDIA GPUs (all platforms)
  - `wmic` and PowerShell Performance Counters for Windows (all GPU vendors)

### AI & Security
- **Google Generative AI** (`github.com/google/generative-ai-go`): Gemini 1.5 Flash integration
- **VirusTotal API v3**: Malware scanning via REST API (manual HTTP calls)

### Configuration
- **godotenv** (`github.com/joho/godotenv`): Environment variable management for API keys

## Build System

### Makefile Targets

```bash
# Development
make install          # Download dependencies
make build           # Build for current platform
make run             # Run without building

# Production
make build-optimized # Build with size optimization (-ldflags="-s -w")
make build-all       # Cross-compile for all platforms

# Maintenance
make clean           # Remove build artifacts
```

### Cross-Platform Builds

The project supports cross-compilation for:
- Windows (amd64): `graveyard.exe`
- Linux (amd64): `graveyard`
- Linux (arm64): `graveyard-arm`
- macOS Intel (amd64): `graveyard-darwin`
- macOS Apple Silicon (arm64): `graveyard-darwin-arm`

Build scripts available in `setup/`:
- `build.sh` / `build.bat`: Platform-specific build scripts
- `install.sh` / `install.bat`: Installation scripts
- `uninstall.sh` / `uninstall.bat`: Uninstallation scripts

## Common Commands

### Development Workflow

```bash
# Initial setup
go mod download
cp .env.example .env
# Edit .env with API keys

# Development
go run cmd/graveyard/main.go
go fmt ./...

# Testing builds
make build
./bin/graveyard

# Release builds
make build-all
```

### Code Formatting

```bash
# Format all Go files
go fmt ./...

# Standard Go conventions apply
# Use gofmt before committing
```

## Configuration

### Environment Variables (.env)

```env
GEMINI_API_KEY=your_gemini_api_key_here
VIRUSTOTAL_API_KEY=your_virustotal_api_key_here
```

- Keys are loaded via `godotenv` at startup
- File should have `0600` permissions for security
- `.env` is gitignored; use `.env.example` as template

## Logging

- File-based logging to `graveyard.log`
- Custom logger in `internal/logger/`
- Logs errors, warnings, and info messages
- Used for debugging and troubleshooting

## Concurrency Model

- **Goroutines**: Asynchronous data fetching and UI updates
- **Channels**: Thread-safe communication between components
- **sync.RWMutex**: Protects shared state in `AppState`
- **Context**: Graceful shutdown handling

## Performance Characteristics

- ~1-2% CPU usage during monitoring
- ~20-30MB memory footprint
- 2-second refresh interval for process/system data
- Non-blocking AI and VirusTotal API calls
- Top 100 processes by CPU usage to reduce overhead
