package main

import "fmt"

func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func RGB(r, g, b int) string {
	// Your code here
	r = clamp(r, 0, 255)
	g = clamp(g, 0, 255)
	b = clamp(b, 0, 255)

	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}
