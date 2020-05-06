package main

import "testing"

func Test0(t *testing.T) {
	nums, target := []int{2, 5, 6, 0, 0, 1, 2}, 0
	if !search(nums, target) {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	nums, target := []int{2, 5, 6, 0, 0, 1, 2}, 3
	if search(nums, target) {
		t.Fail()
	}
}
