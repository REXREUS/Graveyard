# Project Structure

## Directory Layout

```
process-monitor-cli/
├── cmd/graveyard/          # Application entry point
│   └── main.go            # Main function, initialization, monitoring loop
├── internal/              # Private application code
│   ├── app/               # Application state management
│   │   └── state.go       # Thread-safe state with channels
│   ├── logger/            # Logging utility
│   │   └── logger.go      # File-based logging
│   ├── model/             # Data models (no business logic)
│   │   ├── config.go      # Configuration model
│   │   ├── metrics.go     # System metrics models
│   │   └── process.go     # Process model
│   ├── service/           # Business logic layer
│   │   ├── ai.go          # Gemini AI integration
│   │   ├── config.go      # Configuration management
│   │   ├── process.go     # Process monitoring (gopsutil)
│   │   ├── system.go      # System metrics (CPU, Memory, GPU)
│   │   └── virustotal.go  # VirusTotal API integration
│   └── ui/                # Terminal UI layer
│       └── manager.go     # tview UI management
├── bin/                   # Build output (gitignored)
├── docs/                  # Documentation
│   ├── ARCHITECTURE.md    # System design
│   ├── BUILD.md           # Build instructions
│   ├── FEATURES.md        # Feature documentation
│   ├── GPU_MONITORING.md  # GPU implementation details
│   ├── INSTALL.md         # Installation guide
│   ├── PROJECT_SUMMARY.md # High-level overview
│   ├── QUICKSTART.md      # Quick start guide
│   ├── SECURITY.md        # Security considerations
│   ├── UI_LAYOUT.md       # UI design guide
│   ├── VIRUSTOTAL_INTEGRATION.md     # VT guide (English)
│   └── VIRUSTOTAL_INTEGRATION_ID.md  # VT guide (Indonesian)
├── setup/                 # Installation scripts
│   ├── build.sh           # Linux/macOS build script
│   ├── build.bat          # Windows build script
│   ├── install.sh         # Linux/macOS installer
│   ├── install.bat        # Windows installer
│   ├── uninstall.sh       # Linux/macOS uninstaller
│   └── uninstall.bat      # Windows uninstaller
├── .env                   # API keys (gitignored)
├── .env.example           # Environment template
├── .gitignore             # Git ignore rules
├── CHANGELOG.md           # Version history
├── CONTRIBUTING.md        # Contribution guidelines
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
├── graveyard.log          # Application logs (gitignored)
├── LICENSE                # MIT License
├── Makefile               # Build automation
└── README.md              # Main documentation
```

## Architecture Layers

### 1. Entry Point (`cmd/graveyard/`)
- Contains only `main.go`
- Initializes all services and dependencies
- Sets up signal handling for graceful shutdown
- Starts monitoring goroutine and UI event loop

### 2. UI Layer (`internal/ui/`)
- Single file: `manager.go`
- Manages tview application and three-panel layout
- Handles keyboard events (navigation, actions)
- Renders process list, metrics, AI responses, VirusTotal results
- Coordinates with AppState for thread-safe updates

### 3. Service Layer (`internal/service/`)
- **ai.go**: Google Gemini API integration for process and security analysis
- **config.go**: Manages `.env` file and API key retrieval
- **process.go**: Fetches process list using gopsutil
- **system.go**: Retrieves CPU, memory, and GPU metrics (platform-specific)
- **virustotal.go**: VirusTotal API v3 integration for malware scanning

### 4. Application State (`internal/app/`)
- **state.go**: Central, thread-safe state management
- Uses `sync.RWMutex` for concurrent access
- Channels for async communication:
  - `ProcessChan`: Process list updates
  - `MetricsChan`: System metrics updates
  - `AIResponseChan`: AI analysis responses
  - `VTResponseChan`: VirusTotal scan results

### 5. Model Layer (`internal/model/`)
- **config.go**: Configuration structure
- **metrics.go**: SystemMetrics, GPUInfo structures
- **process.go**: Process structure with helper methods
- Pure data definitions, no business logic

### 6. Logger (`internal/logger/`)
- **logger.go**: File-based logging utility
- Writes to `graveyard.log`
- Used throughout the application for debugging

## Code Organization Principles

### Separation of Concerns
- UI code never directly calls external APIs
- Services handle all external interactions
- Models contain only data structures
- State management isolated in `app` package

### Dependency Flow
```
main.go → UI Manager → Services → Models
              ↓
          AppState (channels)
```

### File Naming
- One primary struct per file (e.g., `state.go` contains `AppState`)
- Service files named after their domain (e.g., `ai.go` for AI service)
- Clear, descriptive names matching Go conventions

### Package Structure
- `internal/` prevents external imports
- Each subdirectory is a separate package
- Minimal inter-package dependencies
- Services are independent and composable

## Key Files

- **cmd/graveyard/main.go**: Application bootstrap and monitoring loop
- **internal/ui/manager.go**: All UI logic and event handling
- **internal/app/state.go**: Central state with channel-based updates
- **internal/service/system.go**: Platform-specific GPU monitoring implementation
- **Makefile**: Build automation and cross-compilation
- **.env.example**: Template for API key configuration
