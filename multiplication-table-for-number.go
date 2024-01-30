package main

import "fmt"

func MultiTable(number int) string {
	result := ""
	for i := 1; i <= 10; i++ {
		result += fmt.Sprint(i, " * ", number, " = ", i*number)
		if i < 10 {
			result += "\n"
		}
	}
	return result
}
