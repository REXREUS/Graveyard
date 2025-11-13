package service

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/yourusername/process-monitor-cli/internal/model"
	"google.golang.org/api/option"
)

type AIService struct {
	client *genai.Client
	model  *genai.GenerativeModel
	apiKey string
}

func NewAIService(apiKey string) (*AIService, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	// Try gemini-1.5-flash-latest or gemini-pro
	model := client.GenerativeModel("gemini-2.5-flash")

	return &AIService{
		client: client,
		model:  model,
		apiKey: apiKey,
	}, nil
}

func (a *AIService) AnalyzeProcess(ctx context.Context, proc model.Process) (string, error) {
	prompt := a.buildPrompt(proc)

	resp, err := a.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "No response from AI", nil
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}

func (a *AIService) buildPrompt(proc model.Process) string {
	return fmt.Sprintf(`You are a system process analyzer. Provide a concise analysis of the following process:

Process Name: %s
PID: %d
CPU Usage: %.1f%%
Memory Usage: %dMB

Please provide:
1. Purpose: What is this process and what does it do?
2. Safety: Is this process safe or potentially harmful?
3. Recommendation: Can this process be safely terminated?

Keep the response clear, professional, and under 200 words.`, proc.Name, proc.PID, proc.CPUPercent, proc.MemoryMB)
}

func (a *AIService) AnalyzeVirusTotalResult(ctx context.Context, proc model.Process, vtResult *model.VTScanResult) (string, error) {
	prompt := a.buildVTAnalysisPrompt(proc, vtResult)

	resp, err := a.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "No response from AI", nil
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}

func (a *AIService) buildVTAnalysisPrompt(proc model.Process, vtResult *model.VTScanResult) string {
	detectionsStr := "None"
	if len(vtResult.Detections) > 0 {
		detectionsStr = ""
		for i, d := range vtResult.Detections {
			detectionsStr += fmt.Sprintf("%d. %s\n", i+1, d)
		}
	}

	return fmt.Sprintf(`You are a cybersecurity analyst. Analyze the following VirusTotal scan result for a running process:

Process Information:
- Name: %s
- PID: %d
- CPU Usage: %.1f%%
- Memory Usage: %dMB
- File Path: %s
- File Hash (SHA256): %s

VirusTotal Scan Results:
- Threat Level: %s
- Malicious Detections: %d
- Suspicious Detections: %d
- Harmless: %d
- Undetected: %d
- Total Engines: %d

Top Detections:
%s

Please provide:
1. Risk Assessment: Evaluate the overall risk level based on the scan results
2. Analysis: Explain what the detections mean and whether they are false positives
3. Recommendation: Should this process be terminated? What actions should be taken?
4. Additional Context: Any relevant information about this process type

Keep the response clear, professional, and actionable. Use bullet points for clarity.`,
		proc.Name, proc.PID, proc.CPUPercent, proc.MemoryMB,
		vtResult.FilePath, vtResult.FileHash,
		vtResult.GetThreatLevel(),
		vtResult.Malicious, vtResult.Suspicious, vtResult.Harmless, vtResult.Undetected,
		vtResult.Malicious+vtResult.Suspicious+vtResult.Harmless+vtResult.Undetected,
		detectionsStr)
}

func (a *AIService) Close() error {
	if a.client != nil {
		return a.client.Close()
	}
	return nil
}
