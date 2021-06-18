package xcode

import "testing"

func Test1(t *testing.T) {
	nums, goal := []int{5, -7, 3, 5}, 6
	if output := minAbsDifference(nums, goal); output != 0 {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	nums, goal := []int{7, -9, 15, -2}, -5
	if output := minAbsDifference(nums, goal); output != 1 {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	nums, goal := []int{1, 2, 3}, -7
	if output := minAbsDifference(nums, goal); output != 7 {
		t.Fatal(output)
	}
}
