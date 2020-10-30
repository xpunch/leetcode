/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for s := 1; s <= n; s++ {
		for i := 1; i <= n-s+1; i++ {
			var count int
			for j := i; j <= i+s-1; j++ {
				var l, r int
				if i == j {
					l = 1
				} else {
					l = dp[i][j-1]
				}
				if j == i+s-1 {
					r = 1
				} else {
					r = dp[j+1][i+s-1]
				}
				count += l * r
			}
			dp[i][i+s-1] = count
		}
	}
	return dp[1][n]
}

// @lc code=end

