package scheduler

import "adaptive-task-scheduler/internal/worker"

type RoundRobinScheduler struct {
	workers []*worker.Worker
	index   int
}

func NewRoundRobinScheduler(workers []*worker.Worker) *RoundRobinScheduler {
	return &RoundRobinScheduler{workers: workers}
}

func (r *RoundRobinScheduler) Schedule(task func()) {
	w := r.workers[r.index]
	r.index = (r.index + 1) % len(r.workers)
	w.TaskCh <- task
}
