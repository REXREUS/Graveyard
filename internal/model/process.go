package model

import "fmt"

type Process struct {
	PID           int32
	Name          string
	CPUPercent    float64
	MemoryMB      uint64
	MemoryPercent float32
	Username      string
	Command       string
}

func (p Process) String() string {
	return fmt.Sprintf("%-8d %-20s %.1f%%  %dMB",
		p.PID, p.Name, p.CPUPercent, p.MemoryMB)
}

