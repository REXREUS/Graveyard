# Changelog

## v1.2.1 - Panel Focus & Scrolling Fix (Current Build)

### 🐛 Bug Fixes

#### Fixed AI Assistant & VirusTotal Panel Scrolling
- **Panel focus system** - Added Tab key to cycle focus between panels
- **Visual focus indicator** - Focused panel shows yellow border
- **Scrollable panels** - AI Assistant and VirusTotal panels now properly scrollable
- **Keyboard navigation**:
  - `Tab` - Cycle focus: Process List → AI Assistant → VirusTotal → Process List
  - `↑↓` - Navigate in process list OR scroll in focused panel
  - `Esc` - Return focus to process list (or quit if already focused)
- **Updated UI hints** - Panel titles show "(Tab to focus, ↑↓ to scroll)"
- **Updated footer** - Shows Tab key functionality

#### Fixed Settings Dialog Button Interaction
- **Input capture management** - Disabled global key bindings when dialogs are open
- **Button focus** - Save, Delete All, and Cancel buttons now properly clickable
- **Modal focus** - All modals (Settings, Kill Process, Messages) now receive focus correctly
- **Key binding restoration** - Global key bindings automatically restored when dialogs close

### 📚 Documentation

- **New PROJECT_SUMMARY.md** - Comprehensive project overview
  - High-level project description
  - Key features and capabilities
  - Technology stack information
  - Use cases and platform support
  - Links to all documentation
- **Updated README.md** - Added comprehensive documentation index
  - Organized by category (Getting Started, Features, Development, Reference)
  - Clear language indicators for translated docs
  - Improved navigation structure
- **Added cross-references** - All documentation files now include "See Also" sections
  - Easy navigation between related topics
  - Consistent link formats across all files
  - Better documentation discoverability
- **Updated UI_LAYOUT.md** - Added panel focus and scrolling documentation
- **Updated keyboard shortcuts** - Documented Tab and Esc key behavior
- **New SECURITY.md** - Comprehensive security documentation
  - Explains why API keys are stored in plaintext
  - Best practices for protecting API keys
  - Threat model and security measures
  - What to do if keys are compromised
  - Alternative security approaches

---

## v1.2.0 - Dedicated VirusTotal UI Panel

### 🎨 Major UI Redesign

#### New Dedicated VirusTotal Panel
- **Three-panel layout** - Added center panel specifically for VirusTotal results
- **Visual threat indicators** - Color-coded icons for threat levels:
  - ⚠ RED (HIGH) - Multiple malicious detections
  - ⚡ YELLOW (MEDIUM) - Suspicious detections  
  - ✓ GREEN (SAFE) - Clean, no threats
  - ? WHITE (UNKNOWN) - Insufficient data
- **Progress bars** - Visual detection ratio display (e.g., 7/70 engines)
- **Structured information display**:
  - Process Information (Name, PID)
  - File Details (Path, Hash)
  - Detection Results (with progress bar)
  - Summary (threat assessment)
  - Top Detections (up to 5 engines)
- **Real-time updates** - Panel updates during scan process
- **Better space utilization** - Fills empty space between metrics and process list

#### Enhanced Dual-Panel Analysis
- **Separated concerns** - VirusTotal results in center, AI analysis in left panel
- **Improved readability** - Each panel focuses on specific information
- **Synchronized updates** - Both panels update in real-time during scan
- **Better workflow** - Easier to compare VT data with AI recommendations

### 🔧 Technical Improvements

#### State Management
- **New VT channel** - Dedicated channel for VirusTotal panel updates
- **Improved update loop** - Handles VT panel updates separately from AI panel
- **Better separation** - VT and AI responses managed independently

#### Model Enhancements
- **Extended VTScanResult** - Added TotalEngines and Scans map fields
- **ScanResult struct** - Individual engine results with Detected flag and Result string
- **Improved threat calculation** - More accurate threat level based on percentage
- **Better summary generation** - Context-aware summary messages

## v1.1.0 - VirusTotal Integration

### 🛡️ New Features

#### VirusTotal Integration
- **Malware scanning** - Press `t` to scan any process with VirusTotal
- **File hash analysis** - Automatically calculates SHA256 hash of process executable
- **Multi-engine detection** - Aggregates results from 70+ antivirus engines
- **Threat classification** - Shows Malicious, Suspicious, Harmless, and Undetected counts
- **Detection details** - Displays top 5 antivirus detections with threat names
- **Threat levels** - Color-coded risk levels (HIGH/MEDIUM/SAFE/UNKNOWN)

#### AI-Powered Security Analysis
- **Combined analysis** - Gemini AI analyzes VirusTotal scan results
- **Risk assessment** - AI evaluates overall security risk and confidence level
- **False positive detection** - Identifies potential false positives
- **Actionable recommendations** - Provides clear guidance on whether to terminate process
- **Context-aware analysis** - Considers process type and behavior patterns

#### Enhanced Configuration
- **Dual API key support** - Manage both Gemini and VirusTotal API keys
- **Updated settings dialog** - Configure both API keys in one place
- **Persistent storage** - Both keys saved securely to `.env` file
- **Key validation** - Validates both API key formats

### 🎨 UI Improvements

#### Keyboard Shortcuts
- **New `t` key** - Scan process with VirusTotal + AI analysis
- **Updated footer** - Shows new keyboard shortcut in UI

#### AI Assistant Panel
- **Enhanced display** - Shows VirusTotal results with color-coded threat levels
- **Structured output** - Clear separation between VT results and AI analysis
- **Progress indicators** - Shows scanning and analysis progress
- **Better error messages** - Specific error messages for VT and AI failures

### 🔧 Backend Improvements

#### New Services
- **VirusTotalService** - Complete VirusTotal API integration
  - File hash calculation (SHA256)
  - Process executable path resolution
  - VT API queries with proper error handling
  - Result parsing and aggregation
- **Enhanced AIService** - New method for analyzing VT results
  - `AnalyzeVirusTotalResult()` - Analyzes VT scan data with AI
  - Comprehensive prompt engineering for security analysis
  - Structured output format

#### New Models
- **VTScanResult** - Stores VirusTotal scan results
  - Process information (name, PID, path, hash)
  - Detection statistics (malicious, suspicious, harmless, undetected)
  - Top detections list
  - Threat level calculation
  - Summary generation
- **Config** - Enhanced configuration model
  - Gemini API key
  - VirusTotal API key

### 📚 Documentation

#### New Documentation
- **VIRUSTOTAL_INTEGRATION.md** - Complete integration guide
  - Feature overview
  - Setup instructions
  - Usage examples
  - Result interpretation
  - Best practices
  - Troubleshooting
  - API rate limits
  - Security considerations

#### Updated Documentation
- **README.md** - Added VirusTotal features and configuration
- **FEATURES.md** - Documented VirusTotal and AI analysis features
- **.env.example** - Added VirusTotal API key template

### 🔒 Security Considerations

- File hashes (not actual files) are sent to VirusTotal
- API keys stored securely in `.env` file with 0600 permissions
- Rate limiting awareness (4 requests/minute for free tier)
- Privacy implications documented

### 📋 Requirements

- **VirusTotal API key** - Get from https://www.virustotal.com/gui/my-apikey
- **Gemini API key** - Get from https://makersuite.google.com/app/apikey
- **Internet connection** - Required for both VT and AI features
- **File access permissions** - May require elevated privileges for some processes

### 🐛 Known Limitations

1. **API Rate Limits**
   - Free tier: 4 requests/minute, 500/day
   - Requires waiting between scans

2. **File Access**
   - Some system processes may require administrator privileges
   - Protected processes may not be accessible

3. **Database Coverage**
   - New or custom applications may not be in VT database
   - Requires manual upload for unknown files

### 🚀 Future Enhancements

- File upload for unknown files
- Historical scan results caching
- Batch scanning multiple processes
- Export scan reports to JSON/CSV
- Real-time monitoring with automatic scanning

---

## v1.0.2 - Latest Updates

### UI Improvements

#### Process List
- **Ultra-compact layout** - Process list now uses minimal horizontal space
- **Fixed column widths** - Columns: PID (5 chars), Name (15 chars), CPU% (4 chars), RAM% (4 chars), RAM (5 chars)
- **Truncated long names** - Process names longer than 15 characters are truncated with "..."
- **Better formatting** - All columns have consistent, compact width
- **Adjusted panel width** - Left panel increased to 50 characters for better readability
- **Smart CPU% formatting** - Shows percentage with proper alignment:
  - `0.1%` for values < 1%
  - `5.2%` for values 1-10%
  - `15%` for values >= 10%
- **RAM percentage** - Shows RAM usage percentage for each process
- **RAM size** - Shows actual RAM usage in MB (e.g., "42M")
- **Sorted by CPU** - Processes sorted by CPU usage (highest first)
- **Example format**: `11676 msedgewebvie...  0.0%  5%   42M`
  - PID: 11676
  - Name: msedgewebview2.exe (truncated)
  - CPU%: 0.0% (CPU usage percentage)
  - RAM%: 5% (RAM usage percentage)
  - RAM: 42M (RAM size in MB)

#### AI Assistant Panel
- **Scrollable content** - AI responses can now be scrolled if they're too long
- **Word wrap enabled** - Long text wraps properly without cutting off
- **Better title** - Shows "(↑↓ to scroll)" hint in the title
- **Improved error messages** - More user-friendly error messages instead of technical errors

#### GPU Monitoring
- **NVIDIA GPU metrics** - ✅ Full support for NVIDIA GPUs via nvidia-smi
  - GPU usage percentage with progress bar
  - Memory usage (used/total) with automatic MB/GB formatting
  - Memory percentage with progress bar
- **Intel GPU metrics** - ✅ Full support for Intel integrated GPUs
  - GPU usage percentage with progress bar via Performance Counters
  - Intelligent memory detection (dedicated + shared system memory)
  - Automatic detection of shared memory pool (typically 50% of system RAM)
  - Current memory usage tracking via multiple methods
- **AMD GPU metrics** - ✅ Full support for AMD GPUs via Windows Performance Counters
  - GPU usage percentage with progress bar
  - Memory detection via WMIC and CIM
  - Dedicated and shared memory usage tracking
- **Universal memory formatting** - ✅ Intelligent size display
  - Shows in GB for memory >= 1GB (e.g., "2.5GB/5.9GB")
  - Shows in MB for memory < 1GB (e.g., "512MB/768MB")
  - Automatic unit selection for best readability
- **Universal GPU detection** - ✅ Detects all GPU types using multiple methods:
  - nvidia-smi for NVIDIA GPUs (highest priority)
  - WMIC and CIM for GPU name and dedicated memory
  - PowerShell WMI for comprehensive memory info (dedicated + shared)
  - Performance Counters for real-time usage metrics
  - Automatic fallback between methods
- **Integrated GPU support** - ✅ Special handling for Intel/AMD integrated GPUs
  - Detects shared system memory allocation
  - Calculates total available memory (typically 50% of system RAM)
  - Tracks both dedicated and shared memory usage
- **GPU name display** - Shows actual GPU name (e.g., "Intel(R) UHD Graphics", "NVIDIA GeForce RTX 3060", "AMD Radeon RX 6800")
- **Improved parsing** - Better string parsing to handle different GPU name formats
- **Status messages** - Shows GPU name when detected, "Not detected" when unavailable
- **Expanded GPU panel** - Increased height to 7 lines to show all metrics
- **Clean layout** - GPU panel integrated between Memory and AI Assistant

#### CPU Monitoring
- **Percentage display** - CPU usage now shows percentage in cyan color
- **Core count** - Shows number of CPU cores (e.g., "Usage: 45.2% (8 cores)")
- **Better visibility** - Highlighted percentage for easier reading

### Backend Improvements

#### GPU Detection & Monitoring
- **nvidia-smi integration** - Uses nvidia-smi for NVIDIA GPUs to get real-time metrics
  - GPU utilization percentage
  - Memory used and total (in MB)
  - Automatic fallback to WMIC if nvidia-smi not available
- **WMIC/CIM integration** - Uses Windows WMIC and CIM to query video controller information
  - GPU name detection for all GPU types
  - AdapterRAM detection for dedicated memory
  - CSV format parsing for reliable data extraction
  - CIM queries for enhanced compatibility
- **PowerShell WMI queries** - Comprehensive memory detection for integrated GPUs
  - Detects dedicated memory (AdapterRAM)
  - Calculates shared system memory allocation
  - Determines total available memory (dedicated + shared)
  - Special handling for integrated GPUs (Intel/AMD)
- **PowerShell Performance Counters** - Real-time usage metrics for Intel/AMD GPUs
  - GPU Engine Utilization Percentage for 3D workloads
  - GPU Adapter Memory Dedicated Usage for dedicated memory tracking
  - GPU Adapter Memory Total Committed for total usage
  - GPU Adapter Memory Shared Usage for shared memory tracking
  - Multiple fallback methods for maximum compatibility
- **Multi-method detection** - Tries multiple methods in order:
  1. nvidia-smi (NVIDIA GPUs - most accurate)
  2. WMIC with AdapterRAM (all GPUs - basic info)
  3. PowerShell WMI (comprehensive memory info)
  4. PowerShell Performance Counters (real-time usage)
  5. Simple WMIC name query (fallback)
- **Intelligent memory calculation** - For integrated GPUs:
  - Detects if GPU has less than 128MB dedicated memory
  - Automatically calculates shared memory pool (50% of system RAM)
  - Provides accurate total available memory
- Detects presence of GPU (NVIDIA, AMD, Intel, etc.)
- Stores GPU name and metrics in system metrics
- Helper functions for parsing nvidia-smi, WMIC, and PowerShell output

#### CPU Calculation
- **Caching mechanism** - Caches process objects for accurate CPU readings
- **Interval-based** - Uses previous readings to calculate actual CPU usage
- **Cache cleanup** - Removes stale processes every 10 seconds
- **First-run handling** - Initializes CPU tracking on first detection

#### System Service
- Added `os/exec` for running system commands
- Improved error handling for GPU detection
- Maintains compatibility with existing CPU/Memory monitoring

### Known Limitations

1. **GPU Metrics accuracy** - Different methods have different accuracy levels
   - NVIDIA GPUs: Most accurate via nvidia-smi ✅
   - Intel/AMD GPUs: Uses Windows Performance Counters (may have slight delays)
   - Performance Counter queries may require administrator privileges for some metrics

2. **AI Model** - Using `gemini-1.5-flash-latest`
   - Requires valid API key from Google AI Studio
   - May need adjustment based on API availability

### Requirements

- **For NVIDIA GPU monitoring**: nvidia-smi must be installed (comes with NVIDIA drivers)
- **For Intel/AMD GPU monitoring**: Windows Performance Counters (built-in to Windows)
- **For best GPU metrics**: Run as administrator (optional, but recommended)

### Performance Notes

- GPU metrics are queried using PowerShell for Intel/AMD GPUs
- First query may take slightly longer (~1-2 seconds)
- Subsequent queries are cached and faster
- NVIDIA GPUs use native nvidia-smi which is very fast

See `GPU_MONITORING.md` for detailed implementation guide.

