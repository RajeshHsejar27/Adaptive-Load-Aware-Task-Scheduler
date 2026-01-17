package scheduler

import "adaptive-task-scheduler/internal/worker"

type LeastLoadedScheduler struct {
	workers []*worker.Worker
}

func NewLeastLoadedScheduler(workers []*worker.Worker) *LeastLoadedScheduler {
	return &LeastLoadedScheduler{
		workers: workers,
	}
}

func (l *LeastLoadedScheduler) Schedule(task func()) {
	var chosen *worker.Worker
	minCPU := 100.0

	for _, w := range l.workers {
		if w.Metrics.CPUUsage < minCPU {
			minCPU = w.Metrics.CPUUsage
			chosen = w
		}
	}

	if chosen == nil {
		chosen = l.workers[0]
	}

	chosen.TaskCh <- task
}
