package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	dp := make([][][]*TreeNode, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]*TreeNode, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][i] = []*TreeNode{{Val: i}}
	}
	for s := 1; s <= n; s++ {
		for i := 1; i <= n-s+1; i++ {
			var result []*TreeNode
			for j := i; j <= i+s-1; j++ {
				var lt, rt []*TreeNode
				if i == j {
					lt = []*TreeNode{nil}
				} else {
					lt = dp[i][j-1]
				}
				if j == i+s-1 {
					rt = []*TreeNode{nil}
				} else {
					rt = dp[j+1][i+s-1]
				}
				for _, ln := range lt {
					for _, rn := range rt {
						result = append(result, &TreeNode{Val: j, Left: ln, Right: rn})
					}
				}
			}
			dp[i][i+s-1] = result
		}
	}
	return dp[1][n]
}

func dfs(start, end int) []*TreeNode {
	var result []*TreeNode
	if start > end {
		result = append(result, nil)
		return result
	}
	for i := start; i <= end; i++ {
		l := dfs(start, i-1)
		r := dfs(i+1, end)
		for _, m := range l {
			for _, n := range r {
				result = append(result, &TreeNode{Val: i, Left: m, Right: n})
			}
		}
	}
	return result
}

// @lc code=end
