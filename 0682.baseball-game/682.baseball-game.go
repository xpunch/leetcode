/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */
package xcode

import (
	"strconv"
)

// @lc code=start
func calPoints(ops []string) int {
	points, ans := make([]int, 0, len(ops)), 0
	for _, p := range ops {
		switch p {
		case "C":
			ans -= points[len(points)-1]
			points = points[:len(points)-1]
		case "D":
			score := points[len(points)-1] * 2
			ans += score
			points = append(points, score)
		case "+":
			score := points[len(points)-1] + points[len(points)-2]
			ans += score
			points = append(points, score)
		default:
			score, _ := strconv.Atoi(p)
			ans += score
			points = append(points, score)
		}
	}
	return ans
}

// @lc code=end
