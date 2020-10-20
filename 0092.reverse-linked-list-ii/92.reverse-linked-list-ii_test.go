package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test0(t *testing.T) {
	head := buildLinkedList([]int{1, 2, 3, 4, 5})
	reverse := reverseBetween(head, 2, 4)
	result := convert(reverse)
	if !reflect.DeepEqual(result, []int{1, 4, 3, 2, 5}) {
		t.Fatal(result)
	}
}

func Test1(t *testing.T) {
	head := buildLinkedList([]int{1, 2, 3, 4, 5})
	reverse := reverseBetween(head, 1, 4)
	result := convert(reverse)
	if !reflect.DeepEqual(result, []int{4, 3, 2, 1, 5}) {
		t.Fatal(result)
	}
}

func Test2(t *testing.T) {
	head := buildLinkedList([]int{1, 2, 3, 4, 5})
	reverse := reverseBetween(head, 1, 5)
	result := convert(reverse)
	if !reflect.DeepEqual(result, []int{5, 4, 3, 2, 1}) {
		t.Fatal(result)
	}
}

func buildLinkedList(nums []int) *ListNode {
	var root, cur *ListNode
	for i, n := range nums {
		if i == 0 {
			cur = &ListNode{Val: n}
			root = cur
		} else {
			cur.Next = &ListNode{Val: n}
			cur = cur.Next
		}
	}
	return root
}

func convert(head *ListNode) []int {
	tmp := make(map[*ListNode]struct{})
	nums, cur := make([]int, 0), head
	for cur != nil {
		if _, ok := tmp[cur]; ok {
			fmt.Printf("Cycle: %v %v\n", nums, cur)
			break
		}
		tmp[cur] = struct{}{}
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	return nums
}

// @lc code=end
