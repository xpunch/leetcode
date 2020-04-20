package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result, n := [][]int{{}}, len(nums)
	for i := 1; i <= n; i++ {
		permutate(n, i, -1, nums, []int{}, &result)
	}
	return result
}

func permutate(n, k, m int, nums, current []int, result *[][]int) {
	if k == 0 {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
	} else {
		for i := m + 1; i < n; i++ {
			current = append(current, nums[i])
			permutate(n, k-1, i, nums, current, result)
			current = current[:len(current)-1]
		}
	}
}

// @lc code=end
