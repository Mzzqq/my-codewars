package main

func GetSize(w, h, l int) [2]int {
	result := [2]int{}
	result[0] = ((2 * l * w) + (2 * l * h) + (2 * w * h))
	result[1] = w * h * l
	return result
}
