# Product Overview

Graveyard is a cross-platform, terminal-based process monitoring tool that combines real-time system metrics with AI-powered security analysis. It's designed for developers, system administrators, and security professionals who need intelligent process management from the command line.

## Core Value Proposition

- **Real-time monitoring**: Live CPU, memory, and GPU tracking with auto-refresh
- **AI-powered insights**: Google Gemini integration for process analysis and security assessment
- **Malware detection**: VirusTotal scanning with AI interpretation of results
- **Cross-platform**: Supports Windows, Linux, and macOS (including ARM architectures)

## Key Features

- Process monitoring with accurate resource usage calculations
- System metrics visualization (CPU, memory, GPU utilization)
- AI-powered process inspection and security recommendations
- VirusTotal malware scanning with dual-panel workflow (raw data + AI analysis)
- Safe process termination with confirmation dialogs
- Keyboard-driven interface for efficient navigation

## Target Users

- System administrators monitoring server/workstation processes
- Developers debugging resource consumption and performance issues
- Security analysts investigating suspicious processes
- General users troubleshooting system performance

## Security & Privacy

- API keys stored locally in `.env` with restricted permissions
- No telemetry or data collection
- Only file hashes (SHA256) sent to VirusTotal, never actual files
- Privacy-first design with local-only data processing
