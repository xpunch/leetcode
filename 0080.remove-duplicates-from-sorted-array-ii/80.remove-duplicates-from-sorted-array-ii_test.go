package main

import (
	"log"
	"testing"
)

func Test0(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	result := removeDuplicates(nums)
	if result != 5 {
		t.Fatal(result)
	}
	log.Print(nums)
}

func Test1(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	result := removeDuplicates(nums)
	if result != 7 {
		t.Fatal(result)
	}
	log.Print(nums)
}
