package task

import "math"

func CPUHeavyWork(iterations int) {
	x := 0.0001
	for i := 0; i < iterations; i++ {
		x += math.Sqrt(x)
	}
}
