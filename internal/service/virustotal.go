package service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/yourusername/process-monitor-cli/internal/model"
)

type VirusTotalService struct {
	apiKey     string
	httpClient *http.Client
}

type VTFileReport struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats struct {
				Malicious int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Undetected int `json:"undetected"`
				Harmless   int `json:"harmless"`
			} `json:"last_analysis_stats"`
			LastAnalysisResults map[string]struct {
				Category string `json:"category"`
				Result   string `json:"result"`
			} `json:"last_analysis_results"`
			Names []string `json:"names"`
		} `json:"attributes"`
	} `json:"data"`
}

func NewVirusTotalService(apiKey string) *VirusTotalService {
	return &VirusTotalService{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (v *VirusTotalService) ScanProcess(ctx context.Context, proc model.Process) (*model.VTScanResult, error) {
	// Get process executable path
	exePath, err := v.getProcessExecutablePath(proc.PID)
	if err != nil {
		return nil, fmt.Errorf("failed to get executable path: %w", err)
	}

	// Calculate file hash
	fileHash, err := v.calculateFileHash(exePath)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate hash: %w", err)
	}

	// Query VirusTotal API
	report, err := v.getFileReport(ctx, fileHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get VT report: %w", err)
	}

	// Build scans map
	scans := make(map[string]model.ScanResult)
	detectionCount := 0
	for engine, engineResult := range report.Data.Attributes.LastAnalysisResults {
		detected := engineResult.Category == "malicious" || engineResult.Category == "suspicious"
		scans[engine] = model.ScanResult{
			Detected: detected,
			Result:   engineResult.Result,
		}
		if detected {
			detectionCount++
		}
	}

	totalEngines := report.Data.Attributes.LastAnalysisStats.Malicious +
		report.Data.Attributes.LastAnalysisStats.Suspicious +
		report.Data.Attributes.LastAnalysisStats.Undetected +
		report.Data.Attributes.LastAnalysisStats.Harmless

	scanResult := &model.VTScanResult{
		ProcessName:  proc.Name,
		PID:          proc.PID,
		FilePath:     exePath,
		FileHash:     fileHash,
		Malicious:    report.Data.Attributes.LastAnalysisStats.Malicious,
		Suspicious:   report.Data.Attributes.LastAnalysisStats.Suspicious,
		Undetected:   report.Data.Attributes.LastAnalysisStats.Undetected,
		Harmless:     report.Data.Attributes.LastAnalysisStats.Harmless,
		ScanDate:     time.Now(),
		TotalEngines: totalEngines,
		Scans:        scans,
	}

	// Get top detections
	for engine, engineResult := range report.Data.Attributes.LastAnalysisResults {
		if engineResult.Category == "malicious" || engineResult.Category == "suspicious" {
			detection := fmt.Sprintf("%s: %s", engine, engineResult.Result)
			scanResult.Detections = append(scanResult.Detections, detection)
			if len(scanResult.Detections) >= 5 {
				break
			}
		}
	}

	return scanResult, nil
}

func (v *VirusTotalService) getProcessExecutablePath(pid int32) (string, error) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return "", err
	}
	
	exePath, err := proc.Exe()
	if err != nil {
		return "", err
	}
	
	return exePath, nil
}

func (v *VirusTotalService) calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func (v *VirusTotalService) getFileReport(ctx context.Context, fileHash string) (*VTFileReport, error) {
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s", fileHash)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-apikey", v.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("file not found in VirusTotal database")
	}

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("VT API error (status %d): %s", resp.StatusCode, string(body))
	}

	var report VTFileReport
	if err := json.NewDecoder(resp.Body).Decode(&report); err != nil {
		return nil, err
	}

	return &report, nil
}
