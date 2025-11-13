package ui

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/yourusername/process-monitor-cli/internal/app"
	"github.com/yourusername/process-monitor-cli/internal/logger"
	"github.com/yourusername/process-monitor-cli/internal/model"
	"github.com/yourusername/process-monitor-cli/internal/service"
)

type UIManager struct {
	app            *tview.Application
	root           tview.Primitive
	processList    *tview.Table
	cpuGauge       *tview.TextView
	memoryGauge    *tview.TextView
	gpuGauge       *tview.TextView
	infoPanel      *tview.TextView
	vtPanel        *tview.TextView
	selectedIndex  int
	state          *app.AppState
	processService *service.ProcessService
	systemService  *service.SystemService
	aiService      *service.AIService
	configService  *service.ConfigService
	vtService      *service.VirusTotalService
	focusedPanel   string // "process", "ai", or "vt"
}

func NewUIManager(state *app.AppState, ps *service.ProcessService, ss *service.SystemService, as *service.AIService, cs *service.ConfigService, vts *service.VirusTotalService) *UIManager {
	return &UIManager{
		app:            tview.NewApplication(),
		state:          state,
		processService: ps,
		systemService:  ss,
		aiService:      as,
		configService:  cs,
		vtService:      vts,
		focusedPanel:   "process", // Default focus on process list
	}
}

func (u *UIManager) InitializeLayout() {
	u.processList = tview.NewTable().
		SetBorders(false).
		SetSelectable(true, false)

	u.cpuGauge = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)
	u.cpuGauge.SetBorder(true).SetTitle("CPU")

	u.memoryGauge = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)
	u.memoryGauge.SetBorder(true).SetTitle("Memory")

	u.gpuGauge = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)
	u.gpuGauge.SetBorder(true).SetTitle("GPU")

	u.infoPanel = tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true).
		SetScrollable(true).
		SetWrap(true)
	u.infoPanel.SetBorder(true).SetTitle("AI Assistant (Tab to focus, ↑↓ to scroll)")
	u.infoPanel.SetText("[cyan]Select a process and press 'i' to inspect\n\n[yellow]Note:[white] AI features require a valid Gemini API key.\nPress 's' to configure in Settings.")
	
	// Enable input capture for scrolling
	u.infoPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event // Allow default scrolling behavior
	})

	u.vtPanel = tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true).
		SetScrollable(true).
		SetWrap(true)
	u.vtPanel.SetBorder(true).SetTitle("VirusTotal Scan Results (Tab to focus, ↑↓ to scroll)")
	u.vtPanel.SetText("[cyan]Press 't' to scan selected process\n\n[yellow]Status:[white] Ready\n\n[dim]No scan performed yet[-]")
	
	// Enable input capture for scrolling
	u.vtPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event // Allow default scrolling behavior
	})

	processBox := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(u.processList, 0, 1, true)
	processBox.SetBorder(true).SetTitle("Process List")

	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(u.cpuGauge, 4, 0, false).
		AddItem(u.memoryGauge, 4, 0, false).
		AddItem(u.gpuGauge, 7, 0, false).
		AddItem(u.infoPanel, 0, 1, false)

	mainLayout := tview.NewFlex().
		AddItem(leftPanel, 50, 0, false).
		AddItem(u.vtPanel, 45, 0, false).
		AddItem(processBox, 0, 1, true)

	header := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Graveyard v1.0.0[-]     [s] Settings  [q] Quit")
	header.SetTextAlign(tview.AlignCenter)

	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Tab[white]: Switch Panel  [yellow]↑↓[white]: Navigate/Scroll  [yellow]i[white]: Inspect  [yellow]t[white]: VT Scan  [yellow]k[white]: Kill  [yellow]s[white]: Settings  [yellow]q[white]: Quit")
	footer.SetTextAlign(tview.AlignCenter)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 1, 0, false).
		AddItem(mainLayout, 0, 1, true).
		AddItem(footer, 1, 0, false)

	u.setupKeyBindings()
	u.root = layout
	u.app.SetRoot(layout, true)
}

func (u *UIManager) setupKeyBindings() {
	// Global key handler
	u.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			// Cycle through panels: process -> ai -> vt -> process
			u.cycleFocus()
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				u.app.Stop()
				return nil
			case 'i':
				u.inspectProcess()
				return nil
			case 't', 'T':
				u.scanWithVirusTotal()
				return nil
			case 'k':
				u.killProcess()
				return nil
			case 's':
				u.ShowSettingsDialog()
				return nil
			}
		case tcell.KeyEscape:
			// Return focus to process list
			if u.focusedPanel != "process" {
				u.setFocus("process")
				return nil
			}
			u.app.Stop()
			return nil
		}
		return event
	})

	u.processList.SetSelectionChangedFunc(func(row, col int) {
		if row > 0 {
			u.selectedIndex = row - 1
			u.state.SetSelectedProcess(row - 1)
		}
	})
}

func (u *UIManager) cycleFocus() {
	switch u.focusedPanel {
	case "process":
		u.setFocus("ai")
	case "ai":
		u.setFocus("vt")
	case "vt":
		u.setFocus("process")
	}
}

func (u *UIManager) setFocus(panel string) {
	u.focusedPanel = panel
	
	// Update border colors to show focus
	u.processList.SetBorderColor(tcell.ColorWhite)
	u.infoPanel.SetBorderColor(tcell.ColorWhite)
	u.vtPanel.SetBorderColor(tcell.ColorWhite)
	
	switch panel {
	case "process":
		u.processList.SetBorderColor(tcell.ColorYellow)
		u.app.SetFocus(u.processList)
	case "ai":
		u.infoPanel.SetBorderColor(tcell.ColorYellow)
		u.app.SetFocus(u.infoPanel)
	case "vt":
		u.vtPanel.SetBorderColor(tcell.ColorYellow)
		u.app.SetFocus(u.vtPanel)
	}
}

func (u *UIManager) RenderProcessList(processes []model.Process) {
	u.processList.Clear()

	headers := []string{"PID", "Name", "CPU%", "RAM%", "RAM"}
	
	for i, header := range headers {
		cell := tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignLeft).
			SetSelectable(false)
		u.processList.SetCell(0, i, cell)
	}

	for i, proc := range processes {
		// Truncate process name if too long (max 15 chars)
		name := proc.Name
		if len(name) > 15 {
			name = name[:12] + "..."
		}
		
		// Format PID (max 6 chars)
		pidStr := fmt.Sprintf("%5d", proc.PID)
		
		// Format CPU% (max 5 chars)
		var cpuStr string
		if proc.CPUPercent >= 10 {
			cpuStr = fmt.Sprintf("%3.0f%%", proc.CPUPercent)
		} else if proc.CPUPercent >= 1 {
			cpuStr = fmt.Sprintf("%3.1f%%", proc.CPUPercent)
		} else if proc.CPUPercent > 0 {
			cpuStr = fmt.Sprintf("%.1f%%", proc.CPUPercent)
		} else {
			cpuStr = "  0%%"
		}
		
		// Format RAM% (max 4 chars)
		var ramPercentStr string
		if proc.MemoryPercent >= 10 {
			ramPercentStr = fmt.Sprintf("%2.0f%%", proc.MemoryPercent)
		} else if proc.MemoryPercent >= 1 {
			ramPercentStr = fmt.Sprintf("%.1f%%", proc.MemoryPercent)
		} else {
			ramPercentStr = fmt.Sprintf("%.1f%%", proc.MemoryPercent)
		}
		
		// Format memory size (max 5 chars)
		memStr := fmt.Sprintf("%4dM", proc.MemoryMB)
		
		u.processList.SetCell(i+1, 0, tview.NewTableCell(pidStr))
		u.processList.SetCell(i+1, 1, tview.NewTableCell(name))
		u.processList.SetCell(i+1, 2, tview.NewTableCell(cpuStr))
		u.processList.SetCell(i+1, 3, tview.NewTableCell(ramPercentStr))
		u.processList.SetCell(i+1, 4, tview.NewTableCell(memStr))
	}
}

func (u *UIManager) RenderSystemMetrics(metrics model.SystemMetrics) {
	// CPU with percentage display
	cpuBar := u.createProgressBar(metrics.CPU.Usage, 20)
	u.cpuGauge.SetText(fmt.Sprintf("Usage: [cyan]%.1f%%[white] (%d cores)\n%s", 
		metrics.CPU.Usage, metrics.CPU.Cores, cpuBar))

	// Memory with percentage
	memBar := u.createProgressBar(metrics.Memory.Percentage, 20)
	u.memoryGauge.SetText(fmt.Sprintf("%.1fGB / %.1fGB ([cyan]%.1f%%[white])\n%s",
		metrics.Memory.UsedGB, metrics.Memory.TotalGB, metrics.Memory.Percentage, memBar))

	// GPU with metrics
	if metrics.GPU != nil && metrics.GPU.Available {
		gpuName := metrics.GPU.Name
		if len(gpuName) > 25 {
			gpuName = gpuName[:22] + "..."
		}
		
		if metrics.GPU.MemoryTotalMB > 0 {
			gpuMemPercent := float64(metrics.GPU.MemoryUsedMB) / float64(metrics.GPU.MemoryTotalMB) * 100
			gpuBar := u.createProgressBar(gpuMemPercent, 20)
			gpuUsageBar := u.createProgressBar(metrics.GPU.Usage, 20)
			
			// Format memory size intelligently (MB or GB)
			var memUsedStr, memTotalStr string
			if metrics.GPU.MemoryTotalMB >= 1024 {
				// Use GB for large memory (>= 1GB)
				memUsedGB := float64(metrics.GPU.MemoryUsedMB) / 1024.0
				memTotalGB := float64(metrics.GPU.MemoryTotalMB) / 1024.0
				memUsedStr = fmt.Sprintf("%.1fGB", memUsedGB)
				memTotalStr = fmt.Sprintf("%.1fGB", memTotalGB)
			} else {
				// Use MB for small memory (< 1GB)
				memUsedStr = fmt.Sprintf("%dMB", metrics.GPU.MemoryUsedMB)
				memTotalStr = fmt.Sprintf("%dMB", metrics.GPU.MemoryTotalMB)
			}
			
			u.gpuGauge.SetText(fmt.Sprintf("[cyan]%s[white]\nUsage: [cyan]%.1f%%[white]\n%s\nMem: %s/%s ([cyan]%.1f%%[white])\n%s",
				gpuName, metrics.GPU.Usage, gpuUsageBar, 
				memUsedStr, memTotalStr, gpuMemPercent, gpuBar))
		} else {
			u.gpuGauge.SetText(fmt.Sprintf("[cyan]%s[white]\n[dim]Metrics not available[-]", gpuName))
		}
	} else {
		u.gpuGauge.SetText("[dim]Not detected[-]")
	}
}

func (u *UIManager) createProgressBar(percent float64, width int) string {
	filled := int(percent / 100 * float64(width))
	if filled > width {
		filled = width
	}

	var color string
	if percent < 50 {
		color = "green"
	} else if percent < 80 {
		color = "yellow"
	} else {
		color = "red"
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return fmt.Sprintf("[%s]%s[-]", color, bar)
}

func (u *UIManager) RenderInfoPanel(content string) {
	u.infoPanel.SetText(content)
}

func (u *UIManager) RenderVTPanel(content string) {
	u.vtPanel.SetText(content)
}

func (u *UIManager) inspectProcess() {
	proc := u.state.GetSelectedProcess()
	if proc == nil {
		return
	}

	if u.aiService == nil {
		u.RenderInfoPanel("[red]AI Service not available. Please configure API key in Settings (press 's')[-]")
		return
	}

	u.RenderInfoPanel("[yellow]Analyzing process...[-]")

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		response, err := u.aiService.AnalyzeProcess(ctx, *proc)
		if err != nil {
			logger.Error("AI analysis failed:", err)
			u.state.SetAIResponse(fmt.Sprintf("[red]AI Analysis Failed[-]\n\n[yellow]Possible reasons:[-]\n• Invalid or expired API key\n• Network connection issue\n• API service unavailable\n\n[cyan]Please check your API key in Settings (press 's')[-]"))
		} else {
			u.state.SetAIResponse(response)
		}
	}()
}

func (u *UIManager) scanWithVirusTotal() {
	proc := u.state.GetSelectedProcess()
	if proc == nil {
		return
	}

	if u.vtService == nil {
		u.RenderInfoPanel("[red]VirusTotal Service not available. Please configure API key in Settings (press 's')[-]")
		return
	}

	if u.aiService == nil {
		u.RenderInfoPanel("[red]AI Service not available. Please configure Gemini API key in Settings (press 's')[-]")
		return
	}

	// Update VT Panel via state
	u.state.SetVTResponse(fmt.Sprintf("[yellow]⟳ Scanning Process...[-]\n\n"+
		"[cyan]Process:[-] %s\n"+
		"[cyan]PID:[-] %d\n\n"+
		"[dim]Connecting to VirusTotal...\n"+
		"Please wait...[-]", proc.Name, proc.PID))

	// Update AI Panel
	u.state.SetAIResponse("[yellow]Waiting for VirusTotal scan...[-]")

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		// Step 1: Scan with VirusTotal
		vtResult, err := u.vtService.ScanProcess(ctx, *proc)
		if err != nil {
			logger.Error("VirusTotal scan failed:", err)
			u.state.SetVTResponse(fmt.Sprintf("[red]✗ Scan Failed[-]\n\n"+
				"[yellow]Process:[-] %s (PID: %d)\n\n"+
				"[red]Error:[-] %v\n\n"+
				"[yellow]Possible reasons:[-]\n"+
				"• Invalid or expired API key\n"+
				"• File not found in VT database\n"+
				"• Network connection issue\n"+
				"• Rate limit exceeded\n\n"+
				"[cyan]Press 's' to configure API key[-]", proc.Name, proc.PID, err))
			return
		}

		// Show VT results in dedicated panel
		vtDisplay := u.formatVTResults(vtResult)
		u.state.SetVTResponse(vtDisplay)

		// Update AI panel status
		u.state.SetAIResponse("[yellow]Analyzing with AI...\n\nPlease wait...[-]")

		// Step 2: Analyze with Gemini AI
		if u.aiService != nil {
			aiResponse, err := u.aiService.AnalyzeVirusTotalResult(ctx, *proc, vtResult)
			if err != nil {
				logger.Error("AI analysis failed:", err)
				u.state.SetAIResponse("[red]AI Analysis Failed[-]\n\n" +
					"Please review the VirusTotal results in the center panel.\n\n" +
					"[cyan]Press 's' to configure Gemini API key[-]")
				return
			}

			// Show AI analysis
			u.state.SetAIResponse(fmt.Sprintf("[cyan]AI Security Analysis[-]\n"+
				"[cyan]═══════════════════════════════════[-]\n\n"+
				"%s", aiResponse))
		} else {
			u.state.SetAIResponse("[yellow]AI Service not available[-]\n\n" +
				"Configure Gemini API key in Settings (press 's') for AI analysis.")
		}
	}()
}

func (u *UIManager) getThreatColor(level string) string {
	switch level {
	case "HIGH":
		return "red"
	case "MEDIUM":
		return "yellow"
	case "SAFE":
		return "green"
	default:
		return "white"
	}
}

func (u *UIManager) formatVTResults(vtResult *model.VTScanResult) string {
	threatLevel := vtResult.GetThreatLevel()
	threatColor := u.getThreatColor(threatLevel)
	
	// Threat icon
	var threatIcon string
	switch threatLevel {
	case "HIGH":
		threatIcon = "⚠"
	case "MEDIUM":
		threatIcon = "⚡"
	case "SAFE":
		threatIcon = "✓"
	default:
		threatIcon = "?"
	}

	// Detection ratio bar
	detectionCount := vtResult.Malicious + vtResult.Suspicious
	detectionPercent := 0.0
	if vtResult.TotalEngines > 0 {
		detectionPercent = float64(detectionCount) / float64(vtResult.TotalEngines) * 100
	}
	detectionBar := u.createProgressBar(detectionPercent, 20)

	// Format file path (truncate if too long)
	filePath := vtResult.FilePath
	if len(filePath) > 35 {
		filePath = "..." + filePath[len(filePath)-32:]
	}

	// Format hash (show first 16 chars)
	hashShort := vtResult.FileHash
	if len(hashShort) > 16 {
		hashShort = hashShort[:16] + "..."
	}

	// Build the display
	display := fmt.Sprintf("[%s]%s %s[-]\n\n", threatColor, threatIcon, threatLevel)
	display += fmt.Sprintf("[cyan]Process Information[-]\n")
	display += fmt.Sprintf("Name: [white]%s[-]\n", vtResult.ProcessName)
	display += fmt.Sprintf("PID:  [white]%d[-]\n\n", vtResult.PID)
	
	display += fmt.Sprintf("[cyan]File Details[-]\n")
	display += fmt.Sprintf("Path: [white]%s[-]\n", filePath)
	display += fmt.Sprintf("Hash: [white]%s[-]\n\n", hashShort)
	
	display += fmt.Sprintf("[cyan]Detection Results[-]\n")
	display += fmt.Sprintf("Engines: [white]%d / %d[-] detected\n", detectionCount, vtResult.TotalEngines)
	display += fmt.Sprintf("%s\n\n", detectionBar)
	
	display += fmt.Sprintf("[cyan]Summary[-]\n")
	display += fmt.Sprintf("[white]%s[-]\n\n", vtResult.GetSummary())

	// Show top detections if any
	if detectionCount > 0 && len(vtResult.Scans) > 0 {
		display += fmt.Sprintf("[cyan]Top Detections[-]\n")
		count := 0
		for engine, result := range vtResult.Scans {
			if result.Detected && count < 5 {
				malwareName := result.Result
				if len(malwareName) > 25 {
					malwareName = malwareName[:22] + "..."
				}
				display += fmt.Sprintf("• [yellow]%s[-]: %s\n", engine, malwareName)
				count++
			}
		}
		if detectionCount > 5 {
			display += fmt.Sprintf("\n[dim]...and %d more[-]", detectionCount-5)
		}
	}

	return display
}

func (u *UIManager) killProcess() {
	proc := u.state.GetSelectedProcess()
	if proc == nil {
		return
	}

	modal := tview.NewModal().
		SetText(fmt.Sprintf("Kill process?\n\nName: %s\nPID: %d", proc.Name, proc.PID)).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				err := u.processService.KillProcess(proc.PID)
				if err != nil {
					logger.Error("Failed to kill process:", err)
					u.ShowMessage(fmt.Sprintf("Failed to kill process: %v", err), "error")
				} else {
					u.ShowMessage(fmt.Sprintf("Process %d killed successfully", proc.PID), "success")
				}
			} else {
				u.setupKeyBindings() // Re-enable key bindings
				u.app.SetRoot(u.root, true)
			}
		})

	// Disable global input capture while modal is open
	u.app.SetInputCapture(nil)
	
	u.app.SetRoot(modal, true)
	u.app.SetFocus(modal)
}

func (u *UIManager) ShowSettingsDialog() {
	form := tview.NewForm()
	
	// Gemini API Key
	currentKey := u.configService.GetAPIKey()
	if currentKey != "" && len(currentKey) > 10 {
		currentKey = currentKey[:10] + "..."
	}
	
	// VirusTotal API Key
	currentVTKey := u.configService.GetVTAPIKey()
	if currentVTKey != "" && len(currentVTKey) > 10 {
		currentVTKey = currentVTKey[:10] + "..."
	}

	form.AddInputField("Gemini API Key", currentKey, 50, nil, nil)
	form.AddInputField("VirusTotal API Key", currentVTKey, 50, nil, nil)
	
	form.AddButton("Save", func() {
		apiKey := form.GetFormItem(0).(*tview.InputField).GetText()
		vtKey := form.GetFormItem(1).(*tview.InputField).GetText()
		
		saved := false
		
		// Save Gemini API Key
		if u.configService.ValidateAPIKey(apiKey) && !strings.HasSuffix(apiKey, "...") {
			err := u.configService.SaveAPIKey(apiKey)
			if err != nil {
				logger.Error("Failed to save Gemini API key:", err)
				u.ShowMessage("Failed to save Gemini API key", "error")
				u.setupKeyBindings() // Re-enable key bindings
				u.app.SetRoot(u.root, true)
				return
			}
			newAI, err := service.NewAIService(apiKey)
			if err == nil {
				if u.aiService != nil {
					u.aiService.Close()
				}
				u.aiService = newAI
			}
			saved = true
		}
		
		// Save VirusTotal API Key
		if u.configService.ValidateAPIKey(vtKey) && !strings.HasSuffix(vtKey, "...") {
			err := u.configService.SaveVTAPIKey(vtKey)
			if err != nil {
				logger.Error("Failed to save VirusTotal API key:", err)
				u.ShowMessage("Failed to save VirusTotal API key", "error")
				u.setupKeyBindings() // Re-enable key bindings
				u.app.SetRoot(u.root, true)
				return
			}
			u.vtService = service.NewVirusTotalService(vtKey)
			saved = true
		}
		
		if saved {
			u.ShowMessage("API keys saved successfully", "success")
		}
		u.setupKeyBindings() // Re-enable key bindings
		u.app.SetRoot(u.root, true)
	})
	
	form.AddButton("Delete All", func() {
		u.configService.DeleteAPIKey()
		u.ShowMessage("All API keys deleted", "success")
		u.setupKeyBindings() // Re-enable key bindings
		u.app.SetRoot(u.root, true)
	})
	
	form.AddButton("Cancel", func() {
		u.setupKeyBindings() // Re-enable key bindings
		u.app.SetRoot(u.root, true)
	})

	form.SetBorder(true).SetTitle("Settings - API Configuration").SetTitleAlign(tview.AlignCenter)
	
	// Disable global input capture while form is open
	u.app.SetInputCapture(nil)
	
	u.app.SetRoot(form, true)
	u.app.SetFocus(form)
}

func (u *UIManager) ShowMessage(message, msgType string) {
	color := "white"
	if msgType == "error" {
		color = "red"
	} else if msgType == "success" {
		color = "green"
	}

	modal := tview.NewModal().
		SetText(fmt.Sprintf("[%s]%s[-]", color, message)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			u.setupKeyBindings() // Re-enable key bindings
			u.app.SetRoot(u.root, true)
		})

	// Disable global input capture while modal is open
	u.app.SetInputCapture(nil)
	
	u.app.SetRoot(modal, true)
	u.app.SetFocus(modal)
}

func (u *UIManager) Run() error {
	return u.app.Run()
}

func (u *UIManager) Stop() {
	u.app.Stop()
}

func (u *UIManager) StartUpdateLoop(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case processes := <-u.state.ProcessChan:
				u.app.QueueUpdateDraw(func() {
					u.RenderProcessList(processes)
				})
			case metrics := <-u.state.MetricsChan:
				u.app.QueueUpdateDraw(func() {
					u.RenderSystemMetrics(metrics)
				})
			case response := <-u.state.AIResponseChan:
				u.app.QueueUpdateDraw(func() {
					u.RenderInfoPanel(response)
				})
			case vtResponse := <-u.state.VTResponseChan:
				u.app.QueueUpdateDraw(func() {
					u.RenderVTPanel(vtResponse)
				})
			}
		}
	}()
}
