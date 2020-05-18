package main

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	n := len(heights)
	var result int
	for i := 0; i < n; i++ {
		height := heights[i]
		if height == 0 || (i > 0 && height <= heights[i-1]) {
			continue
		}
		area := height
		for j := i + 1; j < n; j++ {
			right := heights[j]
			if right == 0 {
				break
			}
			if right < height {
				height = right
			}
			tmp := height * (j - i + 1)
			if tmp > area {
				area = tmp
			}
		}
		if area > result {
			result = area
		}
	}
	return result
}

// @lc code=end
