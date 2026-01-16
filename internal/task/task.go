package task

type Task struct {
	ID       string
	Priority int
	Duration int // ms
}

// Task Execution Happens Concurrently
// Once assigned:
// Scheduler → Worker Channel → Goroutine executes task
// Worker executes task
// Measures how long it took
// Updates its metrics
// Reports metrics to Prometheus
// All workers can run tasks in parallel.
