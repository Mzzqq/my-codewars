package main

import "sort"

func Gimme(array [3]int) int {
	sortedArray := make([]int, 3)
	copy(sortedArray, array[:])

	sort.Ints(sortedArray)

	middleElement := sortedArray[1]

	for i, num := range array {
		if num == middleElement {
			return i
		}
	}

	return -1
}
