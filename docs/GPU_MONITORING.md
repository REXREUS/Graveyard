# GPU Monitoring

## Current Implementation

The application includes an advanced, platform-aware GPU monitoring system capable of displaying detailed metrics for a wide range of GPUs.

### Current Status: Full Metrics

The current implementation provides comprehensive GPU metrics, including:
- **GPU Name**: The model of the graphics card.
- **GPU Utilization**: Real-time usage percentage.
- **GPU Memory**: Used and total dedicated/shared video memory.

This is achieved without requiring external Go libraries like `go-nvml` or `go-ole`. Instead, it leverages built-in system commands for maximum compatibility and zero-dependency installation.

### Implementation by Platform

#### Windows (All Vendors: NVIDIA, AMD, Intel)

- **Mechanism**: The `system.go` service uses a combination of `nvidia-smi`, `wmic`, and `powershell` commands.
- **Detection Flow**:
    1.  It first attempts to use `nvidia-smi` for detailed NVIDIA metrics.
    2.  If that fails, it falls back to `wmic` to get the GPU name and base memory information.
    3.  Finally, it uses complex `powershell` scripts to query Windows Performance Counters (`\GPU Engine(*)\Utilization Percentage`, `\GPU Adapter Memory(*)\Dedicated Usage`, etc.) to get real-time utilization and memory usage for any GPU vendor.
- **Integrated GPUs**: The PowerShell script includes special logic to correctly calculate the total memory for integrated GPUs, which use a mix of dedicated and shared system memory.

#### Linux / macOS (NVIDIA Only)

- **Mechanism**: The implementation is designed to parse the output of the `nvidia-smi` command.
- **Metrics**: Provides GPU name, utilization, and memory details for NVIDIA cards.
- **Limitation**: Monitoring for AMD and Intel GPUs on Linux/macOS is not yet implemented, as there is no universal command-line tool equivalent to `nvidia-smi` for them.

### Code Location

The core logic for GPU monitoring can be found in:
- `internal/service/system.go` within the `GetGPUUsage()` and `tryGetGPUUsageFromPerfCounter()` functions.

### UI Display

The GPU panel in the UI will show:
- The detected GPU name.
- A color-coded progress bar for GPU utilization.
- A color-coded progress bar for GPU memory usage.
- Text displaying `Used/Total` memory and usage percentages.
- A "Not detected" message if no GPU can be found.

## Future Enhancements

While the current implementation is robust, future improvements could include:
- **Linux/macOS Support for AMD/Intel**: Implementing support for other GPU vendors on non-Windows platforms, possibly by reading from `/sys/class/drm` on Linux.
- **Per-Process GPU Usage**: Adding the ability to see which specific processes are utilizing the GPU. This is a complex feature that often requires driver-level APIs.

## See Also

- [Features Overview](FEATURES.md) - All features
- [Architecture](ARCHITECTURE.md) - System design
- [Project Summary](PROJECT_SUMMARY.md) - High-level overview
