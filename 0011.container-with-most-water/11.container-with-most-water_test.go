package main

import "testing"

func Test6(t *testing.T) {
	if maxArea([]int{1, 1}) != 1 {
		t.Fail()
	}
}
func Test16(t *testing.T) {
	h := []int{1, 2, 4, 3}
	r := maxArea(h)
	t.Logf("%v: %d", h, r)
	if r != 49 {
		t.Fail()
	}
}
