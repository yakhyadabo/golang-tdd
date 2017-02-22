package main

import (
	"strings"
)

type Character_Utils struct {
}

func (Character_Utils) IsPalindrome(word string) bool {
	reverse := ""
	for i := len(word) - 1; i >= 0; i-- {
		reverse += string(word[i])
	}
	return strings.ToUpper(reverse) == strings.ToUpper(word)
}

func (ch Character_Utils) IsPalindromeR(word string) bool {
	if len(word) <= 1 {
		return true
	}

	if strings.ToUpper(string(word[len(word)-1])) == strings.ToUpper(string(word[0])) {
		return ch.IsPalindromeR(word[1:len(word)-1])
	}

	return false
}
