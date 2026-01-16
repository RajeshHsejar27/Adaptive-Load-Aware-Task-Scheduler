# Adaptive Load-Aware Task Scheduler

A production-inspired Go scheduler that dynamically assigns tasks using real-time CPU, memory, and latency feedback.

## Problem
Traditional schedulers (round-robin, FIFO) ignore system load, causing hotspots and high tail latency.

## Solution
This scheduler computes a live score for each worker and always schedules tasks to the least-loaded worker.

## Features
- Load-aware scheduling
- Priority queue based worker selection
- Concurrent execution with goroutines
- Prometheus metrics support
- Extensible scheduling policies

## Tech Stack
- Go
- Goroutines & Channels
- Heap-based priority queues
- Prometheus

## Run
```bash
# Start scheduler
# This starts ONE Go process on your laptop.
# Inside that process:
# Multiple goroutines run concurrently
# No VMs, no Docker, no cloud
go run cmd/scheduler/main.go
```
At startup, the scheduler does:
Creates a Scheduler object
Creates Workers (logical workers, not OS processes)
Registers workers with the scheduler
Starts Prometheus metrics server (:2112)
Workers ≠ CPU cores
Workers = logical execution units managed by Go
Start Prometheus
./prometheus --config.file=prometheus.yml

# (Optional) Start Grafana
grafana-server
# Run benchmarks
go test -bench=. ./internal/scheduler


After enough tasks run:

    Scheduler calculates average latency

    If latency rises:

        Increases importance of latency

        Decreases CPU/memory weight

    If system is stable:

        CPU/memory regain importance

Result:

    Scheduler self-adjusts

    No human tuning needed

    ## Scheduler Comparison: Adaptive vs Round-Robin

This project includes a live comparison between:
- An adaptive, load-aware scheduler
- A traditional round-robin scheduler

Both schedulers:
- Run in parallel
- Execute identical workloads
- Compete for the same system resources

### Observability
Metrics are exposed via Prometheus and visualized in Grafana:
- Per-worker CPU usage
- Task throughput
- p95 task latency

Grafana dashboards clearly show that the adaptive scheduler
reduces tail latency under CPU contention compared to round-robin.
