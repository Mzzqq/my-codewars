package main

func Divisors(n int) int {
	counter := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			counter += 1
		}
	}
	return counter
}
