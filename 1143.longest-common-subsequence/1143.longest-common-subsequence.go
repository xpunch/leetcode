package xcode

/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == text2 {
		return len(text1)
	}
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[n1][n2]
}

// @lc code=end
