package main

/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	l, r := 0, n-1
	for l <= r {
		mid := (r + l) >> 1
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] == nums[mid] {
				l++
			}
			if nums[r] == nums[mid] {
				r--
			}
		}
	}
	return false
}

// @lc code=end
