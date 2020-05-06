/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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
	var root, pre *ListNode
	val, next, cur, duplicated := head.Val, head.Next, head, false
	for next != nil {
		if val == next.Val {
			duplicated = true
		} else {
			if !duplicated {
				if root == nil {
					root = cur
				}
				if pre == nil {
					pre = cur
				} else {
					pre.Next = cur
					pre = cur
				}
			}
			cur = next
			duplicated = false
			val = next.Val
		}
		next = next.Next
	}
	if duplicated {
		if pre != nil {
			pre.Next = nil
		}
	} else {
		if root == nil {
			root = cur
		}
		if pre != nil {
			pre.Next = cur
		}
	}
	return root
}

// @lc code=end
