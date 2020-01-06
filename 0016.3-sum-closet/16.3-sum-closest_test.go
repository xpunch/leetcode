package main

import "testing"

func TestThreeSumClosest(t *testing.T) {
	input := []int{1, 2, 4, 8, 16, 32, 64, 128}
	target := 82
	closest := threeSumClosest(input, target)
	if closest != 82 {
		t.Fatal(closest)
	}
}
