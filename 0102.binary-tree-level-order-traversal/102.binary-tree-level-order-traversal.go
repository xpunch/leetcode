package xcode

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		nums := make([]int, 0)
		children := make([]*TreeNode, 0)
		for _, n := range stack {
			if n == nil {
				continue
			}
			nums = append(nums, n.Val)
			if n.Left != nil {
				children = append(children, n.Left)
			}
			if n.Right != nil {
				children = append(children, n.Right)
			}
		}
		result = append(result, nums)
		stack = children
	}
	return result
}

// @lc code=end
