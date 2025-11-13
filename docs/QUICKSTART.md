# Quick Start Guide

Get Graveyard up and running in just a few minutes!

## Step 1: Build the Application

First, ensure you have Go 1.21 or higher installed. If not, download it from [go.dev](https://go.dev/dl/).

```bash
# Clone the repository
git clone https://github.com/yourusername/process-monitor-cli.git
cd process-monitor-cli

# Install dependencies
make install

# Build the application
make build
```

## Step 2: Run Graveyard

```bash
# On Linux/macOS
./bin/graveyard

# On Windows
.\bin\graveyard.exe
```

You should now see the Graveyard interface with the process list, system metrics, and empty AI/VirusTotal panels.

## Step 3: Configure API Keys (Optional, for AI & VT features)

Press `s` to open the **Settings** dialog.
- Enter your **Gemini API Key**.
- Enter your **VirusTotal API Key**.
- Click **Save**.

Alternatively, create a `.env` file in the project root and add your keys:
```
GEMINI_API_KEY=your_gemini_api_key
VIRUSTOTAL_API_KEY=your_virustotal_api_key
```

## Step 4: Use the Application

### Basic Navigation
- **↑/↓**: Navigate the process list.
- **Tab**: Cycle focus between the Process List, AI Assistant, and VirusTotal panels.
- **i**: Inspect the selected process with AI.
- **t**: Scan the selected process with VirusTotal (requires both API keys).
- **k**: Kill the selected process (with confirmation).
- **q** or **Esc**: Quit.

### Common Tasks

**Inspect a Process with AI**
1. Select a process with ↑/↓.
2. Press `i`.
3. Read the AI analysis in the **AI Assistant** (left panel).

**Scan for Malware**
1. Select a process with ↑/↓.
2. Press `t`.
3. Review the results in the **VirusTotal Panel** (center).
4. Read the AI's interpretation in the **AI Assistant** (left panel).

**Kill a Process**
1. Select the process with ↑/↓.
2. Press `k`.
3. Confirm with "Yes" in the dialog.

## Troubleshooting

**Graveyard doesn't start.**
- Ensure the build was successful.
- Check for any error messages in `graveyard.log`.

**AI features aren't working.**
- Press `s` and verify your Gemini API key is saved.
- Check your internet connection.

**VirusTotal scan fails.**
- Press `s` and verify your VirusTotal API key.
- Check if you've exceeded the free tier rate limit (4 requests/minute).

## Next Steps

- Read the [Features Overview](FEATURES.md) to learn about all capabilities.
- Check out the [UI Layout Guide](UI_LAYOUT.md) to master the interface.
- Understand the system better by reading the [Architecture Overview](ARCHITECTURE.md).

Happy monitoring! 🚀
