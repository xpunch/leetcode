package xcode

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	stack := []*TreeNode{root}
	for swap := false; len(stack) > 0; swap = !swap {
		nums, children := make([]int, 0), make([]*TreeNode, 0)
		for _, n := range stack {
			if n == nil {
				continue
			}
			if swap {
				nums = append([]int{n.Val}, nums...)
			} else {
				nums = append(nums, n.Val)
			}
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
