package main

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	if k > n {
		return nil
	}
	result := make([][]int, 0)
	permutate(n, k, 0, &result, []int{})
	return result
}

func permutate(n int, k int, m int, result *[][]int, nums []int) {
	if k == 0 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
	} else {
		for i := m + 1; i <= n; i++ {
			nums = append(nums, i)
			permutate(n, k-1, i, result, nums)
			nums = nums[:len(nums)-1]
		}
	}
}

// @lc code=end
