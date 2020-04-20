package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	n, k := 4, 2
	result := combine(n, k)
	if len(result) != cmn(n, k) {
		t.Fail()
	}
}

func Test18(t *testing.T) {
	n, k := 5, 4
	result := combine(n, k)
	fmt.Println(result)
	if len(result) != cmn(n, k) {
		t.Fatal(result)
	}
	if !equal(result[0], []int{1, 2, 3, 4}) {
		t.Fatal(result)
	}
}

func cmn(m, n int) int {
	tmp := 1
	for i := 0; i < n; i++ {
		tmp *= (m - i)
	}
	for i := n; i > 1; i-- {
		tmp /= i
	}
	return tmp
}

func equal(source, target []int) bool {
	if len(source) != len(target) {
		return false
	}
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			return false
		}
	}
	return true
}
