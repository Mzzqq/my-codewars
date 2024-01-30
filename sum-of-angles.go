package main

func Angle(n int) int {
	if n < 3 {
		return 0
	}

	return (n - 2) * 180
}
