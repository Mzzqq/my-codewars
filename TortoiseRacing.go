package main

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	time := g * 3600 / (v2 - v1)
	var result [3]int
	result[0] = time / 3600
	result[1] = (time % 3600) / 60
	result[2] = time % 60
	return result
}
