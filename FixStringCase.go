package main

import "strings"

func solve(str string) string {
	upperCount := 0
	lowerCount := 0

	for _, word := range str {
		if word >= 'A' && word <= 'Z' {
			upperCount += 1
		} else if word >= 'a' && word <= 'z' {
			lowerCount += 1
		}
	}

	if upperCount > lowerCount {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}

}
