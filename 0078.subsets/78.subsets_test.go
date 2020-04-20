package main

import (
	"fmt"
	"testing"
)

func Test0(t *testing.T) {
	input := []int{4, 7, 9}
	output := subsets(input)
	fmt.Println(output)
	if len(output) != count(input) {
		t.Fatal(output)
	}
}

func Test1(t *testing.T) {
	input := []int{1, 2, 3}
	output := subsets(input)
	fmt.Println(output)
	if len(output) != count(input) {
		t.Fatal(output)
	}
}

func count(input []int) int {
	result, n := 2, len(input)
	for i := 1; i < n; i++ {
		result += cmn(n, i)
	}
	return result
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
