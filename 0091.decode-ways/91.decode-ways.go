package leetcode

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if valid(s[i-1 : i]) {
			dp[i] = dp[i-1]
		}
		if i > 1 && valid(s[i-2:i]) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func valid(s string) bool {
	switch len(s) {
	case 1:
		return s[0] != 48
	case 2:
		if s[0] == 49 {
			return true
		}
		if s[0] == 50 {
			return s[1] >= 48 && s[1] <= 54
		}
		return false
	default:
		return false
	}
}

// @lc code=end
