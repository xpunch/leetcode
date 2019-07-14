package main

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i] += nums[j]
				nums[j] = nums[i] - nums[j]
				nums[i] -= nums[j]
			}
		}
	}
	count := len(nums)
	result := make([][]int, 0)
	for i := 0; i < count-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := count - 1; j > i; j-- {
			if j != count-1 && nums[j] == nums[j+1] {
				continue
			}
			for l, k, r := i, (i+j+1)/2, j; k < r && k > l; {
				// if k != count-2 && nums[l] == nums[k] {
				// 	l = k
				// 	k = (r + k + 1) / 2
				// 	continue
				// }
				// if k != r-1 && nums[r] == nums[k] {
				// 	r = k
				// 	k = (l + k + 1) / 2
				// 	continue
				// }
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					result = append(result, []int{nums[i], nums[k], nums[j]})
					break
				} else if sum > 0 {
					r = k
					k = (l + k + 1) / 2
				} else if sum < 0 {
					l = k
					k = (r + k + 1) / 2
				}
			}
		}
	}
	return result
}
