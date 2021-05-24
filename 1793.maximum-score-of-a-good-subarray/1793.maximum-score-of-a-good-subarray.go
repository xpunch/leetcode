package xcode

/*
 * @lc app=leetcode id=1793 lang=golang
 *
 * [1793] Maximum Score of a Good Subarray
 */

// @lc code=start
func maximumScore(nums []int, k int) int {
	result, min, n := nums[k], nums[k], len(nums)
	for l, r := k, k; l > 0 || r < n-1; {
		if l == 0 {
			r++
			if nums[r] < min {
				min = nums[r]
			}
		} else if r == n-1 {
			l--
			if nums[l] < min {
				min = nums[l]
			}
		} else {
			if nums[l-1] > nums[r+1] {
				l--
				if nums[l] < min {
					min = nums[l]
				}
			} else {
				r++
				if nums[r] < min {
					min = nums[r]
				}
			}
		}
		tmp := min * (r - l + 1)
		if tmp > result {
			result = tmp
		}
	}
	return result
}

// @lc code=end
