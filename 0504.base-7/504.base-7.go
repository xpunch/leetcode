package xcode

import "fmt"

/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	result, negative := "", false
	if num < 0 {
		negative = true
		num = num * -1
	}
	for num > 0 {
		result = fmt.Sprint(num%7) + result
		num /= 7
	}
	if negative {
		result = "-" + result
	}
	return result
}

// @lc code=end
