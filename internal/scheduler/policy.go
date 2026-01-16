package scheduler

import "adaptive-task-scheduler/internal/worker"

// func ComputeScore(m worker.Metrics) float64 {
// 	return 0.5*m.CPUUsage + 0.3*m.MemoryUsage + 0.2*m.Latency
// }

// Updated*
func ComputeScore(m worker.Metrics) float64 {
	return wCPU*m.CPUUsage +
		wMemory*m.MemoryUsage +
		wLatency*m.Latency
}

// Lower score = healthier worker
var (
	wCPU     = 0.5
	wMemory  = 0.3
	wLatency = 0.2
)

// Feedback-based adjustment
func AdjustWeights(avgLatency float64) {
	if avgLatency > 0.8 {
		wLatency += 0.05
		wCPU -= 0.03
		wMemory -= 0.02
	}
}
