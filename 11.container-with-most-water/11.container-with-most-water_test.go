package main

import "testing"

func Test0(t *testing.T) {
	if maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) != 49 {
		t.Fail()
	}
}
