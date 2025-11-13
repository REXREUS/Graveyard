package service

import (
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/yourusername/process-monitor-cli/internal/model"
)

type ProcessService struct {
	cache     map[int32]*process.Process
	cpuCache  map[int32]float64
	lastCheck time.Time
}

func NewProcessService() *ProcessService {
	return &ProcessService{
		cache:     make(map[int32]*process.Process),
		cpuCache:  make(map[int32]float64),
		lastCheck: time.Now(),
	}
}

func (p *ProcessService) GetProcessList() ([]model.Process, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	// Clean up cache periodically (every 10 seconds)
	if time.Since(p.lastCheck) > 10*time.Second {
		p.cleanupCache(procs)
		p.lastCheck = time.Now()
	}

	var processes []model.Process
	for _, proc := range procs {
		procInfo, err := p.getProcessInfo(proc)
		if err != nil {
			continue
		}
		processes = append(processes, *procInfo)
	}

	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPUPercent > processes[j].CPUPercent
	})

	if len(processes) > 100 {
		processes = processes[:100]
	}

	return processes, nil
}

func (p *ProcessService) cleanupCache(currentProcs []*process.Process) {
	// Create map of current PIDs
	currentPIDs := make(map[int32]bool)
	for _, proc := range currentProcs {
		currentPIDs[proc.Pid] = true
	}
	
	// Remove cached processes that no longer exist
	for pid := range p.cache {
		if !currentPIDs[pid] {
			delete(p.cache, pid)
			delete(p.cpuCache, pid)
		}
	}
}

func (p *ProcessService) getProcessInfo(proc *process.Process) (*model.Process, error) {
	name, _ := proc.Name()
	
	// Get CPU percent with interval for accurate reading
	// Use cached value if available, otherwise get new reading
	var cpuPercent float64
	pid := proc.Pid
	
	// Check if we have a cached process
	if cachedProc, exists := p.cache[pid]; exists {
		// Get CPU percent from cached process (this will show actual usage)
		cpuPercent, _ = cachedProc.CPUPercent()
	} else {
		// First time seeing this process, initialize
		cpuPercent, _ = proc.CPUPercent()
	}
	
	// Update cache
	p.cache[pid] = proc
	p.cpuCache[pid] = cpuPercent
	
	memInfo, _ := proc.MemoryInfo()
	memPercent, _ := proc.MemoryPercent()
	username, _ := proc.Username()
	cmdline, _ := proc.Cmdline()

	var memMB uint64
	if memInfo != nil {
		memMB = memInfo.RSS / 1024 / 1024
	}

	return &model.Process{
		PID:           pid,
		Name:          name,
		CPUPercent:    cpuPercent,
		MemoryMB:      memMB,
		MemoryPercent: memPercent,
		Username:      username,
		Command:       cmdline,
	}, nil
}

func (p *ProcessService) KillProcess(pid int32) error {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return err
	}
	return proc.Kill()
}
