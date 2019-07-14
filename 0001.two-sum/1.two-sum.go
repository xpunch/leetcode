/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		last := target - nums[i]
		if j, ok := cache[last]; ok {
			return []int{i, j}
		}
		cache[nums[i]] = i
	}
	return nil
}
