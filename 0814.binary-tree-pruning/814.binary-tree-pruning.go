/*
 * @lc app=leetcode id=814 lang=golang
 *
 * [814] Binary Tree Pruning
 */
package xcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func pruneTree(root *TreeNode) *TreeNode {
	if satisfy(root) {
		return root
	}
	return nil
}

func satisfy(node *TreeNode) bool {
	if node == nil {
		return false
	}
	l, r := satisfy(node.Left), satisfy(node.Right)
	if node.Val == 0 && !l && !r {
		return false
	}
	if !l {
		node.Left = nil
	}
	if !r {
		node.Right = nil
	}
	return true
}

// @lc code=end
