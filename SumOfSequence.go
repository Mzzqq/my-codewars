package main

func SequenceSum(start, end, step int) int {
	total := 0
	for i := start; i <= end; i += step {
		total += i
	}
	return total

}
