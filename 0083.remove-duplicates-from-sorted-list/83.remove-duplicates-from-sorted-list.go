/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root, pre, val, next := head, head, head.Val, head.Next
	for next != nil {
		if val == next.Val {
		} else {
			pre.Next = next
			pre = next
			val = next.Val
		}
		next = next.Next
	}
	if pre != nil && pre.Next != nil {
		pre.Next = nil
	}
	return root
}

// @lc code=end
