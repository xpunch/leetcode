package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}, nums}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := 1; j < len(nums); j++ {
			result = append(result, dsf(nums, []int{nums[i]}, i+1, j)...)
		}
	}
	return result
}

func dsf(nums []int, set []int, start int, size int) [][]int {
	result := [][]int{}
	if len(set) == size {
		return append(result, set)
	}
	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		tmp := append([]int(nil), set...)
		tmp = append(tmp, nums[i])
		result = append(result, dsf(nums, tmp, i+1, size)...)
	}
	return result
}

// @lc code=end
