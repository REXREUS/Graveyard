# Contributing to Process Monitor CLI

Thank you for your interest in contributing! This document provides guidelines for contributing to the project.

## Development Setup

1. **Install Go 1.21+**
   - Download from: https://go.dev/dl/

2. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/process-monitor-cli.git
   cd process-monitor-cli
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Set up environment**
   ```bash
   cp .env.example .env
   # Add your GEMINI_API_KEY
   ```

5. **Build and run**
   ```bash
   make build
   ./bin/graveyard
   ```

## Project Structure

See [ARCHITECTURE.md](docs/ARCHITECTURE.md) for detailed architecture information.

## Making Changes

### Code Style

- Follow standard Go conventions
- Use `gofmt` to format code
- Add comments for exported functions
- Keep functions small and focused

### Testing

Before submitting:

```bash
# Format code
go fmt ./...

# Run tests (when available)
go test ./...

# Build for all platforms
make build-all
```

### Commit Messages

Use clear, descriptive commit messages:

```
Add GPU monitoring support

- Implement GPU detection using gopsutil
- Add GPU panel to UI
- Update documentation
```

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

### PR Guidelines

- Describe what your PR does
- Reference any related issues
- Include screenshots for UI changes
- Ensure builds pass on all platforms

## Feature Requests

Open an issue with:
- Clear description of the feature
- Use cases and benefits
- Possible implementation approach

## Bug Reports

Include:
- Steps to reproduce
- Expected behavior
- Actual behavior
- System information (OS, Go version)
- Relevant logs from `graveyard.log`

## Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow

## Questions?

Open an issue with the "question" label or reach out to the maintainers.

Thank you for contributing! 🎉

## See Also

- [Architecture](docs/ARCHITECTURE.md) - System design and components
- [Build Instructions](docs/BUILD.md) - Building for different platforms
- [Security](docs/SECURITY.md) - Security considerations
- [Project Summary](docs/PROJECT_SUMMARY.md) - High-level overview
