package main

import "testing"

func Test0(t *testing.T) {
	if isMatch("aa", "a") {
		t.Fail()
	}
}
func Test1(t *testing.T) {
	if !isMatch("aab", "c*a*b*") {
		t.Fail()
	}
}
func Test4(t *testing.T) {
	if isMatch("mississippi", "mis*is*p*.") {
		t.Fail()
	}
}
func Test5(t *testing.T) {
	if isMatch("aaa", "aa") {
		t.Fail()
	}
}
func Test6(t *testing.T) {
	if !isMatch("aa", "a*") {
		t.Fail()
	}
}
func Test9(t *testing.T) {
	if isMatch("abcd", "d*") {
		t.Fail()
	}
}
func Test10(t *testing.T) {
	if isMatch("ab", ".*c") {
		t.Fail()
	}
}
func Test13(t *testing.T) {
	if isMatch("aaa", "aaaa") {
		t.Fail()
	}
}
func Test38(t *testing.T) {
	if !isMatch("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s") {
		t.Fail()
	}
}
func Test45(t *testing.T) {
	if !isMatch("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s") {
		t.Fail()
	}
}
func Test46(t *testing.T) {
	if isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c") {
		t.Fail()
	}
}
func Test86(t *testing.T) {
	if isMatch("bbab", "b*a*") {
		t.Fail()
	}
}
func Test247(t *testing.T) {
	if isMatch("bbabacccbcbbcaaab", "a*b*a*a*c*aa*c*bc*") {
		t.Fail()
	}
}
func Test407(t *testing.T) {
	if !isMatch("bbbba", ".*a*a") {
		t.Fail()
	}
}
func Test418(t *testing.T) {
	if !isMatch("aaa", "a*a") {
		t.Fail()
	}
}
func Test439(t *testing.T) {
	if !isMatch("cabbbbcbcacbabc", ".*b.*.ab*.*b*a*c") {
		t.Fail()
	}
}
func Test446(t *testing.T) {
	if !isMatch("abbaaaabaabbcba", "a*.*ba.*c*..a*.a*.") {
		t.Fail()
	}
}
