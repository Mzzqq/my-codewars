package main

import "strings"

func Accum(s string) string {
	var result string

	for i, char := range s {
		repeated := strings.Repeat(strings.ToLower(string(char)), i)
		result += strings.ToUpper(string(char)) + repeated

		if i != len(s)-1 {
			result += "-"
		}
	}

	return result
}
