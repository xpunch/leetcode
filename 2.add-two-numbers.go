/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root, parent *ListNode
	carry := false
	for l1 != nil || l2 != nil {
		node := new(ListNode)
		if parent != nil {
			parent.Next = node
			parent = node
		}
		if root == nil {
			root = node
			parent = node
		}
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry {
			sum++
			carry = false
		}
		if sum >= 10 {
			sum = sum - 10
			carry = true
		}
		node.Val = sum
	}
	if carry {
		parent.Next = &ListNode{Val: 1, Next: nil}
	}
	return root
}
