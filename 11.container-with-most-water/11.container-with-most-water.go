package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			tmp := (j - i) * min(height[i], height[j])
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(s int, t int) int {
	if s < t {
		return s
	}
	return t
}
