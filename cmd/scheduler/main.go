package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"adaptive-task-scheduler/internal/metrics"
	"adaptive-task-scheduler/internal/scheduler"
	"adaptive-task-scheduler/internal/task"
	"adaptive-task-scheduler/internal/worker"
)

func main() {
	// -------------------------------
	// Register Prometheus metrics
	// -------------------------------
	prometheus.MustRegister(
		metrics.WorkerCPU,
		metrics.WorkerMemory,
		metrics.TaskLatency,
	)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	// -------------------------------
	// Adaptive Scheduler setup
	// -------------------------------
	adaptive := scheduler.NewScheduler()

	aw1 := worker.NewWorker("a-1")
	aw2 := worker.NewWorker("a-2")

	adaptive.RegisterWorker(aw1)
	adaptive.RegisterWorker(aw2)

	// -------------------------------
	// Round-Robin Scheduler setup
	// -------------------------------
	rw1 := worker.NewWorker("r-1")
	rw2 := worker.NewWorker("r-2")

	rr := scheduler.NewRoundRobinScheduler(
		[]*worker.Worker{rw1, rw2},
	)

	// -------------------------------
	// Least-Loaded Scheduler setup
	// -------------------------------
	lw1 := worker.NewWorker("l-1")
	lw2 := worker.NewWorker("l-2")

	leastLoaded := scheduler.NewLeastLoadedScheduler(
		[]*worker.Worker{lw1, lw2},
	)

	// -------------------------------
	// Shared workload generator
	// -------------------------------
	go func() {
		for i := 0; ; i++ {
			i := i

			// work := func() {
			// 	if i%3 == 0 {
			// 		// task.CPUHeavyWork(50_000_000)
			// 		task.CPUHeavyWork(150_000_000)

			// 	} else {
			// 		time.Sleep(100 * time.Millisecond)
			// 	}
			// }

			work := func() {
				// Every 5th task is VERY heavy
				if i%5 == 0 {
					task.CPUHeavyWork(300_000_000)
					return
				}

				// Normal tasks
				if i%3 == 0 {
					task.CPUHeavyWork(150_000_000)
				} else {
					time.Sleep(100 * time.Millisecond)
				}
			}

			// Adaptive (worker + latency metrics handled inside scheduler)
			adaptive.Schedule(func() {
				work()
			})

			// Round-Robin (latency only)
			rr.Schedule(func() {
				start := time.Now()
				work()
				metrics.TaskLatency.
					WithLabelValues("round_robin").
					Observe(time.Since(start).Seconds())
			})

			// Least-Loaded (latency only)
			leastLoaded.Schedule(func() {
				start := time.Now()
				work()
				metrics.TaskLatency.
					WithLabelValues("least_loaded").
					Observe(time.Since(start).Seconds())
			})

			time.Sleep(300 * time.Millisecond)
		}
	}()

	// -------------------------------
	// Keep app alive forever
	// -------------------------------
	select {}
}

//replace the below go function to run the tasks continuously
// go func() {
// 	for i := 0; ; i++ {
// 		i := i
// 		s.Schedule(func() {
// 			if i%3 == 0 {
// 				task.CPUHeavyWork(50_000_000)
// 			} else {
// 				time.Sleep(100 * time.Millisecond)
// 			}
// 		})
// 		time.Sleep(300 * time.Millisecond)
// 	}
// }()
