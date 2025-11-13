# Architecture Overview

## Project Structure

```
process-monitor-cli/
├── cmd/graveyard/          # Application entry point
│   └── main.go            # Main function, initialization
├── internal/
│   ├── app/               # Application state management
│   │   └── state.go       # Thread-safe state with channels
│   ├── model/             # Data models
│   │   ├── config.go      # Configuration model
│   │   ├── metrics.go     # System metrics models
│   │   └── process.go     # Process model
│   ├── service/           # Business logic layer
│   │   ├── ai.go          # Gemini AI integration
│   │   ├── config.go      # Configuration management
│   │   ├── process.go     # Process monitoring (gopsutil)
│   │   ├── system.go      # System metrics (CPU, Memory, GPU)
│   │   └── virustotal.go  # VirusTotal API integration
│   ├── ui/                # Terminal UI layer
│   │   └── manager.go     # tview UI management
│   └── logger/            # Logging utility
│       └── logger.go      # File-based logging
├── bin/                   # Build output (gitignored)
├── .env                   # API keys (gitignored)
├── .env.example           # Environment template
├── go.mod                 # Go module definition
├── Makefile               # Build automation
└── README.md              # Documentation
```

## Component Layers

### 1. UI Layer (`internal/ui/`)
- Manages terminal interface using `tview`.
- Handles keyboard events for navigation and actions.
- Renders process list, system metrics, AI responses, and VirusTotal results.
- Features a three-panel layout for simultaneous data display.
- Coordinates with `AppState` for thread-safe data updates.

### 2. Service Layer (`internal/service/`)
- **ProcessService**: Fetches and manages processes using `gopsutil`.
- **SystemService**: Retrieves CPU and Memory metrics. Implements advanced, platform-specific GPU monitoring using `nvidia-smi` (for NVIDIA) and PowerShell/WMIC (for Windows).
- **AIService**: Integrates with Google Gemini API for process and security analysis.
- **VirusTotalService**: Integrates with VirusTotal API v3 for malware scanning via file hashes.
- **ConfigService**: Manages `.env` file and API keys securely.

### 3. Application Layer (`internal/app/`)
- **AppState**: Central, thread-safe state management.
- Uses channels for asynchronous communication between the monitoring goroutine and the UI update loop.
- Stores processes, metrics, AI/VT responses, and UI state.

### 4. Model Layer (`internal/model/`)
- Defines data structures for `Process`, `SystemMetrics`, `GPUInfo`, `VTScanResult`, and `Config`.
- Contains no business logic, only data definitions and helper methods (e.g., `GetThreatLevel`).

## Data Flow

```
┌─────────────────────────────────────────────────────────┐
│                    User Input (Keyboard)                 │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              UI Manager (Event Handler)                  │
│  - Navigation (↑↓)                                      │
│  - Inspect (i) → AI Service                             │
│  - VT Scan (t) → VirusTotal Service → AI Service        │
│  - Kill (k) → Process Service                           │
│  - Settings (s) → Config Service                        │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                   App State (Channels)                   │
│  - processChan: []Process                               │
│  - metricsChan: SystemMetrics                           │
│  - aiResponseChan: string                               │
│  - vtResponseChan: string                               │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              UI Update Loop (Goroutine)                  │
│  - Listens on all channels                              │
│  - app.QueueUpdateDraw() for thread-safe rendering      │
└─────────────────────────────────────────────────────────┘
```

## Concurrency Model

### Goroutines

1.  **Main Goroutine**: Runs the `tview` application and event loop.
2.  **Monitoring Goroutine**: Periodically fetches process and system data every 2 seconds.
3.  **UI Update Goroutine**: Listens on state channels and queues UI updates.
4.  **AI & VT Request Goroutines**: Each AI or VirusTotal request is spawned in a new goroutine to prevent blocking the UI.

### Channels

- `processChan`: Sends process list updates.
- `metricsChan`: Sends system metrics updates.
- `aiResponseChan`: Sends AI analysis responses for the left panel.
- `vtResponseChan`: Sends VirusTotal scan results for the center panel.

### Thread Safety

- `sync.RWMutex` in `AppState` for safe concurrent access to shared data.
- `app.QueueUpdateDraw()` for thread-safe UI updates from any goroutine.
- Buffered channels to prevent blocking during state updates.

## Key Technologies

- **Go**: Core programming language.
- **tview**: Terminal UI framework.
- **gopsutil**: Cross-platform system and process utilities.
- **generative-ai-go**: Google Gemini AI SDK.
- **VirusTotal API v3**: For malware scanning.
- **System Commands**: `nvidia-smi`, `wmic`, and `powershell` for advanced GPU monitoring on specific platforms.
- **godotenv**: Environment variable management for API keys.

## Build Process

```bash
# Development build
make build

# Optimized build (smaller binary)
make build-optimized

# Cross-compilation for all supported platforms
make build-all
```

## Error Handling Strategy

1.  **Service Errors**: Logged to `graveyard.log` and user-friendly messages are displayed in the UI.
2.  **API Errors**: Shown in the relevant UI panel (AI or VT), suggesting configuration checks.
3.  **Process Errors**: The application continues with cached data and retries on the next cycle.
4.  **Kill Errors**: A modal dialog is shown with the specific error reason (e.g., permission denied).

## Performance Considerations

- Limits process list to the top 100 by CPU usage to reduce overhead.
- 2-second refresh interval balances responsiveness and system load.
- Asynchronous goroutines for non-blocking API calls.
- Channel buffering to prevent blocking between data fetching and UI rendering.
- Efficient UI updates using `QueueUpdateDraw()`.
- Caching of process data for accurate CPU percentage calculation over time.
