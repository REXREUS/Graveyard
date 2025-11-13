# VirusTotal Integration Guide

## Overview

Graveyard integrates with VirusTotal to provide a seamless, two-step security analysis workflow. When you scan a process, Graveyard not only queries VirusTotal for malware detection results but also uses Google Gemini AI to interpret the results and provide a clear, actionable assessment.

## How It Works

The integration combines automated malware detection with AI-powered analysis:

1.  **VirusTotal Scan**: Graveyard calculates the SHA256 hash of a process's executable and queries the VirusTotal API v3 database.
2.  **AI Interpretation**: Google Gemini AI 2.5 Flash model analyzes the VirusTotal data and provides a security recommendation.

This dual approach gives you both the raw data and expert-level interpretation, making it easier to understand complex scan results.

## Key Features

-   **Multi-Engine Detection**: Leverages over 70 antivirus engines through a single VirusTotal query.
-   **Threat Categorization**: Automatically classifies results into Malicious, Suspicious, Harmless, or Undetected.
-   **Risk-Based Analysis**: The AI provides a risk level assessment and specific recommendations for action.
-   **Combined Workflow**: Results from both VirusTotal and AI are displayed simultaneously in the UI for a comprehensive view.

## Setting Up

### 1. Get a VirusTotal API Key

1.  Visit [VirusTotal.com](https://www.virustotal.com/) and create a free account.
2.  Go to your [API key section](https://www.virustotal.com/gui/my-apikey) to find your private API key.

### 2. Get a Gemini API Key

1.  Visit [Google AI Studio](https://aistudio.google.com/).
2.  Click "Get API Key" and create a key in a new or existing project.

### 3. Configure Graveyard

You can configure your API keys in two ways:

**Option A: Settings UI (Recommended)**
1.  Run Graveyard.
2.  Press `s` to open the Settings dialog.
3.  Paste your keys into the respective fields.
4.  Click "Save".

**Option B: `.env` File**
1.  In the project directory, create a `.env` file (or edit if it exists) from the template (`.env.example`).
2.  Add your keys:
    ```
    GEMINI_API_KEY=your_gemini_api_key_here
    VIRUSTOTAL_API_KEY=your_virustotal_api_key_here
    ```

## Scanning a Process

1.  **Select a Process**: Navigate to the process you want to analyze using the arrow keys in the Process List.
2.  **Start the Scan**: Press `t`.
3.  **View the Results**:
    -   **Center Panel (VirusTotal Data)**: Shows the raw scan data, including detection statistics, threat level, and the top antivirus detections.
    -   **Left Panel (AI Analysis)**: After the scan completes, the AI will provide its interpretation, risk assessment, and recommended actions.

## Understanding the Results

### VirusTotal Panel (Center)

This panel presents the raw data from VirusTotal.

-   **Threat Level & Icon**: A visual indicator of the overall threat.
    -   ⚠ **RED**: High threat (multiple malicious detections).
    -   ⚡ **YELLOW**: Medium threat (suspicious detections).
    -   ✓ **GREEN**: Safe (no threats detected).
    -   ? **WHITE**: Unknown (insufficient data).

-   **Process Information**: The name and PID of the scanned process.
-   **File Details**: The full file path and a truncated hash of the executable.
-   **Detection Results**: A progress bar and count showing how many of the 70+ engines flagged the file as malicious or suspicious.
-   **Summary**: A concise text summary of the findings.
-   **Top Detections**: A list of up to 5 of the most notable detections, including the engine name and the specific threat name it reported.

### AI Assistant Panel (Left)

This panel, powered by Gemini, provides the expert analysis.

The AI will provide:
1.  **Risk Assessment**: An overall evaluation of the threat level.
2.  **Analysis of Detections**: An explanation of what the detections mean, including a discussion of potential false positives.
3.  **Recommended Actions**: Clear guidance on what you should do next (e.g., terminate the process, quarantine the file, or continue monitoring).
4.  **Additional Context**: Information about the process type and its potential significance on the system.

## Example Workflow

### Scenario: A Suspicious System Process

1.  You notice `svchost.exe` (a Windows system process) is using more CPU than usual.
2.  You select it and press `t` to scan.
3.  **VirusTotal Panel Shows**:
    -   **Threat Level**: ⚡ (Medium)
    -   **Detections**: "Engines: 7 / 70 detected"
    -   **Top Detection**: "Kaspersky: Trojan.Win32.Generic"
4.  **AI Panel Shows**:
    -   **Risk Assessment**: "Medium-Low (likely false positive)"
    -   **Analysis**: "Svchost.exe is a critical Windows process. The detections are likely false positives due to its complex, multi-service nature and network activity. Many antivirus engines flag legitimate system files when they exhibit certain behaviors."
    -   **Recommendation**: "Do not terminate. This is a system process. Instead, use Windows Task Manager to identify and terminate the specific service (svchost.exe -k ...) that's causing the high CPU usage."

This demonstrates how the AI helps you avoid a critical system error by contextualizing the data.

## Best Practices

-   **Scan Unknown Processes**: Always scan processes you don't recognize before taking action.
-   **Trust, But Verify**: Treat VirusTotal detections as a starting point, not a verdict. Use the AI analysis to understand the context.
-   **Consider the Source**: System processes are known for false positives. Be extra cautious when dealing with `svchost.exe`, `lsass.exe`, `csrss.exe`, etc.
-   **Monitor Trends**: If a previously safe process starts showing suspicious detections, it may be compromised.
-   **Mind the Rate Limits**: The free VirusTotal API key is limited to 4 requests per minute.

## API Rate Limits

-   **Free Tier**: 4 requests/minute, 500/day, 15,500/month.
-   **Premium Tier**: Higher limits and additional features.

## Troubleshooting

-   **"File not found in VirusTotal database"**: This is common for custom or newly developed applications. The file hasn't been scanned before.
-   **"Rate limit exceeded"**: Wait a minute before scanning again. Be more selective about what you scan.
-   **"AI Service not available"**: Ensure your Gemini API key is configured correctly in Settings.

## Privacy

-   **No File Upload**: Graveyard only sends a SHA256 hash to VirusTotal; the actual file is never uploaded.
-   **Hash Sharing**: VirusTotal may share hashes with the broader security community to improve detection capabilities.

## See Also

-   [Features Overview](FEATURES.md) - Complete feature list
-   [UI Layout](UI_LAYOUT.md) - Understanding the user interface
-   [Security Considerations](SECURITY.md) - How we protect your API keys
