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

func (ch Character_Utils) toUpperString(char byte) string {
	return strings.ToUpper(string(char))
}

func (ch Character_Utils) IsPalindromeR(word string) bool {

	return len(word) <= 1 ||
			ch.toUpperString(word[len(word)-1]) == ch.toUpperString(word[0]) &&
		 	ch.IsPalindromeR(word[1:len(word)-1])
}
