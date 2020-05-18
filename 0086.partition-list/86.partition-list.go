package main

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var left, pre, right *ListNode
	root, cur := head, head
	for cur != nil {
		next := cur.Next
		if cur.Val >= x {
			if right == nil {
				right = cur
			}
			pre = cur
		} else {
			if right != nil {
				if left == nil {
					root = cur
				} else {
					left.Next = cur
				}
				if pre != nil {
					pre.Next = cur.Next
				}
			} else {
				pre = cur
			}
			left = cur
			if right != nil {
				left.Next = right
			}
		}
		cur = next
	}
	return root
}

// @lc code=end
