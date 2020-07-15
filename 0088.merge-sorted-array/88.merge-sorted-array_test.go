package main

import (
	"fmt"
	"testing"
)

func Test0(t *testing.T) {
	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	for i := 0; i < m+n-2; i++ {
		if nums1[i] > nums1[i+1] {
			t.Fail()
		}
	}
}

func Test1(t *testing.T) {
	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0, 0}, 3, []int{4, 5, 6}, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	for i := 0; i < m+n-2; i++ {
		if nums1[i] > nums1[i+1] {
			t.Fail()
		}
	}
}

func Test2(t *testing.T) {
	nums1, m, nums2, n := []int{7, 9, 23, 0, 0, 0, 0}, 3, []int{4, 5, 6}, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	for i := 0; i < m+n-2; i++ {
		if nums1[i] > nums1[i+1] {
			t.Fail()
		}
	}
}
