package main

/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	result, heights := 0, make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 48 {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		stack := make([]int, 0)
		for k := 0; k <= n; k++ {
			for len(stack) > 0 && (k == n || heights[k] <= heights[stack[len(stack)-1]]) {
				height, width := heights[stack[len(stack)-1]], k
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					width = k - (stack[len(stack)-1] + 1)
				}
				area := height * width
				if area > result {
					result = area
				}
			}
			stack = append(stack, k)
		}
	}
	return result
}

// @lc code=end
