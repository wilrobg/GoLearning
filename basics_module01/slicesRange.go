package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Text is palindrome")
	} else {
		fmt.Println("Is not palindrome")
	}
}

func main() {
	slice := []string{"what", "you", "doing"}

	for i, value := range slice {
		fmt.Println(i, value)
	}

	isPalindrome("Ama")
}
