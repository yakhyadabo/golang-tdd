package main

import (
	"testing"
	"../main"
)

var  character_utils  = main.Character_Utils{}

func TestIsPalindrome(t *testing.T) {
	result := character_utils.IsPalindrome("Elle")
	if !result {
		t.Error("Expected true, got ", result)
	}
}

func TestIsNotPalindrome(t *testing.T) {
	result := character_utils.IsPalindrome("Hello")
	if result {
		t.Error("Expected true, got ", result)
	}
}

func TestIsPalindromeR(t *testing.T) {
	result := character_utils.IsPalindromeR("Ele")
	if !result {
		t.Error("Expected true, got ", result)
	}
}

func TestIsNotPalindromeR(t *testing.T) {
	result := character_utils.IsPalindromeR("EllEs")
	if result {
		t.Error("Expected true, got ", result)
	}
}