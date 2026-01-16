package scheduler

import (
	"container/heap"
	"sync"
	"time"

	"adaptive-task-scheduler/internal/metrics"
	"adaptive-task-scheduler/internal/worker"
)

type Scheduler struct {
	Workers map[string]*worker.Worker
	mu      sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Workers: make(map[string]*worker.Worker),
	}
}

func (s *Scheduler) RegisterWorker(w *worker.Worker) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Workers[w.ID] = w
}

func (s *Scheduler) Schedule(task func()) {
	pq := NewPQ()

	for _, w := range s.Workers {
		score := ComputeScore(w.Metrics)
		heap.Push(pq, Item{
			WorkerID: w.ID,
			Score:    score,
		})
	}

	best := heap.Pop(pq).(Item)
	w := s.Workers[best.WorkerID]

	// 🔴 IMPORTANT: scheduler label = "adaptive"
	metrics.WorkerCPU.
		WithLabelValues("adaptive", w.ID).
		Set(w.Metrics.CPUUsage)

	metrics.WorkerMemory.
		WithLabelValues("adaptive", w.ID).
		Set(w.Metrics.MemoryUsage)

	w.TaskCh <- func() {
		start := time.Now()
		task()
		metrics.TaskLatency.
			WithLabelValues("adaptive").
			Observe(time.Since(start).Seconds())
	}
}
