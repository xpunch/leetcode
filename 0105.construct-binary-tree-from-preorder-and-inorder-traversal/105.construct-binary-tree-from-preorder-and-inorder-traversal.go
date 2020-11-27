package xcode

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	return dfs(preorder, inorderMap, 0, 0, len(preorder))
}

func dfs(preorder []int, inorderMap map[int]int, i, l, r int) *TreeNode {
	root := &TreeNode{Val: preorder[i]}
	rootIndex := inorderMap[preorder[i]]
	if rootIndex > l {
		root.Left = dfs(preorder, inorderMap, i+1, l, rootIndex)
	}
	if rootIndex+1 < r {
		root.Right = dfs(preorder, inorderMap, i+1+rootIndex-l, rootIndex+1, r)
	}
	return root
}

// @lc code=end
