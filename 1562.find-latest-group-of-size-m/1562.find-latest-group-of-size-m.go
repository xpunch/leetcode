/*
 * @lc app=leetcode id=1562 lang=golang
 *
 * [1562] Find Latest Group of Size M
 */
package xcode

// @lc code=start
func findLatestStep(arr []int, m int) int {
	n := len(arr)
	if m == n {
		return n
	}
	bits := make([]int, n+1)
	counts := make([]int, n+1)
	ans := -1
	for j, i := range arr {
		l, r := 0, 0
		if i > 1 && bits[i-1] > 0 {
			l = bits[i-1]
		}
		if i < n && bits[i+1] > 0 {
			r = bits[i+1]
		}
		bits[i] = l + r + 1
		counts[bits[i]]++
		if i > 1 && l > 0 {
			bits[i-1] = bits[i]
			if l > 1 {
				bits[i-l] = bits[i]
			}
			counts[l]--
		}
		if i < n && r > 0 {
			bits[i+1] = bits[i]
			if r > 1 {
				bits[i+r] = bits[i]
			}
			counts[r]--
		}
		if counts[m] > 0 {
			ans = j + 1
		}
	}
	return ans
}

// @lc code=end
