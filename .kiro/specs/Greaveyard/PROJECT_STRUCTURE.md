# Project Structure

```
process-monitor-cli/
├── cmd/
│   └── graveyard/
│       └── main.go                 # Entry point
├── internal/
│   ├── app/
│   │   └── state.go               # AppState with channels and mutexes
│   ├── model/
│   │   ├── process.go             # Process struct
│   │   ├── metrics.go             # SystemMetrics, CPUInfo, MemoryInfo, GPUInfo
│   │   └── config.go              # Config struct
│   ├── service/
│   │   ├── process.go             # ProcessService (using gopsutil)
│   │   ├── system.go              # SystemService (CPU, Memory, GPU)
│   │   ├── ai.go                  # AIService (Gemini integration)
│   │   └── config.go              # ConfigService (.env management)
│   ├── ui/
│   │   └── manager.go             # UIManager (tview UI)
│   └── logger/
│       └── logger.go              # Logging utility
├── bin/                           # Build output (gitignored)
│   ├── graveyard.exe             # Windows binary
│   ├── graveyard                 # Linux binary
│   └── graveyard-arm             # ARM binary
├── .env.example                   # Example environment variables
├── .env                          # User's API key (gitignored)
├── go.mod                        # Go module definition
├── go.sum                        # Go dependencies checksum
├── Makefile                      # Build automation
├── README.md                     # Documentation
└── .gitignore                    # Git ignore rules
```

## Key Dependencies (go.mod)

```go
module github.com/yourusername/process-monitor-cli

go 1.21

require (
    github.com/rivo/tview v0.0.0-20231126152417-33a1d271f2b6
    github.com/gdamore/tcell/v2 v2.6.0
    github.com/shirou/gopsutil/v3 v3.23.11
    github.com/google/generative-ai-go v0.5.0
    github.com/joho/godotenv v1.5.1
    google.golang.org/api v0.152.0
)
```

## Build Commands

```bash
# Install dependencies
go mod download

# Build for current platform
go build -o bin/graveyard cmd/graveyard/main.go

# Build for all platforms
make build-all

# Build optimized (smaller binary)
go build -ldflags="-s -w" -o bin/graveyard cmd/graveyard/main.go

# Run without building
go run cmd/graveyard/main.go
```

## Binary Sizes (Approximate)

- Windows (amd64): ~10-12 MB
- Linux (amd64): ~10-12 MB
- Linux (arm64): ~10-12 MB
- Optimized with -ldflags: ~8-10 MB

## Advantages of Go Implementation

1. **Single Binary**: No runtime dependencies needed
2. **Cross-Platform**: Compile once for all platforms
3. **Fast Startup**: <10ms startup time
4. **Low Memory**: ~20-30MB memory footprint
5. **Concurrent**: Goroutines for efficient monitoring
6. **Easy Distribution**: Just download and run
