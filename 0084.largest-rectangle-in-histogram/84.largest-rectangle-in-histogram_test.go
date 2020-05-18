package main

import "testing"

func Test0(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(heights)
	if area != 10 {
		t.Fatal(area)
	}
}
func Test26(t *testing.T) {
	heights := []int{0, 9}
	area := largestRectangleArea(heights)
	if area != 9 {
		t.Fatalf("%v %d", heights, area)
	}
}

func Test53(t *testing.T) {
	heights := []int{2, 1, 2}
	area := largestRectangleArea(heights)
	if area != 3 {
		t.Fatal(area)
	}
}
