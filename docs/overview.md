
```md
## Overview

The system consists of a central scheduler and multiple workers.

Each worker continuously reports metrics. The scheduler computes a score and assigns tasks dynamically.

This mirrors real-world schedulers like Kubernetes and CI runners.

To validate scheduling decisions, the system runs an adaptive
scheduler and a round-robin scheduler side-by-side on the same
machine. This allows real-time comparison under identical load,
mirroring how scheduling strategies are evaluated in production systems.
