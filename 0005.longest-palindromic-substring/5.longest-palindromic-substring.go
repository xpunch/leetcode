package main

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	l, r := 0, 0
	dp := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n, n)
	}
	for j := 1; j <= n-1; j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i > r-l {
					l = i
					r = j
				}
			}
		}
	}
	return s[l : r+1]
}
