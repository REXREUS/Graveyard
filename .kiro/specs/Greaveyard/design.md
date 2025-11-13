# Design Document

## Overview

The Process Monitor CLI is a cross-platform terminal application built with Go that provides real-time system monitoring with AI-powered process analysis. The application compiles to a single binary executable with no runtime dependencies, making it easy to distribute. It uses the `tview` library for terminal UI rendering, `gopsutil` for cross-platform system metrics, and Google Gemini AI SDK for intelligent process explanations.

## Architecture

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Terminal UI Layer                        │
│         (tview - Application, Boxes, Lists, Gauges)         │
└──────────────────┬──────────────────────────────────────────┘
                   │
┌──────────────────▼──────────────────────────────────────────┐
│                  Application Layer                           │
│  - UI Manager (Layout, Event Handling, Rendering)           │
│  - App State (Process Data, UI State, Channels)             │
└──────────────────┬──────────────────────────────────────────┘
                   │
┌──────────────────▼──────────────────────────────────────────┐
│                   Service Layer                              │
│  - Process Service (Fetch, Monitor via gopsutil)            │
│  - System Service (CPU, Memory, GPU Metrics)                │
│  - AI Service (Gemini API Integration)                      │
│  - Config Service (API Key Management, .env)                │
└──────────────────┬──────────────────────────────────────────┘
                   │
┌──────────────────▼──────────────────────────────────────────┐
│                  Platform Layer                              │
│  - gopsutil (Cross-platform process/system info)            │
│  - Runtime detection (GOOS for platform-specific logic)     │
└─────────────────────────────────────────────────────────────┘
```

### Component Interaction Flow

1. **Startup**: App initializes → Load config → Initialize UI → Start goroutines
2. **Monitoring Loop**: Goroutine runs every 2s → Fetch processes → Send to channel → UI updates
3. **User Interaction**: Keypress → Event handler → Service call → Update state → Render
4. **AI Query**: Select process → Press 'i' → Goroutine calls Gemini API → Display response
5. **Kill Process**: Select process → Press 'k' → Confirm → Kill → Refresh list

### Concurrency Model

Go's goroutines and channels enable efficient concurrent operations:
- **Main goroutine**: Runs tview UI event loop
- **Monitor goroutine**: Fetches process/system data every 2s
- **AI goroutine**: Handles async AI API calls
- **Channels**: Communicate data between goroutines safely

## Components and Interfaces

### 1. UI Manager (`internal/ui/manager.go`)

Manages the terminal interface using tview library.

**Responsibilities:**
- Initialize tview application and UI components
- Handle layout and positioning
- Manage keyboard event bindings
- Coordinate rendering updates

**Key Methods:**
```go
type UIManager struct {
    app            *tview.Application
    processList    *tview.Table
    cpuGauge       *tview.TextView
    memoryGauge    *tview.TextView
    gpuGauge       *tview.TextView
    infoPanel      *tview.TextView
    vtPanel        *tview.TextView
    selectedIndex  int
    focusedPanel   string // "process", "ai", or "vt"
}

func NewUIManager(state, processService, systemService, aiService, configService, vtService) *UIManager
func (u *UIManager) InitializeLayout()
func (u *UIManager) RenderProcessList(processes []Process)
func (u *UIManager) RenderSystemMetrics(metrics SystemMetrics)
func (u *UIManager) RenderInfoPanel(content string)
func (u *UIManager) RenderVTPanel(content string)
func (u *UIManager) ShowSettingsDialog()
func (u *UIManager) ShowKillConfirmation(process Process) bool
func (u *UIManager) ShowMessage(message, msgType string)
func (u *UIManager) cycleFocus()
func (u *UIManager) setFocus(panel string)
func (u *UIManager) scanWithVirusTotal()
func (u *UIManager) formatVTResults(vtResult *VTScanResult) string
func (u *UIManager) Run() error
func (u *UIManager) Stop()
func (u *UIManager) StartUpdateLoop(ctx context.Context)
```

**UI Layout:**
```
┌─────────────────────────────────────────────────────────────┐
│  Process Monitor v1.0.0          [s] Settings  [q] Quit     │
├──────────────────────┬──────────────────────────────────────┤
│  System Resources    │  Process List                        │
│                      │                                      │
│  CPU:  [████░░] 45%  │  PID    NAME         CPU%   MEM%    │
│  MEM:  [███░░░] 32%  │  1234   chrome       12.5   450MB   │
│  GPU:  [██░░░░] 23%  │  5678   node         5.2    120MB   │
│                      │  9012   svchost      0.1    45MB    │
├──────────────────────┤  ...                                │
│  AI Assistant        │                                      │
│                      │                                      │
│  [Process info       │                                      │
│   will appear here]  │                                      │
│                      │                                      │
└──────────────────────┴──────────────────────────────────────┘
│  ↑↓: Navigate  [i]: Inspect  [k]: Kill  [s]: Settings      │
└─────────────────────────────────────────────────────────────┘
```

### 2. Process Service (`internal/service/process.go`)

Handles process monitoring using gopsutil library.

**Responsibilities:**
- Fetch all running processes
- Get process metrics (CPU, memory)
- Kill processes by PID
- Handle platform-specific differences

**Key Methods:**
```go
type ProcessService struct {
    cache map[int32]*process.Process
}

func NewProcessService() *ProcessService
func (p *ProcessService) GetProcessList() ([]Process, error)
func (p *ProcessService) KillProcess(pid int32) error
func (p *ProcessService) getProcessInfo(proc *process.Process) (*Process, error)
```

**Process Data Structure:**
```go
type Process struct {
    PID           int32
    Name          string
    CPUPercent    float64
    MemoryMB      uint64
    MemoryPercent float32
    Username      string
    Command       string
}
```

### 3. System Service (`internal/service/system.go`)

Retrieves overall system metrics using `gopsutil` library.

**Responsibilities:**
- Get CPU usage percentage
- Get memory usage and total
- Detect and get GPU metrics
- Calculate percentages

**Key Methods:**
```go
type SystemService struct {
    hasGPU bool
}

func NewSystemService() *SystemService
func (s *SystemService) GetCPUUsage() (float64, error)
func (s *SystemService) GetMemoryUsage() (*MemoryInfo, error)
func (s *SystemService) GetGPUUsage() (*GPUInfo, error)
func (s *SystemService) HasGPU() bool
```

**System Metrics Structure:**
```go
type SystemMetrics struct {
    CPU    CPUInfo
    Memory MemoryInfo
    GPU    *GPUInfo // nil if no GPU
}

type CPUInfo struct {
    Usage float64
    Cores int
}

type MemoryInfo struct {
    UsedGB      float64
    TotalGB     float64
    Percentage  float64
}

type GPUInfo struct {
    Available   bool
    Usage       float64
    MemoryUsedMB  uint64
    MemoryTotalMB uint64
}
```

### 4. AI Service (`internal/service/ai.go`)

Integrates with Google Gemini API for process analysis.

**Responsibilities:**
- Initialize Gemini client
- Send process queries asynchronously
- Format responses
- Handle API errors

**Key Methods:**
```go
type AIService struct {
    client *genai.Client
    model  *genai.GenerativeModel
    apiKey string
}

func NewAIService(apiKey string) (*AIService, error)
func (a *AIService) AnalyzeProcess(ctx context.Context, proc Process) (string, error)
func (a *AIService) buildPrompt(proc Process) string
func (a *AIService) formatResponse(response string) string
func (a *AIService) Close() error
```

**AI Prompt Templates:**

*Process Analysis:*
```
You are a system process analyzer. Provide a concise analysis of the following process:

Process Name: {name}
PID: {pid}
CPU Usage: {cpu}%
Memory Usage: {memory}MB

Please provide:
1. Purpose: What is this process and what does it do?
2. Safety: Is this process safe or potentially harmful?
3. Recommendation: Can this process be safely terminated?

Keep the response clear, professional, and under 200 words.
```

*VirusTotal Analysis:*
```
You are a cybersecurity analyst. Analyze the following VirusTotal scan result for a running process:

Process Information:
- Name: {name}
- PID: {pid}
- CPU Usage: {cpu}%
- Memory Usage: {memory}MB
- File Path: {path}
- File Hash (SHA256): {hash}

VirusTotal Scan Results:
- Threat Level: {level}
- Malicious Detections: {malicious}
- Suspicious Detections: {suspicious}
- Harmless: {harmless}
- Undetected: {undetected}
- Total Engines: {total}

Top Detections:
{detections}

Please provide:
1. Risk Assessment: Evaluate the overall risk level based on the scan results
2. Analysis: Explain what the detections mean and whether they are false positives
3. Recommendation: Should this process be terminated? What actions should be taken?
4. Additional Context: Any relevant information about this process type

Keep the response clear, professional, and actionable. Use bullet points for clarity.
```

### 5. Config Service (`internal/service/config.go`)

Manages application configuration and API key storage.

**Responsibilities:**
- Load/save .env file
- Manage multiple API keys (Gemini and VirusTotal)
- Validate configuration

**Key Methods:**
```go
type ConfigService struct {
    configPath string
    config     *Config
}

type Config struct {
    APIKey          string
    VTAPIKey        string
    RefreshInterval int
}

func NewConfigService() *ConfigService
func (c *ConfigService) LoadConfig() error
func (c *ConfigService) SaveAPIKey(key string) error
func (c *ConfigService) SaveVTAPIKey(key string) error
func (c *ConfigService) GetAPIKey() string
func (c *ConfigService) GetVTAPIKey() string
func (c *ConfigService) DeleteAPIKey() error
func (c *ConfigService) ValidateAPIKey(key string) bool
```

### 6. App State (`internal/app/state.go`)

Centralized state management using Go channels and mutexes.

**Responsibilities:**
- Store current processes
- Track selected process
- Manage UI state
- Thread-safe state updates
- Handle VirusTotal response state

**Key Methods:**
```go
type AppState struct {
    mu              sync.RWMutex
    processes       []Process
    systemMetrics   SystemMetrics
    selectedIndex   int
    aiResponse      string
    vtResponse      string
    processChan     chan []Process
    metricsChan     chan SystemMetrics
    aiResponseChan  chan string
    vtResponseChan  chan string
}

func NewAppState() *AppState
func (a *AppState) UpdateProcesses(processes []Process)
func (a *AppState) UpdateSystemMetrics(metrics SystemMetrics)
func (a *AppState) SetSelectedProcess(index int)
func (a *AppState) SetAIResponse(response string)
func (a *AppState) SetVTResponse(response string)
func (a *AppState) GetProcesses() []Process
func (a *AppState) GetSelectedProcess() *Process
```

### 7. VirusTotal Service (`internal/service/virustotal.go`)

Integrates with VirusTotal API for malware scanning.

**Responsibilities:**
- Calculate file hashes (SHA256)
- Query VirusTotal API
- Parse scan results
- Handle API errors and rate limits

**Key Methods:**
```go
type VirusTotalService struct {
    apiKey     string
    httpClient *http.Client
}

func NewVirusTotalService(apiKey string) *VirusTotalService
func (v *VirusTotalService) ScanProcess(ctx context.Context, proc Process) (*VTScanResult, error)
func (v *VirusTotalService) getProcessExecutablePath(pid int32) (string, error)
func (v *VirusTotalService) calculateFileHash(filePath string) (string, error)
func (v *VirusTotalService) getFileReport(ctx context.Context, fileHash string) (*VTFileReport, error)
```

## Data Models

### Process Model (`internal/model/process.go`)
```go
type Process struct {
    PID           int32
    Name          string
    CPUPercent    float64
    MemoryMB      uint64
    MemoryPercent float32
    Username      string
    Command       string
}

func (p Process) String() string {
    return fmt.Sprintf("%-8d %-20s %.1f%%  %dMB", 
        p.PID, p.Name, p.CPUPercent, p.MemoryMB)
}
```

### SystemMetrics Model (`internal/model/metrics.go`)
```go
type SystemMetrics struct {
    CPU       CPUInfo
    Memory    MemoryInfo
    GPU       *GPUInfo
    Timestamp time.Time
}

type CPUInfo struct {
    Usage float64
    Cores int
}

type MemoryInfo struct {
    UsedGB     float64
    TotalGB    float64
    Percentage float64
}

type GPUInfo struct {
    Available     bool
    Name          string
    Usage         float64
    MemoryUsedMB  uint64
    MemoryTotalMB uint64
}

type VTScanResult struct {
    ProcessName  string
    PID          int32
    FilePath     string
    FileHash     string
    Malicious    int
    Suspicious   int
    Undetected   int
    Harmless     int
    Detections   []string
    ScanDate     time.Time
    TotalEngines int
    Scans        map[string]ScanResult
}

type ScanResult struct {
    Detected bool
    Result   string
}

func (v *VTScanResult) GetThreatLevel() string
func (v *VTScanResult) GetSummary() string
```

### Config Model (`internal/model/config.go`)
```go
type Config struct {
    APIKey          string `json:"api_key"`
    VTAPIKey        string `json:"vt_api_key"`
    RefreshInterval int    `json:"refresh_interval"`
}

func DefaultConfig() *Config {
    return &Config{
        APIKey:          "",
        VTAPIKey:        "",
        RefreshInterval: 2000,
    }
}
```

## Error Handling

### Error Types

1. **Process Fetch Error**: When unable to retrieve process list
   - Display error message in UI
   - Retry after interval
   - Keep showing cached data

2. **AI API Error**: When Gemini API fails
   - Show error in info panel
   - Suggest checking API key
   - Allow retry

3. **Kill Process Error**: When unable to terminate process
   - Show error dialog with reason
   - Suggest running with elevated privileges

4. **Configuration Error**: When .env file is corrupted
   - Reset to defaults
   - Prompt user to reconfigure

### Error Handling Strategy

```javascript
try {
  // Operation
} catch (error) {
  logger.error(error)
  uiManager.showError(error.message)
  // Graceful degradation
}
```

## Testing Strategy

### Unit Tests

**Target Coverage: Core business logic**

1. **Process Service Tests**
   - Test process parsing for each platform
   - Test data normalization
   - Mock system commands

2. **AI Service Tests**
   - Test prompt generation
   - Test response formatting
   - Mock Gemini API

3. **Config Service Tests**
   - Test .env file operations
   - Test API key validation

### Integration Tests

1. **UI Integration**
   - Test keyboard navigation
   - Test state updates trigger renders
   - Test dialog flows

2. **Service Integration**
   - Test process service → UI flow
   - Test AI service → UI flow

### Manual Testing

1. **Cross-Platform Testing**
   - Test on Windows 10/11
   - Test on Ubuntu/Debian Linux
   - Test on macOS (ARM)

2. **UI/UX Testing**
   - Test with different terminal sizes
   - Test with many processes (1000+)
   - Test keyboard shortcuts

3. **Error Scenarios**
   - Test without API key
   - Test with invalid API key
   - Test killing protected processes
   - Test without GPU

## Technology Stack

### Core Dependencies

```go
require (
    github.com/rivo/tview v0.0.0-20231126152417-33a1d271f2b6
    github.com/gdamore/tcell/v2 v2.6.0
    github.com/shirou/gopsutil/v3 v3.23.11
    github.com/google/generative-ai-go v0.5.0
    github.com/joho/godotenv v1.5.1
    google.golang.org/api v0.152.0
)
```

**Library Descriptions:**
- `tview` - Terminal UI framework (like blessed for Node.js)
- `tcell` - Low-level terminal handling (used by tview)
- `gopsutil` - Cross-platform system and process utilities
- `generative-ai-go` - Google Gemini AI SDK for Go
- `godotenv` - .env file management
- `google.golang.org/api` - Google API client (required by Gemini SDK)

### Build Configuration

**Cross-Compilation:**
```bash
# Build for Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o bin/graveyard.exe cmd/graveyard/main.go

# Build for Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -o bin/graveyard cmd/graveyard/main.go

# Build for Linux ARM (64-bit)
GOOS=linux GOARCH=arm64 go build -o bin/graveyard-arm cmd/graveyard/main.go

# Build for macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o bin/graveyard-darwin cmd/graveyard/main.go

# Build for macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o bin/graveyard-darwin-arm cmd/graveyard/main.go
```

**Optimized Build (smaller binary):**
```bash
go build -ldflags="-s -w" -o bin/graveyard cmd/graveyard/main.go
```

### Platform-Specific Considerations

**All Platforms:**
- `gopsutil` handles platform differences automatically
- No need for platform-specific commands (ps, tasklist, etc.)
- Process killing uses `process.Kill()` which works cross-platform

**GPU Detection:**
- Use `gopsutil` GPU detection
- Fallback gracefully if GPU not available
- Support for NVIDIA, AMD, and Intel GPUs

## Performance Considerations

1. **Refresh Rate**: 2-second interval balances responsiveness and CPU usage
2. **Process Filtering**: Limit display to top 100 processes by CPU/Memory
3. **Goroutine Efficiency**: Use goroutines for concurrent data fetching
4. **Channel Buffering**: Buffer channels to prevent blocking
5. **API Caching**: Cache AI responses for same process names (using sync.Map)
6. **Memory Management**: Go's garbage collector handles cleanup automatically
7. **Binary Size**: Compiled binary ~8-12MB (can be reduced with build flags)

## Security Considerations

1. **API Key Storage**: Store in .env file with restricted permissions
2. **Process Killing**: Require confirmation before termination
3. **Command Injection**: Sanitize all process names before display
4. **Privilege Escalation**: Warn users when elevated privileges needed
5. **API Rate Limiting**: Implement cooldown between AI requests

## Future Enhancements

1. Process search/filter functionality
2. Process tree visualization
3. Historical metrics graphs
4. Export process data to CSV
5. Custom themes and color schemes
6. Process grouping by application
7. Network usage per process
8. Disk I/O monitoring
