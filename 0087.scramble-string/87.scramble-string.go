package main

/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	n := len(s1)
	dp := make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			for k := 1; k <= n-max(i, j); k++ {
				if s1[i:i+k] == s2[j:j+k] {
					dp[i][j][k] = true
				} else {
					for l := 1; l < k; l++ {
						if (dp[i][j][l] && dp[i+l][j+l][k-l]) || (dp[i][j+k-l][l] && dp[i+l][j][k-l]) {
							dp[i][j][k] = true
						}
					}
				}
			}
		}
	}
	return dp[0][0][n]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
