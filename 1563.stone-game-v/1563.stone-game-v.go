package xcode

/*
 * @lc app=leetcode id=1563 lang=golang
 *
 * [1563] Stone Game V
 */

// @lc code=start
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = stoneValue[i]
		if i != n-1 {
			if stoneValue[i+1] > stoneValue[i] {
				dp[i][i+1] = stoneValue[i]
			} else {
				dp[i][i+1] = stoneValue[i+1]
			}
		}
	}
	for i := n - 2; i >= 0; i-- {
		left, right := stoneValue[i], stoneValue[i+1]
		for j := i + 2; j < n; j++ {
			right += stoneValue[j]
			max, left2, right2 := 0, left, right
			for k := i; k < j; k++ {
				if left2 > right2 {
					r := dp[k+1][j]
					if j-k > 1 {
						r += right2
					}
					if r > max {
						max = r
					}
				} else if left2 < right2 {
					l := dp[i][k]
					if k-i > 0 {
						l += left2
					}
					if l > max {
						max = l
					}
				} else {
					l, r := dp[i][k], dp[k+1][j]
					if k-i > 0 {
						l += left2
					}
					if j-k > 1 {
						r += right2
					}
					if l >= r {
						max = l
					} else {
						max = r
					}
				}
				left2 += stoneValue[k+1]
				right2 -= stoneValue[k+1]
			}
			dp[i][j] = max
		}
	}
	return dp[0][n-1]
}

// @lc code=end
