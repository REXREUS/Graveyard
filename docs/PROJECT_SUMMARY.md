# Project Summary

## Overview

**Graveyard** is a cross-platform, terminal-based process monitoring tool built with Go. It goes beyond simple process listing by providing **real-time system metrics**, **AI-powered process analysis**, and a comprehensive **VirusTotal malware scanning workflow** with AI interpretation. Designed for developers, system administrators, and security professionals, it offers a powerful, command-line interface for managing and investigating running processes.

## Key Features

### 🔍 Real-Time Monitoring
- **Process Tracking**: Monitor all running processes with accurate CPU and memory usage calculations, sorted by resource consumption.
- **System Metrics**: Visual gauges for CPU, memory, and **advanced GPU metrics** (utilization, memory used/total) on Windows, Linux, and macOS.
- **Auto-Refresh**: Updates every 2 seconds for timely insights.

### 🤖 AI-Powered Analysis (Gemini)
- **Process Inspection**: Get intelligent explanations about any process, its purpose, and safety assessment from Google Gemini AI.
- **Security Context**: Understand the legitimacy and function of system processes.
- **Termination Guidance**: Receive clear recommendations on whether it's safe to kill a process.

### 🛡️ Advanced VirusTotal Integration
- **Malware Scanning**: Scan process executables against a database of over 70 antivirus engines using SHA256 hashing.
- **AI Security Analyst**: Gemini AI interprets complex VirusTotal results, providing a risk assessment and actionable recommendations.
- **Dual-Panel Workflow**: View raw VirusTotal data and AI analysis simultaneously for a complete security picture.
- **Threat Classification**: Clear visual indicators (HIGH/MEDIUM/SAFE/UNKNOWN) with color-coded alerts.

### ⚡ Process Management
- **Safe Termination**: Kill processes with confirmation dialogs and clear error messages.
- **Permission Handling**: Smart error messages guide you on privilege requirements.

### 🎨 Efficient User Interface
- **Three-Panel Design**:
  - **Left**: System metrics + AI Assistant for analysis.
  - **Center**: Dedicated VirusTotal scan results panel.
  - **Right**: Interactive process list.
- **Keyboard-Driven**: Navigate efficiently with simple shortcuts (`Tab`, arrow keys, `i`, `t`, `k`, `s`).

## Technology Stack

### Core Technologies
- **Language**: Go 1.21+
- **UI Framework**: tview (responsive terminal UI)
- **System Monitoring**: gopsutil (cross-platform system/process utilities)
- **AI Integration**: Google Gemini 1.5 Flash (for process analysis and VT interpretation)
- **Security Scanning**: VirusTotal API v3
- **Configuration**: godotenv (secure environment variable management)

### Advanced GPU Monitoring
- **Windows (All Vendors)**: Comprehensive metrics using `nvidia-smi`, `wmic`, and PowerShell Performance Counters.
- **Linux/macOS (NVIDIA)**: Full support via `nvidia-smi`.
- **Integrated GPU Support**: Accurate memory calculation for Intel/AMD integrated graphics.

### Architecture
- **Layered Design**: Clear separation of UI, service, and model layers.
- **Concurrent Processing**: Goroutines for asynchronous data fetching and non-blocking UI updates.
- **Channel-Based Communication**: Thread-safe state management via buffered channels.
- **Event-Driven UI**: Responsive keyboard event handling with panel focus management.

## Use Cases

### System Administration
- Monitor server and workstation processes in real-time.
- Quickly identify resource-intensive applications.
- Safely terminate problematic processes with AI guidance.
- Scan suspicious processes for malware using VirusTotal.

### Development & Debugging
- Track application resource consumption during development.
- Identify memory leaks and CPU bottlenecks.
- Get AI-powered explanations of complex system processes.
- Monitor GPU usage for graphics-intensive workloads.

### Security Analysis
- Investigate unknown or suspicious processes.
- Verify the legitimacy of system processes (e.g., `svchost.exe`).
- Scan executables for malware signatures.
- Receive AI-generated security recommendations based on VirusTotal data.

### General Monitoring
- Keep a pulse on system health.
- Understand what applications are running and their resource impact.
- Troubleshoot system performance issues.

## Platform Support

### Operating Systems
- **Windows 10/11**: Full support, including advanced GPU metrics.
- **Linux (amd64, arm64)**: Full process/system support, advanced GPU metrics for NVIDIA.
- **macOS (Intel, Apple Silicon)**: Full support, advanced GPU metrics for NVIDIA.

### GPU Vendors
- **NVIDIA**: Full metrics (utilization, memory used/total) on all supported OSes.
- **Intel & AMD (Windows)**: Full metrics via Performance Counters.
- **Integrated GPUs**: Special handling for accurate shared/dedicated memory reporting.

## Getting Started

### 1. Build from Source
```bash
git clone https://github.com/yourusername/process-monitor-cli.git
cd process-monitor-cli
make install  # Downloads dependencies
make build    # Compiles the executable
./bin/graveyard
```

### 2. Configure (Optional)
For AI and VirusTotal features, get API keys and configure:
- Press `s` in the application to open the Settings dialog.
- Or manually create a `.env` file with your keys.

## Documentation

### User Documentation
- [Quick Start Guide](QUICKSTART.md) - Get running in 5 minutes
- [Installation Guide](INSTALL.md) - Detailed installation instructions
- [Features Overview](FEATURES.md) - Complete feature list
- [VirusTotal Integration](VIRUSTOTAL_INTEGRATION.md) - Malware scanning guide

### Developer Documentation
- [Architecture](ARCHITECTURE.md) - System design and components
- [Build Instructions](BUILD.md) - Building for different platforms
- [GPU Monitoring](GPU_MONITORING.md) - Advanced GPU monitoring implementation
- [UI Layout](UI_LAYOUT.md) - User interface design guide

## Security & Privacy

- **API Key Protection**: Keys stored locally in `.env` with `0600` permissions.
- **No Telemetry**: The application collects no user data or system telemetry.
- **Privacy-First Scanning**: Only file hashes (SHA256) are sent to VirusTotal, not the actual files.
- **Local AI Analysis**: Process data is analyzed locally by Gemini; no sensitive information is stored.

## Performance

- **Resource Efficiency**: ~1-2% CPU usage, ~20-30MB memory footprint during monitoring.
- **Optimized Updates**: Data is cached and refreshed at 2-second intervals to balance performance and responsiveness.
- **Asynchronous Processing**: AI and VirusTotal calls do not block the UI.

## Future Roadmap

The project is actively maintained with a focus on enhancing the user experience and security analysis capabilities:

- **Process Filtering & Search**: Filter and sort processes by various criteria.
- **Export Features**: Export process lists and analysis reports.
- **Custom Alerts**: Set thresholds for CPU/memory/GPU usage.
- **Batch VirusTotal Scanning**: Scan multiple processes simultaneously.
- **Per-Process GPU Usage**: See which processes are using the GPU.

## Contributing

Contributions are welcome! Whether it's bug fixes, new features, documentation improvements, or GPU support for new vendors, we appreciate community involvement.

## License

MIT License. See [LICENSE](LICENSE) for details.

## Credits

- **tview** - Terminal UI framework
- **gopsutil** - Cross-platform system utilities
- **Google Gemini** - AI-powered process and security analysis
- **VirusTotal** - Malware scanning and threat intelligence

---

**Graveyard**: Transforming command-line process monitoring with AI and security intelligence.
