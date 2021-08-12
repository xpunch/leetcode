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
	var newHead *ListNode
	var first *ListNode
	for cur := head; cur != nil; {
		i, nodes := 0, make([]*ListNode, k)
		for ; i < k && cur != nil; i++ {
			nodes[i] = cur
			cur = cur.Next
		}
		if i < k {
			if first != nil {
				first.Next = nodes[0]
			}
			break
		}
		if first != nil {
			first.Next = nodes[k-1]
		}
		if newHead == nil {
			newHead = nodes[k-1]
		}
		for j := k - 1; j > 0; j-- {
			nodes[j].Next = nodes[j-1]
		}
		if cur == nil {
			nodes[0].Next = nil
		} else {
			first = nodes[0]
		}
	}
	return newHead
}

// @lc code=end
