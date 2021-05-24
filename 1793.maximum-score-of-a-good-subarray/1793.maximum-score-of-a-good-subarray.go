package xcode

/*
 * @lc app=leetcode id=1793 lang=golang
 *
 * [1793] Maximum Score of a Good Subarray
 */

// @lc code=start
func maximumScore(nums []int, k int) int {
	n := len(nums)
	left, right, stack := make([]int, n), make([]int, n), make([]int, 0)
	for i := 0; i < n; i++ {
		left[i] = -1
		for len(stack) > 0 {
			e := len(stack) - 1
			if nums[stack[e]] < nums[i] {
				break
			}
			left[i] = left[stack[e]]
			stack = stack[:e]
		}
		if left[i] == -1 {
			left[i] = i
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		right[i] = -1
		for len(stack) > 0 {
			e := len(stack) - 1
			if nums[stack[e]] < nums[i] {
				break
			}
			right[i] = right[stack[e]]
			stack = stack[:e]
		}
		if right[i] == -1 {
			right[i] = i
		}
		stack = append(stack, i)
	}
	var result int
	for i := 0; i < n; i++ {
		if left[i] <= k && right[i] >= k {
			tmp := nums[i] * (right[i] - left[i] + 1)
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

// @lc code=end
