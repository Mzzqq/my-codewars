package main

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	if dadYearsOld < sonYearsOld {
		return 0
	}
	return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
}
