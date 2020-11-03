package xcode

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(left, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}
	if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

// @lc code=end
