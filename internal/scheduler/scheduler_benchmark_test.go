package scheduler

import (
	"testing"

	"adaptive-task-scheduler/internal/task"
	"adaptive-task-scheduler/internal/worker"
)

func BenchmarkAdaptive(b *testing.B) {
	s := NewScheduler()

	w1 := worker.NewWorker("a1")
	w2 := worker.NewWorker("a2")

	s.RegisterWorker(w1)
	s.RegisterWorker(w2)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Schedule(func() {
			task.CPUHeavyWork(1_000_000)
		})
	}
}

func BenchmarkRoundRobin(b *testing.B) {
	w1 := worker.NewWorker("r1")
	w2 := worker.NewWorker("r2")

	rr := NewRoundRobinScheduler([]*worker.Worker{w1, w2})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rr.Schedule(func() {
			task.CPUHeavyWork(1_000_000)
		})
	}
}

// Benchmarks compare adaptive scheduling vs round-robin
// using CPU-intensive workloads to highlight load imbalance effects.

// to run benchmark
// go test -bench=. ./internal/scheduler
// go test -bench=. ./internal/scheduler -benchmem
