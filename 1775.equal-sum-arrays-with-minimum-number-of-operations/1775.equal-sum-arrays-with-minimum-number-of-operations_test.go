package xcode

import "testing"

func Test1(t *testing.T) {
	nums1, nums2 := []int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}
	if output := minOperations(nums1, nums2); output != 3 {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	nums1, nums2 := []int{1, 1, 1, 1, 1, 1, 1}, []int{6}
	if output := minOperations(nums1, nums2); output != -1 {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	nums1, nums2 := []int{6, 6}, []int{1}
	if output := minOperations(nums1, nums2); output != 3 {
		t.Fatal(output)
	}
}

func Test4(t *testing.T) {
	nums1, nums2 := []int{1}, []int{6, 6}
	if output := minOperations(nums1, nums2); output != 3 {
		t.Fatal(output)
	}
}
