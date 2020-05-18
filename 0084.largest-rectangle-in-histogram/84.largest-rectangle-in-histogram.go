package main

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	n, stack := len(heights), make([]int, 0)
	var result int
	for i := 0; i <= n; i++ {
		for len(stack) > 0 && (i == n || heights[stack[len(stack)-1]] >= heights[i]) {
			height, width := heights[stack[len(stack)-1]], i
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				// height may be duplicated
				width = i - (stack[len(stack)-1] + 1)
			}
			area := height * width
			if area > result {
				result = area
			}
		}
		stack = append(stack, i)
	}
	return result
}

// @lc code=end
