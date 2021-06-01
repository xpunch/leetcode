package xcode

import "sort"

/*
 * @lc app=leetcode id=1291 lang=golang
 *
 * [1291] Sequential Digits
 */

// @lc code=start
func sequentialDigits(low int, high int) []int {
	var result []int
	for i := 1; i <= 9; i++ {
		for x := i; x <= high; {
			r := x % 10
			if r == 0 {
				break
			}
			if x >= low {
				result = append(result, x)
			}
			x = x*10 + r + 1
		}
	}
	sort.Ints(result)
	return result
}

// @lc code=end
