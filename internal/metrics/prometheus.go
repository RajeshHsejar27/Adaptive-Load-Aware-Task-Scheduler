package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	WorkerCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "worker_cpu_usage",
			Help: "CPU usage per worker",
		},
		[]string{"scheduler", "worker"},
	)

	WorkerMemory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "worker_memory_usage",
			Help: "Memory usage per worker",
		},
		[]string{"scheduler", "worker"},
	)

	TaskLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "task_latency_seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"scheduler"},
	)
)
