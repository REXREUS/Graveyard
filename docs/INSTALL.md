# Installation Guide

There are two main ways to install and run the Process Monitor CLI: by building from source or by downloading a pre-compiled binary from the project's releases page.

## Option 1: Building from Source (Recommended)

This method ensures you have the latest version and is the standard way to install Go applications.

### Prerequisites
- **Go**: Version 1.21 or higher.
- **Git**: For cloning the repository.
- **Make**: (Optional) For using simplified build commands.

### Installation Steps

1.  **Set up Go**
    If you don't have Go installed, download it from the [official Go website](https://go.dev/dl/). Follow the instructions for your operating system. Verify the installation by opening a new terminal and running:
    ```bash
    go version
    ```

2.  **Clone the Repository**
    ```bash
    git clone https://github.com/yourusername/process-monitor-cli.git
    cd process-monitor-cli
    ```

3.  **Install Dependencies**
    This command downloads the required Go modules.
    ```bash
    make install
    # Or manually:
    # go mod download
    ```

4.  **Build the Application**
    This compiles the source code into a single executable file.
    ```bash
    make build
    # Or manually:
    # go build -o bin/graveyard cmd/graveyard/main.go
    ```

## Option 2: Using a Pre-compiled Binary

For a quicker setup, you can download a binary for your operating system from the project's **Releases** page on GitHub.

1.  Navigate to the [Releases Page](https://github.com/yourusername/process-monitor-cli/releases).
2.  Find the latest release and download the appropriate archive for your OS and architecture (e.g., `graveyard-windows-amd64.zip`, `graveyard-linux-amd64.tar.gz`).
3.  Extract the archive.
4.  (Optional) Move the executable to a directory in your system's `PATH` (e.g., `/usr/local/bin` or `C:\Windows\System32`) to make it runnable from any location.

## Running the Application

After installation, you can run the application from your terminal.

- If you built from source:
  ```bash
  # On Linux/macOS
  ./bin/graveyard

  # On Windows
  .\bin\graveyard.exe
  ```

- If you moved the binary to your `PATH`, you can simply run:
  ```bash
  graveyard
  ```

## Configuration (API Keys)

To use the AI and VirusTotal features, you need to configure your API keys.

1.  **Create the `.env` file**:
    In the project's root directory, copy the example file:
    ```bash
    cp .env.example .env
    ```

2.  **Add your keys**:
    Edit the `.env` file and paste your API keys:
    ```
    GEMINI_API_KEY=your_gemini_api_key_here
    VIRUSTOTAL_API_KEY=your_virustotal_api_key_here
    ```

Alternatively, you can configure the keys directly within the application by pressing `s` to open the Settings dialog.

## Troubleshooting

- **"Permission denied"**: You may need to make the binary executable on Linux/macOS (`chmod +x bin/graveyard`) or run with elevated privileges (`sudo`) for certain features like killing system processes.
- **AI/VT features not working**: Ensure your API keys are correct in the `.env` file or have been saved correctly via the Settings menu. Check `graveyard.log` for any API-related errors.

## See Also

- [Quick Start Guide](QUICKSTART.md) - Get running in 5 minutes.
- [Build Instructions](BUILD.md) - For advanced compilation options like cross-platform builds.
