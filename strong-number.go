package main

func Strong(n int) string {
	sum := 0
	original := n
	for n > 0 {
		factorial := 1
		for i := 1; i <= n%10; i++ {
			factorial *= i
		}
		sum += factorial
		n /= 10
	}
	if sum == original {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}
