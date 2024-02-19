package main

func multipleOfIndex(ints []int) []int {
	result := []int{}

	for i, val := range ints {
		if i != 0 && val%i == 0 {
			result = append(result, val)
		}
	}

	return result
}
