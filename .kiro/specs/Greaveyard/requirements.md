# Requirements Document

## Introduction

A cross-platform CLI-based process monitoring tool that provides real-time system resource visualization with AI-powered process analysis. The application displays running processes with CPU, Memory, and GPU usage through an interactive terminal interface, and integrates with Gemini AI to provide intelligent explanations about processes. The tool supports Windows, Linux, and ARM architectures.

## Glossary

- **Process Monitor**: The main CLI application that displays system processes
- **Process Entry**: A single running process with associated metadata (PID, name, CPU%, Memory%, GPU%)
- **AI Assistant**: The Gemini AI integration that provides process explanations
- **Settings Panel**: The UI component for managing API key configuration
- **Resource Panel**: The UI section displaying CPU and Memory metrics with progress bars
- **GPU Panel**: The UI section displaying GPU metrics with progress bars (when available)
- **Process List**: The scrollable list of all running processes
- **Info Panel**: The UI section displaying AI-generated process information

## Requirements

### Requirement 1

**User Story:** As a system administrator, I want to view all running processes in a clean terminal interface, so that I can monitor system activity without leaving the command line.

#### Acceptance Criteria

1. WHEN the user executes the application command, THE Process Monitor SHALL display a terminal-based user interface
2. THE Process Monitor SHALL retrieve all running processes from the operating system
3. THE Process Monitor SHALL display each Process Entry with PID, process name, CPU percentage, and memory percentage
4. THE Process Monitor SHALL refresh the Process List every 2 seconds with updated metrics
5. THE Process Monitor SHALL support keyboard navigation using arrow keys to select processes

### Requirement 2

**User Story:** As a developer, I want to see CPU and Memory usage visualized with progress bars, so that I can quickly identify resource-intensive processes.

#### Acceptance Criteria

1. THE Process Monitor SHALL display a Resource Panel showing overall CPU usage as a percentage
2. THE Process Monitor SHALL display a Resource Panel showing overall Memory usage as a percentage
3. THE Process Monitor SHALL render CPU usage with a visual progress bar indicator
4. THE Process Monitor SHALL render Memory usage with a visual progress bar indicator
5. THE Process Monitor SHALL update Resource Panel metrics every 2 seconds

### Requirement 3

**User Story:** As a user with a GPU, I want to see GPU usage metrics, so that I can monitor graphics processing workload.

#### Acceptance Criteria

1. WHEN a GPU is detected, THE Process Monitor SHALL display a GPU Panel
2. THE Process Monitor SHALL display GPU usage as a percentage with a visual progress bar
3. WHEN no GPU is detected, THE Process Monitor SHALL hide the GPU Panel
4. THE Process Monitor SHALL update GPU Panel metrics every 2 seconds
5. THE Process Monitor SHALL display GPU memory usage alongside GPU utilization percentage

### Requirement 4

**User Story:** As a user unfamiliar with system processes, I want to get AI-powered explanations about selected processes, so that I can understand what they do and whether they are safe.

#### Acceptance Criteria

1. WHEN the user selects a Process Entry and presses the inspect key, THE Process Monitor SHALL send the process name to the AI Assistant
2. THE AI Assistant SHALL provide an explanation including process purpose, safety assessment, and termination recommendation
3. THE Process Monitor SHALL display the AI Assistant response in the Info Panel
4. THE Process Monitor SHALL format the AI response with clear sections for readability
5. WHEN the AI Assistant request fails, THE Process Monitor SHALL display an error message in the Info Panel

### Requirement 5

**User Story:** As a user, I want to configure my Gemini API key through the application interface, so that I can enable AI features without editing configuration files manually.

#### Acceptance Criteria

1. WHEN the user presses the settings key, THE Process Monitor SHALL display a Settings Panel overlay
2. THE Settings Panel SHALL provide an input field for entering the Gemini API key
3. THE Settings Panel SHALL provide buttons to save, edit, or delete the API key
4. WHEN the user saves an API key, THE Process Monitor SHALL store the key in a .env file
5. WHEN the user closes the Settings Panel, THE Process Monitor SHALL return to the main interface

### Requirement 6

**User Story:** As a cross-platform user, I want the application to work on Windows, Linux, and ARM systems, so that I can use the same tool across different environments.

#### Acceptance Criteria

1. THE Process Monitor SHALL detect the operating system at runtime
2. WHEN running on Windows, THE Process Monitor SHALL use Windows-specific commands to retrieve process information
3. WHEN running on Linux or ARM, THE Process Monitor SHALL use Unix-specific commands to retrieve process information
4. THE Process Monitor SHALL normalize process data into a consistent format across all platforms
5. THE Process Monitor SHALL handle platform-specific process names and metrics appropriately

### Requirement 7

**User Story:** As a system administrator, I want to terminate problematic processes directly from the interface, so that I can quickly resolve system issues without using separate commands.

#### Acceptance Criteria

1. WHEN the user selects a Process Entry and presses the kill key, THE Process Monitor SHALL display a confirmation dialog
2. THE confirmation dialog SHALL show the process name and PID to be terminated
3. WHEN the user confirms termination, THE Process Monitor SHALL attempt to kill the selected process
4. WHEN process termination succeeds, THE Process Monitor SHALL display a success message and remove the process from the list
5. WHEN process termination fails, THE Process Monitor SHALL display an error message with the failure reason

### Requirement 8

**User Story:** As a user, I want intuitive keyboard controls, so that I can navigate and interact with the application efficiently.

#### Acceptance Criteria

1. THE Process Monitor SHALL support arrow up/down keys for navigating the Process List
2. THE Process Monitor SHALL support 'i' key for inspecting the selected process
3. THE Process Monitor SHALL support 'k' key for killing the selected process
4. THE Process Monitor SHALL support 's' key for opening the Settings Panel
5. THE Process Monitor SHALL support 'q' or 'Esc' key for exiting the application
6. THE Process Monitor SHALL display a help bar showing available keyboard shortcuts

### Requirement 9

**User Story:** As a user, I want the application to handle errors gracefully, so that temporary issues don't crash the monitoring tool.

#### Acceptance Criteria

1. WHEN process retrieval fails, THE Process Monitor SHALL display an error message and retry after 2 seconds
2. WHEN the AI Assistant API is unavailable, THE Process Monitor SHALL display a connection error message
3. WHEN the API key is invalid or missing, THE Process Monitor SHALL prompt the user to configure it in Settings
4. THE Process Monitor SHALL continue displaying cached process data when refresh operations fail
5. THE Process Monitor SHALL log errors to a log file for debugging purposes

### Requirement 10

**User Story:** As a security-conscious user, I want to scan running processes with VirusTotal, so that I can identify potential malware threats using multiple antivirus engines.

#### Acceptance Criteria

1. WHEN the user selects a Process Entry and presses the scan key, THE Process Monitor SHALL calculate the SHA256 hash of the process executable
2. THE Process Monitor SHALL query the VirusTotal API with the file hash
3. THE Process Monitor SHALL display scan results including malicious, suspicious, harmless, and undetected counts
4. THE Process Monitor SHALL show the top 5 antivirus engine detections when threats are found
5. WHEN the VirusTotal API key is missing or invalid, THE Process Monitor SHALL display an error message prompting configuration

### Requirement 11

**User Story:** As a user analyzing security threats, I want AI-powered analysis of VirusTotal scan results, so that I can understand the risk level and get actionable recommendations.

#### Acceptance Criteria

1. WHEN a VirusTotal scan completes, THE Process Monitor SHALL send the scan results to the AI Assistant
2. THE AI Assistant SHALL provide a risk assessment based on detection statistics
3. THE AI Assistant SHALL explain whether detections are likely false positives
4. THE AI Assistant SHALL provide recommendations on whether to terminate the process
5. THE Process Monitor SHALL display the AI analysis in a dedicated panel

### Requirement 12

**User Story:** As a user, I want to manage multiple API keys through the settings interface, so that I can configure both Gemini and VirusTotal services easily.

#### Acceptance Criteria

1. THE Settings Panel SHALL provide separate input fields for Gemini API key and VirusTotal API key
2. WHEN the user saves API keys, THE Process Monitor SHALL store both keys in the .env file
3. THE Settings Panel SHALL display masked versions of existing API keys for security
4. WHEN the user deletes API keys, THE Process Monitor SHALL remove all keys from configuration
5. THE Process Monitor SHALL reinitialize services when new API keys are saved

### Requirement 13

**User Story:** As a user, I want to navigate between multiple information panels, so that I can view process list, AI analysis, and VirusTotal results simultaneously.

#### Acceptance Criteria

1. THE Process Monitor SHALL display three main panels: process list, AI Assistant, and VirusTotal results
2. WHEN the user presses the tab key, THE Process Monitor SHALL cycle focus between panels
3. THE Process Monitor SHALL highlight the focused panel with a distinct border color
4. WHEN a panel is focused, THE Process Monitor SHALL enable scrolling with arrow keys
5. WHEN the user presses escape, THE Process Monitor SHALL return focus to the process list

### Requirement 14

**User Story:** As a user, I want visual threat indicators in VirusTotal results, so that I can quickly assess the security risk level.

#### Acceptance Criteria

1. THE Process Monitor SHALL classify threats as HIGH, MEDIUM, SAFE, or UNKNOWN based on detection ratios
2. WHEN threat level is HIGH, THE Process Monitor SHALL display results with red color and warning icon
3. WHEN threat level is MEDIUM, THE Process Monitor SHALL display results with yellow color and alert icon
4. WHEN threat level is SAFE, THE Process Monitor SHALL display results with green color and checkmark icon
5. THE Process Monitor SHALL display a visual progress bar showing detection percentage
