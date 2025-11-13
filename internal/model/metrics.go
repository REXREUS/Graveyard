package model

import (
	"fmt"
	"time"
)

type SystemMetrics struct {
	CPU       CPUInfo
	Memory    MemoryInfo
	GPU       *GPUInfo
	Timestamp time.Time
}

type CPUInfo struct {
	Usage float64
	Cores int
}

type MemoryInfo struct {
	UsedGB     float64
	TotalGB    float64
	Percentage float64
}

type GPUInfo struct {
	Available     bool
	Name          string
	Usage         float64
	MemoryUsedMB  uint64
	MemoryTotalMB uint64
}

type VTScanResult struct {
	ProcessName  string
	PID          int32
	FilePath     string
	FileHash     string
	Malicious    int
	Suspicious   int
	Undetected   int
	Harmless     int
	Detections   []string
	ScanDate     time.Time
	TotalEngines int
	Scans        map[string]ScanResult
}

// ScanResult represents individual engine scan result
type ScanResult struct {
	Detected bool
	Result   string
}

// GetThreatLevel returns threat level based on detections
func (v *VTScanResult) GetThreatLevel() string {
	total := v.Malicious + v.Suspicious + v.Undetected + v.Harmless
	if total == 0 {
		return "UNKNOWN"
	}

	maliciousPercent := float64(v.Malicious) / float64(total) * 100

	if maliciousPercent >= 10 {
		return "HIGH"
	} else if maliciousPercent > 0 || v.Suspicious > 0 {
		return "MEDIUM"
	}
	return "SAFE"
}

// GetSummary returns a summary of the scan
func (v *VTScanResult) GetSummary() string {
	total := v.Malicious + v.Suspicious + v.Undetected + v.Harmless
	if total == 0 {
		return "No scan data available"
	}

	if v.Malicious == 0 && v.Suspicious == 0 {
		return fmt.Sprintf("Clean - No threats detected by %d engines", total)
	}

	if v.Malicious > 0 {
		return fmt.Sprintf("Malicious - %d/%d engines detected threats", v.Malicious, total)
	}

	return fmt.Sprintf("Suspicious - %d/%d engines flagged as suspicious", v.Suspicious, total)
}
