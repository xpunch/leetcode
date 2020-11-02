package main

/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	if len(s3) == 0 {
		return true
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 {
				if s3[i+j-1] == s1[i-1] && dp[i-1][j] {
					dp[i][j] = true
					continue
				}
			}
			if j > 0 {
				if s3[i+j-1] == s2[j-1] && dp[i][j-1] {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func dfs(s1, s2, s3 string) bool {
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	if len(s3) == 0 {
		return true
	}
	if s3[0] == s1[0] {
		if dfs(s1[1:], s2, s3[1:]) {
			return true
		}
	}
	if s3[0] == s2[0] {
		return dfs(s1, s2[1:], s3[1:])
	}
	return false
}

// @lc code=end
