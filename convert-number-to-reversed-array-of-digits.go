package main

func Digitize(n int) []int {
	result := []int{}

	if n == 0 {
		return []int{0}
	}

	for n > 0 {
		result = append(result, n%10)
		n /= 10
	}
	return result
}
