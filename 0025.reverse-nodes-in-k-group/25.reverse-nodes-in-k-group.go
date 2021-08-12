package main

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	root, left := head, head
	var pre *ListNode
	for left != nil && left.Next != nil {
		i, right := 1, left.Next
		for ; i < k-1 && right.Next != nil; i++ {
			right = right.Next
		}
		if i < k-1 {
			left = nil
			break
		}
		tmp := right.Next
		if root == head {
			root = right
		}
		if pre != nil {
			pre.Next = right
		}
		reverseRange(left, right)
		pre = left
		left = tmp
	}
	return root
}

func reverseRange(head *ListNode, tail *ListNode) {
	cur := head
	for cur != tail {
		tmp := head.Next.Next
		head.Next.Next = cur
		cur = head.Next
		head.Next = tmp
	}
}

// @lc code=end
