# Implementation Plan

- [x] 1. Set up Go project structure and dependencies
  - Initialize Go module with `go mod init`
  - Create directory structure (cmd/graveyard, internal/ui, internal/service, internal/model, internal/app)
  - Set up go.mod with required dependencies (tview, gopsutil, generative-ai-go, godotenv)
  - Create main.go entry point in cmd/graveyard
  - _Requirements: 1.1, 6.1_

- [x] 2. Implement configuration management
  - [x] 2.1 Create Config model and ConfigService
    - Create Config struct in internal/model/config.go
    - Implement ConfigService in internal/service/config.go
    - Implement .env file loading using godotenv
    - Implement API key getter, setter, and delete methods for both Gemini and VirusTotal
    - Add API key validation logic
    - _Requirements: 5.4, 8.3, 12.1, 12.2_
  
  - [x] 2.2 Create .env.example template
    - Document required environment variables (GEMINI_API_KEY, VIRUSTOTAL_API_KEY)
    - Add default configuration values
    - _Requirements: 5.4, 12.1_

- [x] 3. Implement cross-platform process monitoring
  - [x] 3.1 Create Process model
    - Create Process struct in internal/model/process.go
    - Add String() method for formatted output
    - Define helper methods for sorting processes
    - _Requirements: 6.4_
  
  - [x] 3.2 Create ProcessService with gopsutil
    - Implement ProcessService struct in internal/service/process.go
    - Use gopsutil/process to fetch all running processes
    - Extract PID, name, CPU%, memory usage from gopsutil
    - Implement process caching for CPU calculation
    - _Requirements: 1.2, 6.1, 6.2, 6.3_
  
  - [x] 3.3 Implement GetProcessList method
    - Fetch all processes using process.Processes()
    - Get CPU percent, memory info, username for each process
    - Convert memory to MB format
    - Sort processes by CPU or memory usage
    - Handle errors for inaccessible processes
    - _Requirements: 1.2, 6.4_
  
  - [x] 3.4 Implement process kill functionality
    - Add KillProcess method using process.Kill()
    - Handle permission errors gracefully
    - Works cross-platform automatically
    - _Requirements: 7.2, 7.3_

- [x] 4. Implement system metrics monitoring
  - [x] 4.1 Create SystemMetrics models
    - Create SystemMetrics, CPUInfo, MemoryInfo, GPUInfo structs in internal/model/metrics.go
    - Add VTScanResult and ScanResult structs for VirusTotal integration
    - Add helper methods for formatting metrics and threat levels
    - _Requirements: 2.1, 2.2, 10.3, 14.1_
  
  - [x] 4.2 Create SystemService
    - Implement SystemService struct in internal/service/system.go
    - Use gopsutil/cpu for CPU usage (cpu.Percent())
    - Use gopsutil/mem for memory usage (mem.VirtualMemory())
    - Calculate usage percentages
    - _Requirements: 2.1, 2.2, 2.5_
  
  - [x] 4.3 Implement GPU detection and monitoring
    - Detect GPU availability using gopsutil or system checks
    - Retrieve GPU usage, name, and memory metrics if available
    - Return nil for GPU info if not available
    - Handle systems without GPU gracefully
    - _Requirements: 3.1, 3.2, 3.3, 3.5_

- [x] 5. Implement AI service integration
  - [x] 5.1 Create AIService
    - Implement AIService struct in internal/service/ai.go
    - Initialize Google Gemini AI client using generative-ai-go
    - Configure API key from config
    - Add error handling for API failures
    - Implement Close() method for cleanup
    - _Requirements: 4.1, 4.5, 8.2_
  
  - [x] 5.2 Implement AnalyzeProcess method
    - Create buildPrompt() helper method
    - Include process name, PID, CPU, and memory in prompt
    - Request structured response (purpose, safety, recommendation)
    - Use context.Context for timeout control
    - Call Gemini API with GenerateContent()
    - _Requirements: 4.2_
  
  - [x] 5.3 Implement AnalyzeVirusTotalResult method
    - Create buildVTAnalysisPrompt() helper method
    - Include process info, VT scan results, and detections in prompt
    - Request risk assessment, analysis, and recommendations
    - Use context.Context for timeout control
    - Call Gemini API with GenerateContent()
    - _Requirements: 11.1, 11.2, 11.3, 11.4_
  
  - [x] 5.4 Implement response formatting
    - Parse AI response text
    - Format response for terminal display with line breaks
    - Handle incomplete or malformed responses
    - Add error messages for API failures
    - _Requirements: 4.3, 4.4, 11.5_

- [x] 6. Implement state management
  - [x] 6.1 Create AppState
    - Implement AppState struct in internal/app/state.go
    - Add sync.RWMutex for thread-safe access
    - Create channels for process updates, metrics updates, AI responses, and VT responses
    - Implement state storage for processes, metrics, selected index, AI response, VT response
    - _Requirements: 1.4, 13.1_
  
  - [x] 6.2 Implement state update methods
    - Add UpdateProcesses() with mutex locking
    - Add UpdateSystemMetrics() with mutex locking
    - Add SetSelectedProcess() method
    - Add SetAIResponse() method
    - Add SetVTResponse() method
    - Add getter methods (GetProcesses, GetSelectedProcess, etc.)
    - Send updates to channels for UI consumption
    - _Requirements: 1.4, 2.5, 3.4, 11.5_

- [x] 7. Implement VirusTotal integration
  - [x] 7.1 Create VirusTotalService
    - Implement VirusTotalService struct in internal/service/virustotal.go
    - Initialize HTTP client with timeout
    - Configure API key from config
    - Add error handling for API failures and rate limits
    - _Requirements: 10.1, 10.5_
  
  - [x] 7.2 Implement file hash calculation
    - Add getProcessExecutablePath() method using gopsutil
    - Implement calculateFileHash() method using SHA256
    - Handle file access errors gracefully
    - _Requirements: 10.1_
  
  - [x] 7.3 Implement VirusTotal API integration
    - Add getFileReport() method to query VT API
    - Parse JSON response with detection statistics
    - Handle 404 errors for files not in VT database
    - Handle API authentication errors
    - _Requirements: 10.2, 10.3_
  
  - [x] 7.4 Implement ScanProcess method
    - Get process executable path
    - Calculate file hash
    - Query VirusTotal API
    - Build VTScanResult with all detection data
    - Extract top 5 detections
    - _Requirements: 10.2, 10.3, 10.4_

- [x] 8. Implement terminal UI with tview
  - [x] 8.1 Create UIManager and initialize tview application
    - Implement UIManager struct in internal/ui/manager.go with all service dependencies
    - Initialize tview.Application
    - Set up input capture for keyboard events
    - Implement Stop() and StartUpdateLoop() methods
    - Add focusedPanel tracking for multi-panel navigation
    - _Requirements: 1.1, 8.4, 13.1_
  
  - [x] 8.2 Implement main layout structure
    - Create Flex layout with tview.NewFlex()
    - Create header TextView with title and shortcuts
    - Create left Flex for system resources and AI panel (vertical)
    - Create center panel for VirusTotal results
    - Create right panel for process list
    - Create footer TextView with keyboard shortcuts help
    - Use SetBorder() and SetTitle() for visual structure
    - _Requirements: 1.1, 13.1_
  
  - [x] 8.3 Implement system resources panel
    - Create CPU TextView with dynamic content
    - Create Memory TextView with dynamic content
    - Create GPU TextView with dynamic content including GPU name
    - Implement progress bar rendering using ASCII characters with color coding
    - Add RenderSystemMetrics() method to update displays
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 3.1, 3.2_
  
  - [x] 8.4 Implement process list panel
    - Create Table widget with tview.NewTable()
    - Set table headers (PID, Name, CPU%, RAM%, RAM)
    - Implement RenderProcessList() to populate table with formatted data
    - Enable table selection and highlighting
    - Add scroll support for long lists
    - Truncate long process names for better display
    - _Requirements: 1.3, 1.5_
  
  - [x] 8.5 Implement AI assistant info panel
    - Create TextView for displaying AI responses
    - Implement text wrapping with SetWordWrap(true)
    - Enable scrolling with SetScrollable(true)
    - Add loading indicator text during AI requests
    - Implement RenderInfoPanel() method
    - Format AI response with proper line breaks and colors
    - _Requirements: 4.3, 4.4, 13.4_
  
  - [x] 8.6 Implement VirusTotal results panel
    - Create TextView for displaying VT scan results
    - Implement text wrapping and scrolling
    - Add formatVTResults() method with visual indicators
    - Display threat level with color-coded icons
    - Show detection statistics with progress bars
    - Display top 5 detections
    - Implement RenderVTPanel() method
    - _Requirements: 10.3, 10.4, 14.1, 14.2, 14.3, 14.4, 14.5_
  
  - [x] 8.7 Implement settings dialog
    - Create Form using tview.NewForm()
    - Add InputField for Gemini API key entry
    - Add InputField for VirusTotal API key entry
    - Add buttons for Save, Delete All, Cancel
    - Implement ShowSettingsDialog() method
    - Handle form submission and validation
    - Mask existing API keys for security
    - Reinitialize services on save
    - _Requirements: 5.1, 5.2, 5.3, 12.1, 12.2, 12.3, 12.4, 12.5_
  
  - [x] 8.8 Implement kill confirmation dialog
    - Create Modal with tview.NewModal()
    - Show process name and PID in message
    - Add Yes/No buttons
    - Implement ShowKillConfirmation() method returning bool
    - Handle user response
    - _Requirements: 7.1, 7.2_

- [x] 9. Implement keyboard controls and event handling
  - [x] 9.1 Implement navigation controls
    - Use SetInputCapture() on table for arrow key handling
    - Arrow up/down handled automatically by tview.Table
    - Track selected row index in AppState
    - Update selected process when row changes
    - _Requirements: 1.5, 8.1_
  
  - [x] 9.2 Implement inspect process action
    - Bind 'i' key in SetInputCapture()
    - Get selected process from AppState
    - Launch goroutine to call AIService.AnalyzeProcess()
    - Show "Loading..." message in info panel
    - Display AI response when received via channel
    - Handle errors and show error messages
    - _Requirements: 4.1, 8.2_
  
  - [x] 9.3 Implement VirusTotal scan action
    - Bind 't' key in SetInputCapture()
    - Get selected process from AppState
    - Launch goroutine to call VirusTotalService.ScanProcess()
    - Show "Scanning..." message in VT panel
    - Display VT scan results when received
    - Call AIService.AnalyzeVirusTotalResult() for AI analysis
    - Handle errors and show error messages
    - _Requirements: 10.1, 10.2, 11.1_
  
  - [x] 9.4 Implement kill process action
    - Bind 'k' key in SetInputCapture()
    - Get selected process from AppState
    - Show kill confirmation dialog with ShowKillConfirmation()
    - If confirmed, call ProcessService.KillProcess()
    - Display success or error message using ShowMessage()
    - Trigger immediate process list refresh
    - _Requirements: 7.1, 7.2, 7.3, 7.4, 7.5, 8.3_
  
  - [x] 9.5 Implement settings action
    - Bind 's' key in SetInputCapture()
    - Show settings dialog with ShowSettingsDialog()
    - Load current API keys from ConfigService
    - On save, call ConfigService.SaveAPIKey() and SaveVTAPIKey()
    - Reinitialize AIService and VirusTotalService with new keys
    - _Requirements: 5.1, 8.4, 12.2, 12.5_
  
  - [x] 9.6 Implement panel focus cycling
    - Bind Tab key in SetInputCapture()
    - Implement cycleFocus() method to switch between panels
    - Implement setFocus() method to highlight focused panel
    - Update border colors to show focus
    - Enable scrolling in focused panel
    - _Requirements: 13.2, 13.3, 13.4_
  
  - [x] 9.7 Implement exit action
    - Bind 'q' and tcell.KeyEscape in SetInputCapture()
    - Handle Escape to return focus to process list
    - Call app.Stop() to exit tview application
    - Clean up resources (close AI service, stop goroutines)
    - Exit gracefully with os.Exit(0)
    - _Requirements: 8.5, 13.5_

- [x] 10. Implement real-time monitoring loop
  - [x] 10.1 Create monitoring goroutine
    - Create StartMonitoring() function that runs in goroutine
    - Use time.Ticker for 2-second interval
    - Fetch process list using ProcessService.GetProcessList()
    - Fetch system metrics using SystemService methods
    - Send updates to AppState channels
    - Handle context cancellation for graceful shutdown
    - _Requirements: 1.4, 2.5, 3.4_
  
  - [x] 10.2 Implement UI update loop
    - Create goroutine to listen on AppState channels
    - Use select statement to handle multiple channels
    - On process update, call UIManager.RenderProcessList()
    - On metrics update, call UIManager.RenderSystemMetrics()
    - On AI response, call UIManager.RenderInfoPanel()
    - On VT response, call UIManager.RenderVTPanel()
    - Use app.QueueUpdateDraw() for thread-safe UI updates
    - _Requirements: 1.4, 2.5, 3.4, 11.5_

- [x] 11. Implement error handling and logging
  - [x] 11.1 Add error handling for process fetching
    - Check errors from ProcessService.GetProcessList()
    - Display error message in info panel on failure
    - Continue with cached data if available
    - Retry automatically on next interval
    - _Requirements: 9.1, 9.4_
  
  - [x] 11.2 Add error handling for AI service
    - Handle context timeout errors (set 30s timeout)
    - Handle API connection errors
    - Handle invalid API key errors (check for 401/403)
    - Display user-friendly error messages in info panel
    - Prompt to configure API key if missing or invalid
    - _Requirements: 4.5, 9.2, 9.3_
  
  - [x] 11.3 Add error handling for VirusTotal service
    - Handle context timeout errors (set 60s timeout)
    - Handle API connection errors
    - Handle invalid API key errors
    - Handle file not found in VT database (404)
    - Handle rate limit errors
    - Display user-friendly error messages in VT panel
    - _Requirements: 10.5_
  
  - [x] 11.4 Add error handling for process killing
    - Check errors from ProcessService.KillProcess()
    - Handle permission denied errors (EPERM)
    - Handle process not found errors
    - Display error modal with specific reason
    - Suggest running with sudo/admin privileges if needed
    - _Requirements: 7.5_
  
  - [x] 11.5 Implement logging system
    - Create logger package in internal/logger
    - Use log package or logrus for structured logging
    - Log to file in user's home directory or temp
    - Log errors with timestamps and stack traces
    - Log important operations (kill process, API calls, startup, VT scans)
    - _Requirements: 9.5_

- [x] 12. Create CLI entry point and build configuration
  - [x] 12.1 Create main.go entry point
    - Create cmd/graveyard/main.go
    - Initialize all services (Config, Process, System, AI, VirusTotal)
    - Initialize AppState and UIManager with all dependencies
    - Start monitoring goroutine
    - Start UI update goroutine
    - Run tview application
    - Handle graceful shutdown with signal handling (SIGINT, SIGTERM)
    - _Requirements: 1.1_
  
  - [x] 12.2 Create build scripts and Makefile
    - Create Makefile with build targets
    - Add target for building all platforms (Windows, Linux, macOS, ARM)
    - Add target for optimized builds with -ldflags="-s -w"
    - Add clean target to remove binaries
    - Create build output directory (bin/)
    - _Requirements: 1.1_
  
  - [x] 12.3 Create comprehensive documentation
    - Document installation instructions (download binary or build from source)
    - Document usage and keyboard shortcuts table
    - Document API key configuration (.env file for both Gemini and VirusTotal)
    - Create VIRUSTOTAL_INTEGRATION.md with detailed VT usage guide
    - Add ASCII art logo or banner
    - Add screenshots or terminal recordings
    - Document system requirements
    - Add troubleshooting section
    - _Requirements: 5.1, 8.6, 10.5, 12.1_

- [x] 13. Integration and final polish
  - [x] 13.1 Wire all components together in main.go
    - Initialize ConfigService and load config
    - Initialize ProcessService and SystemService
    - Initialize AIService with API key (if available)
    - Initialize VirusTotalService with API key (if available)
    - Initialize AppState with channels
    - Initialize UIManager with all service references
    - Start monitoring goroutine with context
    - Start UI update goroutine
    - Set up signal handling (SIGINT, SIGTERM) for graceful shutdown
    - Run tview application
    - _Requirements: All_
  
  - [x] 13.2 Test cross-platform compatibility
    - Build for Windows (GOOS=windows GOARCH=amd64)
    - Build for Linux (GOOS=linux GOARCH=amd64)
    - Build for ARM (GOOS=linux GOARCH=arm64)
    - Test process fetching on each platform
    - Test kill process on each platform
    - Verify GPU detection works correctly
    - Test with and without API keys configured
    - Test VirusTotal scanning on each platform
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5_
  
  - [x] 13.3 Optimize UI rendering and performance
    - Limit process list to top 100 by CPU/memory
    - Sort processes before rendering
    - Use app.QueueUpdateDraw() efficiently
    - Avoid unnecessary re-renders
    - Test with high process count (1000+)
    - Profile memory usage and optimize if needed
    - _Requirements: 1.4, 2.5_
  
  - [x] 13.4 Add visual polish and styling
    - Apply consistent color scheme using tcell.Color
    - Add borders to all panels with SetBorder(true)
    - Style progress bars with block characters (█░)
    - Add panel titles with SetTitle()
    - Use colors for CPU/Memory levels (green/yellow/red)
    - Add visual indicators for selected process
    - Format numbers with proper alignment
    - Add threat level icons and color coding for VT results
    - Implement focus highlighting with border colors
    - _Requirements: 1.1, 2.3, 2.4, 3.2, 14.2, 14.3, 14.4_
