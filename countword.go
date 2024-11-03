package main

import "strings"

func CountWord(value string) int {

	removeSpace := strings.Fields(strings.TrimSpace(value))

	totalWord := len(removeSpace)
	return totalWord
}
