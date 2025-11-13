package service

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/yourusername/process-monitor-cli/internal/model"
)

type SystemService struct {
	hasGPU bool
}

func NewSystemService() *SystemService {
	return &SystemService{
		hasGPU: false,
	}
}

func (s *SystemService) GetCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, nil
}

func (s *SystemService) GetMemoryUsage() (*model.MemoryInfo, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &model.MemoryInfo{
		UsedGB:     float64(v.Used) / 1024 / 1024 / 1024,
		TotalGB:    float64(v.Total) / 1024 / 1024 / 1024,
		Percentage: v.UsedPercent,
	}, nil
}

func (s *SystemService) GetGPUUsage() (*model.GPUInfo, error) {
	gpuInfo := &model.GPUInfo{
		Available:     false,
		Name:          "",
		Usage:         0.0,
		MemoryUsedMB:  0,
		MemoryTotalMB: 0,
	}
	
	if runtime.GOOS == "windows" {
		// Try nvidia-smi first for NVIDIA GPUs
		cmd := exec.Command("nvidia-smi", "--query-gpu=name,utilization.gpu,memory.used,memory.total", "--format=csv,noheader,nounits")
		output, err := cmd.Output()
		if err == nil {
			outputStr := strings.TrimSpace(string(output))
			parts := strings.Split(outputStr, ",")
			if len(parts) >= 4 {
				gpuInfo.Available = true
				gpuInfo.Name = strings.TrimSpace(parts[0])
				
				// Parse GPU usage
				if usage, err := parseFloat(strings.TrimSpace(parts[1])); err == nil {
					gpuInfo.Usage = usage
				}
				
				// Parse memory used (in MB)
				if memUsed, err := parseUint64(strings.TrimSpace(parts[2])); err == nil {
					gpuInfo.MemoryUsedMB = memUsed
				}
				
				// Parse memory total (in MB)
				if memTotal, err := parseUint64(strings.TrimSpace(parts[3])); err == nil {
					gpuInfo.MemoryTotalMB = memTotal
				}
				
				return gpuInfo, nil
			}
		}
		
		// Try to get GPU info using WMIC with more details
		cmd = exec.Command("wmic", "path", "win32_VideoController", "get", "name,AdapterRAM", "/format:csv")
		output, err = cmd.Output()
		if err == nil {
			outputStr := string(output)
			lines := strings.Split(outputStr, "\n")
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				// Skip empty lines and header
				if trimmed == "" || strings.Contains(strings.ToLower(trimmed), "node") {
					continue
				}
				
				// Parse CSV format: Node,AdapterRAM,Name
				parts := strings.Split(trimmed, ",")
				if len(parts) >= 3 {
					gpuName := strings.TrimSpace(parts[2])
					if gpuName != "" && !strings.Contains(strings.ToLower(gpuName), "name") {
						gpuInfo.Available = true
						gpuInfo.Name = gpuName
						
						// Try to parse AdapterRAM (in bytes)
						if ramStr := strings.TrimSpace(parts[1]); ramStr != "" {
							if ramBytes, err := parseUint64(ramStr); err == nil && ramBytes > 0 {
								gpuInfo.MemoryTotalMB = ramBytes / 1024 / 1024
							}
						}
						
						// Try to get GPU usage from performance counter (Windows only)
						s.tryGetGPUUsageFromPerfCounter(gpuInfo)
						break
					}
				}
			}
		}
		
		// If still not found, try simple name query
		if !gpuInfo.Available {
			cmd = exec.Command("wmic", "path", "win32_VideoController", "get", "name")
			output, err = cmd.Output()
			if err == nil {
				outputStr := string(output)
				lines := strings.Split(outputStr, "\n")
				for _, line := range lines {
					trimmed := strings.TrimSpace(line)
					if trimmed != "" && !strings.Contains(strings.ToLower(trimmed), "name") {
						gpuInfo.Available = true
						gpuInfo.Name = trimmed
						s.tryGetGPUUsageFromPerfCounter(gpuInfo)
						break
					}
				}
			}
		}
	}
	
	return gpuInfo, nil
}

func (s *SystemService) tryGetGPUUsageFromPerfCounter(gpuInfo *model.GPUInfo) {
	// Try to get GPU usage from Windows Performance Counter
	// This works for Intel, AMD, and other GPUs
	cmd := exec.Command("powershell", "-Command", 
		"(Get-Counter '\\GPU Engine(*engtype_3D)\\Utilization Percentage' -ErrorAction SilentlyContinue).CounterSamples | Measure-Object -Property CookedValue -Sum | Select-Object -ExpandProperty Sum")
	output, err := cmd.Output()
	if err == nil {
		outputStr := strings.TrimSpace(string(output))
		if usage, err := parseFloat(outputStr); err == nil {
			gpuInfo.Usage = usage
		}
	}
	
	// Get comprehensive GPU memory info using PowerShell
	// This uses DXGI to get accurate GPU memory information
	psScript := `
# Try to get GPU memory from DXGI (most accurate for Windows)
try {
    Add-Type -TypeDefinition @"
using System;
using System.Runtime.InteropServices;

public class DXGI {
    [DllImport("dxgi.dll")]
    public static extern int CreateDXGIFactory1(ref Guid riid, out IntPtr ppFactory);
    
    [StructLayout(LayoutKind.Sequential)]
    public struct DXGI_ADAPTER_DESC {
        [MarshalAs(UnmanagedType.ByValTStr, SizeConst = 128)]
        public string Description;
        public uint VendorId;
        public uint DeviceId;
        public uint SubSysId;
        public uint Revision;
        public UIntPtr DedicatedVideoMemory;
        public UIntPtr DedicatedSystemMemory;
        public UIntPtr SharedSystemMemory;
    }
}
"@
} catch {}

$dedicatedMem = 0
$sharedMem = 0
$totalMem = 0

# Method 1: Try WMI VideoMemoryType query (Windows 10+)
try {
    $gpu = Get-CimInstance -ClassName Win32_VideoController -ErrorAction SilentlyContinue | Select-Object -First 1
    
    # Get AdapterRAM
    if ($gpu.AdapterRAM -and $gpu.AdapterRAM -gt 0) {
        $dedicatedMem = $gpu.AdapterRAM
    }
    
    # Try to get CurrentUsage and MaxMemorySupported
    if ($gpu.MaxMemorySupported -and $gpu.MaxMemorySupported -gt 0) {
        $totalMem = $gpu.MaxMemorySupported
    }
} catch {}

# Method 2: Query WMIC for VideoMemoryType
if ($totalMem -eq 0) {
    try {
        $wmic = wmic path Win32_VideoController get AdapterRAM,MaxMemorySupported /format:value 2>$null
        foreach ($line in $wmic) {
            if ($line -match "AdapterRAM=(\d+)") {
                $dedicatedMem = [uint64]$matches[1]
            }
            if ($line -match "MaxMemorySupported=(\d+)") {
                $totalMem = [uint64]$matches[1]
            }
        }
    } catch {}
}

# Method 3: For integrated GPUs, calculate from system memory
$systemMem = (Get-CimInstance -ClassName Win32_ComputerSystem).TotalPhysicalMemory

# Check if this is an integrated GPU (small or no dedicated memory)
$isIntegrated = ($dedicatedMem -eq 0 -or $dedicatedMem -lt 134217728)

if ($isIntegrated) {
    # For integrated GPUs, Windows typically allocates:
    # - Small dedicated memory (128MB-1GB) - this is reserved
    # - Large shared memory (up to 50% of system RAM) - this is the "total" shown in Task Manager
    
    if ($dedicatedMem -eq 0) {
        # Estimate dedicated memory for integrated GPU (usually 128MB-1GB)
        $dedicatedMem = 1073741824  # 1GB default
    }
    
    # Shared memory is typically 50% of system RAM
    # This is what Task Manager shows as "Total" for integrated GPUs
    $sharedMem = [math]::Floor($systemMem / 2)
    
    # For integrated GPUs, the "total" is just the shared memory
    # (dedicated is reserved and not counted in the user-visible total)
    $totalMem = $sharedMem
}

# If still no total, use dedicated as fallback
if ($totalMem -eq 0 -and $dedicatedMem -gt 0) {
    $totalMem = $dedicatedMem
}

Write-Output "$dedicatedMem|$sharedMem|$totalMem"
`
	
	cmd = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", psScript)
	output, err = cmd.Output()
	if err == nil {
		outputStr := strings.TrimSpace(string(output))
		parts := strings.Split(outputStr, "|")
		if len(parts) >= 3 {
			// Parse total memory (use the largest value available)
			if totalBytes, err := parseUint64(parts[2]); err == nil && totalBytes > 0 {
				gpuInfo.MemoryTotalMB = totalBytes / 1024 / 1024
			} else if dedicatedBytes, err := parseUint64(parts[0]); err == nil && dedicatedBytes > 0 {
				gpuInfo.MemoryTotalMB = dedicatedBytes / 1024 / 1024
			}
		}
	}
	
	// Try to get current GPU memory usage from Performance Counter
	// Method 1: Try Dedicated Usage counter
	cmd = exec.Command("powershell", "-Command",
		"(Get-Counter '\\GPU Adapter Memory(*)\\Dedicated Usage' -ErrorAction SilentlyContinue).CounterSamples | Measure-Object -Property CookedValue -Sum | Select-Object -ExpandProperty Sum")
	output, err = cmd.Output()
	if err == nil {
		outputStr := strings.TrimSpace(string(output))
		if memBytes, err := parseUint64(outputStr); err == nil && memBytes > 0 {
			gpuInfo.MemoryUsedMB = memBytes / 1024 / 1024
		}
	}
	
	// Method 2: If Method 1 failed, try to get from WMI CurrentUsage
	if gpuInfo.MemoryUsedMB == 0 {
		cmd = exec.Command("powershell", "-Command",
			"(Get-Counter '\\GPU Adapter Memory(*)\\Total Committed' -ErrorAction SilentlyContinue).CounterSamples | Measure-Object -Property CookedValue -Sum | Select-Object -ExpandProperty Sum")
		output, err = cmd.Output()
		if err == nil {
			outputStr := strings.TrimSpace(string(output))
			if memBytes, err := parseUint64(outputStr); err == nil && memBytes > 0 {
				gpuInfo.MemoryUsedMB = memBytes / 1024 / 1024
			}
		}
	}
	
	// Method 3: For integrated GPUs, estimate based on process usage
	if gpuInfo.MemoryUsedMB == 0 && gpuInfo.MemoryTotalMB > 0 {
		// Try to get shared memory usage
		cmd = exec.Command("powershell", "-Command",
			"(Get-Counter '\\GPU Adapter Memory(*)\\Shared Usage' -ErrorAction SilentlyContinue).CounterSamples | Measure-Object -Property CookedValue -Sum | Select-Object -ExpandProperty Sum")
		output, err = cmd.Output()
		if err == nil {
			outputStr := strings.TrimSpace(string(output))
			if memBytes, err := parseUint64(outputStr); err == nil && memBytes > 0 {
				gpuInfo.MemoryUsedMB = memBytes / 1024 / 1024
			}
		}
	}
}

func parseFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}

func parseUint64(s string) (uint64, error) {
	var u uint64
	_, err := fmt.Sscanf(s, "%d", &u)
	return u, err
}

func (s *SystemService) HasGPU() bool {
	return s.hasGPU
}

func (s *SystemService) GetSystemMetrics() (*model.SystemMetrics, error) {
	cpuUsage, err := s.GetCPUUsage()
	if err != nil {
		cpuUsage = 0
	}

	memInfo, err := s.GetMemoryUsage()
	if err != nil {
		memInfo = &model.MemoryInfo{}
	}

	gpuInfo, _ := s.GetGPUUsage()
	
	metrics := &model.SystemMetrics{
		CPU: model.CPUInfo{
			Usage: cpuUsage,
			Cores: runtime.NumCPU(),
		},
		Memory: *memInfo,
		GPU:    gpuInfo,
	}

	return metrics, nil
}
