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
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return permutate(n, k, 0, result, []int{})
}

func permutate(n int, k int, m int, result [][]int, nums []int) [][]int {
	if k == 0 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		result = append(result, tmp)
	} else {
		for i := m + 1; i <= n; i++ {
			newNums := nums[:]
			newNums = append(newNums, i)
			result = permutate(n, k-1, i, result, newNums)
		}
	}
	return result
}

// @lc code=end
