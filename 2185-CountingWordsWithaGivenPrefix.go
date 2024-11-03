package main

import "strings"

func prefixCount(words []string, pref string) int {

	count := 0
	for _, value := range words {

		if strings.HasPrefix(value, pref) {
			count++
		}
	}

	return count

}
