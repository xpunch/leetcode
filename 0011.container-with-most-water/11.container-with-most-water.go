package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	var left, right, max int
	right = len(height) - 1
	for left < right {
		var tmp int
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
			left++
		} else {
			tmp = height[right] * (right - left)
			right--
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
