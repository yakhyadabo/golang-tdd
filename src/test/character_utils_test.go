package main

import (
	"testing"
	"../main"
)

var  character_utils  = main.Character_Utils{}

func TestIsPalindrome(t *testing.T) {
	result := character_utils.IsPalindrome("Hello")
	if result {
		t.Error("Expected true, got ", result)
	}

}
