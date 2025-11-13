package app

import (
	"sync"

	"github.com/yourusername/process-monitor-cli/internal/model"
)

type AppState struct {
	mu             sync.RWMutex
	processes      []model.Process
	systemMetrics  model.SystemMetrics
	selectedIndex  int
	aiResponse     string
	vtResponse     string
	ProcessChan    chan []model.Process
	MetricsChan    chan model.SystemMetrics
	AIResponseChan chan string
	VTResponseChan chan string
}

func NewAppState() *AppState {
	return &AppState{
		processes:      []model.Process{},
		selectedIndex:  0,
		ProcessChan:    make(chan []model.Process, 1),
		MetricsChan:    make(chan model.SystemMetrics, 1),
		AIResponseChan: make(chan string, 1),
		VTResponseChan: make(chan string, 1),
	}
}

func (a *AppState) UpdateProcesses(processes []model.Process) {
	a.mu.Lock()
	a.processes = processes
	a.mu.Unlock()

	select {
	case a.ProcessChan <- processes:
	default:
	}
}

func (a *AppState) UpdateSystemMetrics(metrics model.SystemMetrics) {
	a.mu.Lock()
	a.systemMetrics = metrics
	a.mu.Unlock()

	select {
	case a.MetricsChan <- metrics:
	default:
	}
}

func (a *AppState) SetSelectedProcess(index int) {
	a.mu.Lock()
	a.selectedIndex = index
	a.mu.Unlock()
}

func (a *AppState) SetAIResponse(response string) {
	a.mu.Lock()
	a.aiResponse = response
	a.mu.Unlock()

	select {
	case a.AIResponseChan <- response:
	default:
	}
}

func (a *AppState) SetVTResponse(response string) {
	a.mu.Lock()
	a.vtResponse = response
	a.mu.Unlock()

	select {
	case a.VTResponseChan <- response:
	default:
	}
}

func (a *AppState) GetProcesses() []model.Process {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.processes
}

func (a *AppState) GetSelectedProcess() *model.Process {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if a.selectedIndex >= 0 && a.selectedIndex < len(a.processes) {
		return &a.processes[a.selectedIndex]
	}
	return nil
}

func (a *AppState) GetSelectedIndex() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.selectedIndex
}
