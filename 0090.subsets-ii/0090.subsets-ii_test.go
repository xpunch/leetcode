package leetcode

import (
	"testing"
)

func Test0(t *testing.T) {
	nums := []int{1, 2, 2}
	result := subsetsWithDup(nums)
	exception := [][]int{{2}, {1}, {1, 2, 2}, {2, 2}, {1, 2}, {}}
	if !compare(result, exception) {
		t.Fatalf("Output: %v\nExpected: %v\n", result, exception)
	}
}

func Test1(t *testing.T) {
	nums := []int{1, 2}
	result := subsetsWithDup(nums)
	exception := [][]int{{}, {1}, {1, 2}, {2}}
	if !compare(result, exception) {
		t.Fatalf("Output: %v\nExpected: %v\n", result, exception)
	}
}

func Test2(t *testing.T) {
	nums := []int{1, 2, 3}
	result := subsetsWithDup(nums)
	exception := [][]int{{}, {1, 2, 3}, {1}, {1, 2}, {1, 3}, {2}, {2, 3}, {3}}
	if !compare(result, exception) {
		t.Fatalf("Output: %v\nExpected: %v\n", result, exception)
	}
}

func Test3(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := subsetsWithDup(nums)
	exception := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 4}, {1, 3}, {1, 3, 4}, {1, 4}, {2}, {2, 3}, {2, 3, 4}, {2, 4}, {3}, {3, 4}, {4}}
	if !compare(result, exception) {
		t.Fatalf("Output: %v\nExpected: %v\n", result, exception)
	}
}

func Test4(t *testing.T) {
	nums := []int{5, 3, 1, 2, 2}
	result := subsetsWithDup(nums)
	exception := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 2, 3, 5}, {1, 2, 2, 5}, {1, 2, 3}, {1, 2, 3, 5}, {1, 2, 5}, {1, 3}, {1, 3, 5}, {1, 5}, {2}, {2, 2}, {2, 2, 3}, {2, 2, 3, 5}, {2, 2, 5}, {2, 3}, {2, 3, 5}, {2, 5}, {3}, {3, 5}, {5}}
	if !compare(result, exception) {
		t.Fatalf("Output: %v\nExpected: %v\n", result, exception)
	}
}

func Test13(t *testing.T) {
	nums := []int{2, 1, 2, 1, 3}
	result := subsetsWithDup(nums)
	exception := [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 3}, {1, 1, 2, 3}, {1, 1, 3}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3}, {1, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 3}, {3}}
	if !compare(result, exception) {
		t.Fatalf("Output: %v\nExpected: %v\n", result, exception)
	}
}

func compare(result, exception [][]int) bool {
	if len(result) != len(exception) {
		return false
	}
	for _, r := range result {
		found := false
		for _, e := range exception {
			if len(r) != len(e) {
				continue
			}
			matched := true
			for i := 0; i < len(r); i++ {
				if r[i] != e[i] {
					matched = false
					break
				}
			}
			if matched {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	for _, r := range exception {
		found := false
		for _, e := range result {
			if len(r) != len(e) {
				continue
			}
			matched := true
			for i := 0; i < len(r); i++ {
				if r[i] != e[i] {
					matched = false
					break
				}
			}
			if matched {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
