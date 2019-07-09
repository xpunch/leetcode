package main

import "testing"

func Test0(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	if len(result) != 2 {
		t.Fail()
	}
}
