package main

import "testing"

func Test0(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(heights)
	if area != 10 {
		t.Fatal(area)
	}
}
