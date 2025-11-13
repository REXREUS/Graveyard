# Features Documentation

## Current Features

### 1. Process Monitoring

#### CPU Usage per Process
- **Real-time tracking** - CPU usage is calculated using a caching mechanism for accuracy over time.
- **Accurate readings** - Uses gopsutil's `CPUPercent()` with an interval-based calculation.
- **Smart formatting**:
  - `0.1%` for CPU < 1% (shows precision for low usage)
  - `5.2%` for CPU 1-10% (one decimal place)
  - `15%` for CPU >= 10% (no decimal, more compact)
- **Sorted display** - Processes are sorted by CPU usage (highest first).
- **Top 100 processes** - Shows the top 100 most resource-intensive processes.

#### Memory Usage per Process
- **RSS Memory** - Shows Resident Set Size (actual physical RAM used).
- **Compact format** - Displays in MB (e.g., `256MB`).
- **Percentage display** - Shows memory usage as a percentage of total RAM.

#### Process Information
- **PID** - Process ID
- **Name** - Process executable name (truncated if longer than 15 characters)
- **CPU%** - CPU usage percentage
- **RAM%** - Memory usage percentage
- **RAM** - Memory usage in MB

### 2. System Monitoring

#### CPU Monitoring
- **Overall usage** - Shows total CPU usage across all cores.
- **Visual bar** - Color-coded progress bar:
  - Green: < 50%
  - Yellow: 50-80%
  - Red: > 80%
- **Core count** - Displays the number of logical CPU cores.

#### Memory Monitoring
- **Used/Total** - Shows used and total RAM in GB.
- **Percentage** - Memory usage percentage.
- **Visual bar** - Color-coded progress bar.
- **Real-time updates** - Updates every refresh cycle.

#### GPU Monitoring
- **Advanced Detection** - Automatically detects and monitors GPUs on Windows and systems with NVIDIA drivers.
- **NVIDIA GPUs**: Full metrics via `nvidia-smi`, including GPU name, utilization percentage, and used/total memory.
- **Intel & AMD GPUs (Windows)**: Full metrics via PowerShell and Windows Performance Counters, including GPU name, utilization, and used/total memory.
- **Integrated GPUs**: Special handling to accurately report shared and dedicated memory.
- **Visual bars** - Color-coded progress bars for both GPU utilization and memory usage.

### 3. AI-Powered Analysis

#### Process Analysis
- **AI inspection** - Press `i` on any process to get an AI-powered analysis.
- **Gemini integration** - Uses Google's Gemini 1.5 Flash model for fast and accurate insights.
- **Analysis includes**:
  - Process purpose and function
  - Safety assessment
  - Termination recommendation
- **Scrollable results** - Long AI responses can be scrolled within the AI Assistant panel.

#### VirusTotal Integration
- **Malware scanning** - Press `t` to scan a process with VirusTotal.
- **File hash analysis** - Automatically calculates the SHA256 hash of the process executable.
- **Threat detection** - Shows malicious, suspicious, and harmless detections from over 70 antivirus engines.
- **Dedicated UI panel** - Results are displayed in a dedicated center panel with clear visual indicators.
- **Combined AI analysis** - Gemini AI analyzes the VirusTotal results to provide a comprehensive security assessment and recommendation.
- **Dual-panel display** - View raw VirusTotal data (center) and the AI's interpretation (left) simultaneously.

#### Configuration
- **API key management** - Securely store your Gemini and VirusTotal API keys.
- **Settings dialog** - Press `s` to open a dialog to configure both API keys.
- **Persistent storage** - Saves keys to a local `.env` file.

### 4. Process Management

#### Kill Process
- **Safe termination** - Press `k` to kill the selected process.
- **Confirmation dialog** - Asks for confirmation before terminating to prevent accidental kills.
- **Error handling** - Shows a clear error message if a process cannot be killed.
- **Permission aware** - Notifies the user if elevated privileges are required.

### 5. User Interface

#### Layout
- **Three-panel design**:
  - **Left**: System metrics (CPU, Memory, GPU) and the AI Assistant panel.
  - **Center**: Dedicated VirusTotal scan results panel.
  - **Right**: Main process list.
  - **Top**: Header with version and shortcuts.
  - **Bottom**: Dynamic keyboard shortcuts guide.

#### VirusTotal Panel
- **Dedicated display** - A separate, focused panel for VirusTotal scan results.
- **Visual indicators** - Color-coded threat levels with icons:
  - ⚠ RED (HIGH) - Multiple malicious detections.
  - ⚡ YELLOW (MEDIUM) - Suspicious detections.
  - ✓ GREEN (SAFE) - Clean, no threats detected.
  - ? WHITE (UNKNOWN) - Insufficient data.
- **Progress bars** - Visual display of the detection ratio.
- **Top detections** - Shows up to 5 top antivirus detections.
- **Detailed information**: Process name, PID, file path, file hash, and detection statistics.

#### Navigation
- **Arrow keys** - Navigate the process list (↑/↓).
- **Tab key** - Cycle focus between the Process List, AI Assistant, and VirusTotal panels.
- **Keyboard shortcuts**:
  - `i` - Inspect process with AI.
  - `t` - Scan with VirusTotal + AI analysis.
  - `k` - Kill process.
  - `s` - Open Settings.
  - `q` or `Esc` - Quit.

#### Visual Design
- **Color coding** - Intuitive colors for different states and threat levels.
- **Progress bars** - Visual representation of resource usage.
- **Compact layout** - Efficient use of terminal space.
- **Responsive** - Adapts to the terminal size.

## Performance

### Update Frequency
- **Process list & System metrics** - Updates every 2 seconds.
- **CPU calculation** - Uses cached values for accurate readings over time.

### Resource Usage
- **Low overhead** - Minimal CPU and memory usage during monitoring.
- **Efficient caching** - Caches only necessary data to reduce system calls.
- **Cache cleanup** - Removes stale process data every 10 seconds to prevent memory leaks.

## Limitations

### Current Limitations
1.  **Platform Support for GPU**
    - Full GPU metrics are available on Windows (all vendors) and Linux/macOS (NVIDIA only, via `nvidia-smi`).
    - GPU monitoring is not yet implemented for non-NVIDIA cards on Linux/macOS.
2.  **Process Access**
    - Accessing or killing some system-level processes may require running the application with elevated privileges (`sudo` or "Run as Administrator").
3.  **CPU Calculation**
    - The first reading for a new process may show 0% CPU as it requires one cycle to establish a baseline.

## Future Enhancements

### Planned Features
1.  **Advanced Filtering & Sorting**
    - Filter processes by name.
    - Sort by different columns (PID, Name, CPU, Memory).
2.  **Export Features**
    - Export the process list to a CSV file.
    - Save AI analysis or VirusTotal reports to a text file.
3.  **Alerts**
    - Notify the user when CPU or Memory usage exceeds a configurable threshold.
