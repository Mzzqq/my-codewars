package main

func Tribonacci(signature [3]float64, n int) []float64 {
	result := make([]float64, n)
	if n == 0 {
		return result
	}

	copy(result, signature[:])

	for i := 3; i < n; i++ {
		result[i] = result[i-1] + result[i-2] + result[i-3]
	}

	return result
}
