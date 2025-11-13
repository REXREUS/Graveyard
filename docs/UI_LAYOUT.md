# UI Layout Documentation

## Overview

Process Monitor CLI features a three-panel terminal UI designed for efficient, simultaneous display of system metrics, security analysis, and process data.

## Layout Structure

```
┌─────────────────────────────────────────────────────────────────────────────┐
│              Graveyard v1.0.0          [s] Settings  [q] Quit               │
├──────────────────┬──────────────────────────┬──────────────────────────────┤
│                  │                          │                              │
│  CPU Gauge       │                          │                              │
│  ┌────────────┐  │                          │      Process List            │
│  │ Usage: 45% │  │   VirusTotal Panel       │      ┌──────────────────┐   │
│  │ ████████░░ │  │   (Scan Results)         │      │ PID  Name   CPU% │   │
│  └────────────┘  │                          │      │ 1234 chrome  15% │   │
│                  │   ⚡ MEDIUM              │      │ 5678 code    12% │   │
│  Memory Gauge    │                          │      │ 9012 firefox  8% │   │
│  ┌────────────┐  │   Process: svchost.exe   │      │ ...              │   │
│  │ 8GB / 16GB │  │   PID: 5678              │      └──────────────────┘   │
│  │ ██████░░░░ │  │                          │                              │
│  └────────────┘  │   Detection: 7/70        │                              │
│                  │   ████░░░░░░░░░░░░░░     │                              │
│  GPU Gauge       │                          │                              │
│  ┌────────────┐  │   Top Detections:        │                              │
│  │ NVIDIA RTX │  │   • Kaspersky: Trojan    │                              │
│  │ Usage: 60% │  │   • Avast: Malware       │                              │
│  └────────────┘  │                          │                              │
│                  │                          │                              │
│  AI Assistant    │                          │                              │
│  ┌────────────┐  │                          │                              │
│  │ Analysis:  │  │                          │                              │
│  │ Risk: LOW  │  │                          │                              │
│  │ This is a  │  │                          │                              │
│  │ legitimate │  │                          │                              │
│  │ process... │  │                          │                              │
│  └────────────┘  │                          │                              │
│                  │                          │                              │
└──────────────────┴──────────────────────────┴──────────────────────────────┘
│  Tab: Switch   ↑↓: Nav/Scroll   i: Inspect   t: VT Scan   k: Kill          │
└─────────────────────────────────────────────────────────────────────────────┘
```

## Panel Descriptions

### Left Panel (Fixed Width: 50 chars)

Contains system metrics and the AI analysis engine.

1.  **CPU Gauge** (4 lines)
    - Current overall CPU usage percentage.
    - Number of logical cores.
    - Color-coded progress bar.

2.  **Memory Gauge** (4 lines)
    - Used vs. Total memory in GB.
    - Overall memory usage percentage.
    - Color-coded progress bar.

3.  **GPU Gauge** (7 lines)
    - GPU name/model.
    - GPU utilization percentage (if available).
    - GPU memory usage (if available).
    - Color-coded progress bars for both usage and memory.

4.  **AI Assistant Panel** (Flexible height)
    - Displays AI-powered analysis from Google Gemini.
    - Shows process inspection results (`i`) and VirusTotal scan analysis (`t`).
    - Content is scrollable using ↑↓ arrow keys when the panel is focused.

### Center Panel (Fixed Width: 45 chars)

Dedicated to displaying VirusTotal scan results.

**Features:**
- Visual threat level indicator with a color-coded icon.
- Process information (Name, PID).
- File details (Path, Hash).
- Detection statistics with a progress bar.
- A summary of the threat assessment.
- A list of the top 5 antivirus detections.

**Threat Level Icons:**
- ⚠ RED - **HIGH** threat (multiple malicious detections).
- ⚡ YELLOW - **MEDIUM** threat (suspicious detections).
- ✓ GREEN - **SAFE** (clean, no threats).
- ? WHITE - **UNKNOWN** (insufficient data).

### Right Panel (Flexible Width)

The main interactive list of running processes.

**Columns:**
- PID (Process ID)
- Name (Executable name)
- CPU% (CPU Usage)
- RAM% (Memory Usage Percentage)
- RAM (Memory Usage in MB)

**Features:**
- Shows the top 100 processes, sorted by CPU usage.
- The selected process is highlighted.
- Updates automatically every 2 seconds.

## Workflow & Navigation

### Panel Focus
- The currently focused panel is indicated by a **yellow border**.
- Press `Tab` to cycle focus between the three main panels: **Process List → AI Assistant → VirusTotal Panel**.
- Press `Esc` from the AI or VT panel to immediately return focus to the Process List.

### Scrolling
- When the **Process List** is focused, ↑↓ navigates the list.
- When the **AI Assistant** or **VirusTotal Panel** is focused, ↑↓ scrolls through their content.

### Common Workflows

1.  **Inspect a Process**:
    - Select a process in the Right Panel.
    - Press `i`.
    - Read the AI analysis in the Left Panel.

2.  **Scan for Malware**:
    - Select a process in the Right Panel.
    - Press `t`.
    - View raw scan data in the Center Panel.
    - Read the AI's interpretation and recommendation in the Left Panel.

## Color Coding

### Progress Bars
- **Green**: < 50% usage (Safe)
- **Yellow**: 50-80% usage (Moderate)
- **Red**: > 80% usage (High)

### Threat Levels
- **Red**: HIGH threat
- **Yellow**: MEDIUM threat
- **Green**: SAFE
- **White**: UNKNOWN

## Keyboard Shortcuts

| Key | Action | Description |
|---|---|---|
| `Tab` | Switch Panel | Cycle focus between Process List, AI Assistant, and VirusTotal panels. |
| `↑↓` | Navigate/Scroll | Move selection in the process list OR scroll content in a focused panel. |
| `i` | Inspect | Get an AI-powered analysis of the selected process. |
| `t` | VT Scan | Scan the selected process with VirusTotal and get a combined AI analysis. |
| `k` | Kill | Terminate the selected process (with confirmation). |
| `s` | Settings | Open the dialog to configure API keys. |
| `Esc` | Back / Quit | Return focus to the process list or quit the application. |
| `q` | Quit | Exit the application immediately. |

## Design Rationale

### Why a Three-Panel Layout?
1.  **Separation of Concerns**: Each panel has a distinct purpose: system overview (left), deep security analysis (center), and process management (right).
2.  **Simultaneous Data**: Allows for comparing raw VirusTotal data with the AI's interpretation without switching screens, improving the security workflow.
3.  **Efficient Space Use**: The layout is compact and fills the screen with relevant, non-redundant information.
