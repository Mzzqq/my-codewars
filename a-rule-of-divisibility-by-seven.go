package main

func Seven(n int64) []int {
	step := 0
	for n >= 100 {
		step++
		lastDigit := n % 10
		n = (n - lastDigit) / 10
		n -= 2 * lastDigit
	}
	return []int{int(n), step}
}
