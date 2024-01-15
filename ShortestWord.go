package main

import "strings"

func FindShort(s string) int {
	result := 0
	splitWord := strings.Split(s, " ")
	for i := 0; i < len(splitWord); i++ {
		if result == 0 {
			result = len(splitWord[i])
		} else {
			if result > len(splitWord[i]) {
				result = len(splitWord[i])
			}
		}
	}
	return result
}
