package main

import "strings"

type Character_Utils struct {
}

func (Character_Utils) IsPalindrome(word string) bool {
	reverse := ""
	for  i := len(word) -1; i>= 0; i-- {
		reverse += string(word[i])
	}

	return strings.ToUpper(reverse) == strings.ToUpper(word)
}