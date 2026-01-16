// Each worker is: A Go struct with:
// A task channel
// A metrics snapshot
// Running in its own goroutine

package worker

import (
	"time"
)

type Worker struct {
	ID      string
	Metrics Metrics
	TaskCh  chan func()
}

func NewWorker(id string) *Worker {
	w := &Worker{
		ID:     id,
		TaskCh: make(chan func()),
	}
	go w.start()
	return w
}

func (w *Worker) start() {
	for task := range w.TaskCh {
		start := time.Now()

		task()

		w.Metrics.Latency = time.Since(start).Seconds()
		load := CollectMetrics()
		w.Metrics.CPUUsage = load.CPUUsage
		w.Metrics.MemoryUsage = load.MemoryUsage
	}
}
