package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	cur := head
	var root, mPreNode, mNode, pre *ListNode
	for i := 1; i <= n; i++ {
		if i < m-1 {
			cur = cur.Next
			continue
		}
		if i == m-1 {
			root = head
			mPreNode = cur
			cur = cur.Next
			continue
		}
		next := cur.Next
		if i == m {
			mNode = cur
		} else {
			cur.Next = pre
		}
		if i == n {
			if root == nil {
				root = cur
			} else {
				mPreNode.Next = cur
			}
			mNode.Next = next
		}
		pre = cur
		cur = next
	}
	return root
}

// @lc code=end
