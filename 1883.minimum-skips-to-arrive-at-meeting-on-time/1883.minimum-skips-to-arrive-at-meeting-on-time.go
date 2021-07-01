package xcode

/*
 * @lc app=leetcode id=1883 lang=golang
 *
 * [1883] Minimum Skips to Arrive at Meeting On Time
 */

// @lc code=start
func minSkips(dist []int, speed int, hoursBefore int) int {
	n := len(dist)
	dp := make([]int, n+1)
	// dp[i] represents after i skips, smallest time*speed
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			dp[j] += dist[i]
			if i < n-1 {
				dp[j] = (dp[j] + speed - 1) / speed * speed
			}

			if j > 0 {
				// dp[j] means between i-1 and i
				// dp[j-1] + dist[i] means skip between i and i+1
				dp[j] = min(dp[j], dp[j-1]+dist[i])
			}

			// how to understand:
			// dp[0] means no skips
			// dp[1] means skip 1 time
			// after loop i from 0 to n-1, dp[1] means shortest distance when skip 1 time
			// so dp[j] means shortest distance when skip j times
		}
	}

	for k := 0; k < n; k++ {
		if dp[k] <= speed*hoursBefore {
			return k
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end
