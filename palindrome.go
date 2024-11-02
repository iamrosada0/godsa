package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(value string) bool {
	word := preprocess(value)
	left := 0
	right := len(word) - 1

	// Verificar se a palavra é um palíndromo
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func preprocess(value string) string {
	var pretransform strings.Builder

	for _, char := range value {
		// Adiciona apenas caracteres alfanuméricos
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			pretransform.WriteRune(unicode.ToUpper(char))
		}
	}
	return pretransform.String()
}

func main() {
	testWords := []string{"A man, a plan, a canal, Panama!", "Not a palindrome", "Racecar", "12321", "Hello, World!"}

	for _, word := range testWords {
		if isPalindrome(word) {
			fmt.Printf("\"%s\" is a palindrome.\n", word)
		} else {
			fmt.Printf("\"%s\" is not a palindrome.\n", word)
		}
	}
}
