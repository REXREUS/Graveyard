<div align="center">

# 🪦 Graveyard

### AI-Powered Process Monitor & Security Scanner

*Monitor, Analyze, and Secure Your System Processes with Intelligence*

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-blue?style=for-the-badge)](https://github.com/REXREUS/Graveyard/releases)
[![Release](https://img.shields.io/github/v/release/REXREUS/Graveyard?style=for-the-badge)](https://github.com/REXREUS/Graveyard/releases)

[Features](#-features) • [Installation](#-installation) • [Quick Start](#-quick-start) • [Documentation](#-documentation) • [Contributing](#-contributing)

---

![Graveyard Demo](docs/assets/demo.gif)
*Real-time process monitoring with AI-powered analysis*

</div>

## ✨ Features

<table>
<tr>
<td width="50%">

### 🎯 Core Features
- 📊 **Real-time Monitoring** - Live CPU & Memory tracking
- 🔄 **Auto-refresh** - Updates every second
- 🎨 **Beautiful UI** - Clean terminal interface with progress bars
- ⚡ **Fast & Lightweight** - Minimal resource usage
- 🖥️ **Cross-platform** - Windows, Linux, macOS, ARM

</td>
<td width="50%">

### 🤖 AI-Powered
- 🧠 **Smart Analysis** - Google Gemini AI integration
- 🛡️ **Malware Detection** - VirusTotal scanning
- 🔍 **Combined Intelligence** - AI + VirusTotal analysis
- 💡 **Process Insights** - Understand what's running
- 🔒 **Security Scoring** - Risk assessment

</td>
</tr>
</table>

## 🚀 Installation

### Quick Install

**Linux / macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/rexreus/Graveyard/main/setup/install.sh | bash
```

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/REXREUS/Graveyard/main/setup/install.ps1 | iex
```

### Manual Download

Download the latest binary for your platform:

| Platform | Architecture | Download |
|----------|-------------|----------|
| 🪟 Windows | x64 | [graveyard.exe](https://github.com/REXREUS/Graveyard/releases/latest/download/graveyard.exe) |
| 🐧 Linux | x64 | [graveyard](https://github.com/REXREUS/Graveyard/releases/latest/download/graveyard) |
| 🐧 Linux | ARM64 | [graveyard-arm](https://github.com/REXREUS/Graveyard/releases/latest/download/graveyard-arm) |
| 🍎 macOS | Intel | [graveyard-darwin](https://github.com/REXREUS/Graveyard/releases/latest/download/graveyard-darwin) |
| 🍎 macOS | Apple Silicon | [graveyard-darwin-arm](https://github.com/REXREUS/Graveyard/releases/latest/download/graveyard-darwin-arm) |

### Build from Source

```bash
# Clone the repository
git clone https://github.com/REXREUS/Graveyard.git
cd Graveyard

# Install dependencies
go mod download

# Build
./setup/build.sh    # Linux/Mac
setup\build.bat     # Windows
```

## ⚡ Quick Start

### 1. Run Graveyard
```bash
graveyard
```

### 2. Configure API Keys (Optional)

Press `s` in the app to open Settings, or create `.env` file:

```bash
cp .env.example .env
```

Get your free API keys:
- 🤖 [Gemini API](https://makersuite.google.com/app/apikey) - For AI analysis
- 🛡️ [VirusTotal API](https://www.virustotal.com/gui/my-apikey) - For malware scanning

Add to `.env`:
```env
GEMINI_API_KEY=your_api_key_here
VIRUSTOTAL_API_KEY=your_virustotal_api_key_here
```

### 3. Start Monitoring!

<div align="center">

| Key | Action | Description |
|:---:|--------|-------------|
| `↑` `↓` | Navigate | Move through process list |
| `i` | 🤖 AI Inspect | Analyze process with Gemini AI |
| `t` | 🛡️ Threat Scan | Check with VirusTotal + AI |
| `k` | ⚠️ Kill Process | Terminate selected process |
| `s` | ⚙️ Settings | Configure API keys |
| `q` / `Esc` | 🚪 Quit | Exit application |

</div>

## 📸 Screenshots

<div align="center">

### Main Interface
![Main Interface](https://i.imgur.com/QkpTJ3e.png)

### AI Analysis
![AI Analysis](https://i.imgur.com/bIwwqKI.png)

### VirusTotal Scan
![VirusTotal Scan](https://i.imgur.com/aUt7GG0.png)

</div>

## 🎯 Use Cases

- 🔍 **System Monitoring** - Track resource usage in real-time
- 🛡️ **Security Auditing** - Scan suspicious processes
- 🐛 **Debugging** - Identify resource-hungry applications
- 📊 **Performance Analysis** - Optimize system performance
- 🔒 **Malware Detection** - Check for threats with VirusTotal

## 💡 Why Graveyard?

| Traditional Tools | 🪦 Graveyard |
|-------------------|--------------|
| Basic process list | ✨ AI-powered insights |
| Manual analysis | 🤖 Automated intelligence |
| No security checks | 🛡️ Built-in malware scanning |
| Complex interfaces | 🎨 Clean, intuitive UI |
| Platform-specific | 🌍 Cross-platform |

## 🔧 System Requirements

- **Runtime**: No dependencies needed (standalone binary)
- **Build**: Go 1.21+ (only for building from source)
- **Terminal**: Unicode support recommended
- **Network**: Internet connection for AI features

## 🆘 Troubleshooting

<details>
<summary><b>Permission Denied when Killing Process</b></summary>

Some processes require elevated privileges:
- **Windows**: Run as Administrator
- **Linux/macOS**: `sudo graveyard`
</details>

<details>
<summary><b>AI Features Not Working</b></summary>

1. Press `s` to check API key configuration
2. Verify internet connection
3. Check logs: `graveyard.log`
4. Ensure API keys are valid
</details>

<details>
<summary><b>Binary Not Found After Install</b></summary>

- **Windows**: Restart terminal to refresh PATH
- **Linux/macOS**: Check `/usr/local/bin` is in PATH
</details>

## 📚 Documentation

<table>
<tr>
<td width="33%">

### 🚀 Getting Started
- [Quick Start](docs/QUICKSTART.md)
- [Installation](docs/INSTALL.md)
- [Project Overview](docs/PROJECT_SUMMARY.md)

</td>
<td width="33%">

### 🎨 Features
- [Feature List](docs/FEATURES.md)
- [VirusTotal Guide](docs/VIRUSTOTAL_INTEGRATION.md)
- [Panduan ID](docs/VIRUSTOTAL_INTEGRATION_ID.md)
- [GPU Monitoring](docs/GPU_MONITORING.md)

</td>
<td width="33%">

### 🛠️ Development
- [Architecture](docs/ARCHITECTURE.md)
- [Build Guide](docs/BUILD.md)
- [Contributing](CONTRIBUTING.md)
- [UI Design](docs/UI_LAYOUT.md)

</td>
</tr>
</table>

## 🔒 Security

Graveyard takes security seriously:

- 🔐 API keys stored securely in `.env` (permissions: 0600)
- 🚫 `.env` automatically excluded from git
- 🔄 Supports key rotation
- 📊 No data collection or telemetry

**Security Best Practices:**
```bash
# Never commit secrets
echo ".env" >> .gitignore

# Set proper permissions
chmod 600 .env

# Rotate keys regularly
# Monitor usage in provider dashboards
```

See [SECURITY.md](docs/SECURITY.md) for detailed information.

## 🗑️ Uninstall

**Linux / macOS:**
```bash
 curl -fsSL https://raw.githubusercontent.com/REXREUS/Graveyard/main/setup/uninstall.sh |bash
```

**Windows:**
```ps1
irm https://raw.githubusercontent.com/REXREUS/Graveyard/main/setup/uninstall.ps1 | iex
```

## 🤝 Contributing

We love contributions! Here's how you can help:

1. 🍴 Fork the repository
2. 🌿 Create a feature branch (`git checkout -b feature/amazing`)
3. 💾 Commit your changes (`git commit -m 'Add amazing feature'`)
4. 📤 Push to the branch (`git push origin feature/amazing`)
5. 🎉 Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🌟 Star History

If you find Graveyard useful, please consider giving it a star! ⭐

[![Star History Chart](https://api.star-history.com/svg?repos=REXREUS/Graveyard&type=Date)](https://star-history.com/#REXREUS/Graveyard&Date)

## 💬 Support

- 📖 [Documentation](docs/)
- 🐛 [Issue Tracker](https://github.com/REXREUS/Graveyard/issues)
- 💡 [Feature Requests](https://github.com/REXREUS/Graveyard/issues/new)
- 📧 [Contact](https://github.com/REXREUS)

---

<div align="center">

**Made with ❤️ by [REXREUS](https://github.com/REXREUS)**

[⬆ Back to Top](#-graveyard)

</div>
