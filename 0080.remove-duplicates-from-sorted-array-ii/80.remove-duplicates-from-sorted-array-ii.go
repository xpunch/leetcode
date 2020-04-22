package main

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	duplicates, removed := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			duplicates++
		} else {
			duplicates = 0
		}
		if duplicates > 1 {
			removed++
		}
		if removed > 0 {
			nums[i-removed] = nums[i]
		}
	}
	return len(nums) - removed
}

// @lc code=end
