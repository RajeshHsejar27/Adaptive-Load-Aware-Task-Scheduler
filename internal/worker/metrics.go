package worker

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Metrics struct {
	CPUUsage    float64
	MemoryUsage float64
	Latency     float64
}

func CollectMetrics() Metrics {
	cpuPercents, _ := cpu.Percent(200*time.Millisecond, false)
	memStat, _ := mem.VirtualMemory()

	return Metrics{
		CPUUsage:    cpuPercents[0],      // %
		MemoryUsage: memStat.UsedPercent, // %
	}
}
