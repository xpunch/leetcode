package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	input := "babad"
	output := longestPalindrome(input)
	if output != "bab" {
		t.Fatalf("%s %s", input, output)
	}
}
