package xcode

import "testing"

func Test1(t *testing.T) {
	low, high := 100, 300
	if output := sequentialDigits(low, high); !equal(output, []int{123, 234}) {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	low, high := 10, 1000000000
	if output := sequentialDigits(low, high); !equal(output, []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789, 1234, 2345, 3456, 4567, 5678, 6789, 12345, 23456, 34567, 45678, 56789, 123456, 234567, 345678, 456789, 1234567, 2345678, 3456789, 12345678, 23456789, 123456789}) {
		t.Fatal(output)
	}
}

func equal(s, t []int) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}
