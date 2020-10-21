package leetcode

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	// 1 recursive
	// if root != nil {
	// 	result = append(result, inorderTraversal(root.Left)...)
	// 	result = append(result, root.Val)
	// 	result = append(result, inorderTraversal(root.Right)...)
	// }

	// 2 iterative
	stack, cur := make([]*TreeNode, 0), root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}
	return result
}

// @lc code=end
