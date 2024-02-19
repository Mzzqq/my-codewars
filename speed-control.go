package main

import "math"

func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0
	}

	maxSpeed := 0.0
	for i := 1; i < len(x); i++ {
		deltaDistance := x[i] - x[i-1]
		speed := (3600 * deltaDistance) / float64(s)
		maxSpeed = math.Max(maxSpeed, speed)
	}

	return int(math.Floor(maxSpeed))
}
