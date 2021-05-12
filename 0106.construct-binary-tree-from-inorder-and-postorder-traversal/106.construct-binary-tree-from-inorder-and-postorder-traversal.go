package xcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	inMap := make(map[int]int)
	for i, v := range inorder {
		inMap[v] = i
	}
	return dfs(postorder, inMap, 0, len(postorder)-1, 0)
}

func dfs(postorder []int, inorder map[int]int, postStart, postEnd, inStart int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &TreeNode{Val: postorder[postEnd]}
	leftLen := inorder[postorder[postEnd]] - inStart
	root.Left = dfs(postorder, inorder, postStart, postStart+leftLen-1, inStart)
	root.Right = dfs(postorder, inorder, postStart+leftLen, postEnd-1, inorder[postorder[postEnd]]+1)
	return root
}

// @lc code=end
