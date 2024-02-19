package main

import "strings"

func wave(words string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		if words[i] != ' ' {
			up := strings.ToUpper(string(words[i]))
			wave := words[:i] + up + words[i+1:]
			result = append(result, wave)
		}
	}
	return result
}
