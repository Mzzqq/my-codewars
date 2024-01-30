package main

import "regexp"

func Disemvowel(comment string) string {
	vocal := regexp.MustCompile("[aiueoAIUEO]")
	return vocal.ReplaceAllString(comment, "")
}
