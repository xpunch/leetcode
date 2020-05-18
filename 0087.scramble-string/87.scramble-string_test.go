package main

import "testing"

func Test0(t *testing.T) {
	s1, s2 := "great", "rgeat"
	if !isScramble(s1, s2) {
		t.Fail()
	}
}
func Test1(t *testing.T) {
	s1, s2 := "abcde", "caebd"
	if isScramble(s1, s2) {
		t.Fail()
	}
}

func Test248(t *testing.T) {
	s1, s2 := "ababcbaccbabbcbca", "bbbbbaaaacccccbba"
	if isScramble(s1, s2) {
		t.Fail()
	}
}
