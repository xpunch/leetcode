/*
 * @lc app=leetcode id=1556 lang=golang
 *
 * [1556] Thousand Separator
 */
package xcode

import "fmt"

// @lc code=start
func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	var result string
	for i := n; i > 0; {
		m := i % 1000
		i = i / 1000
		if i == 0 {
			result = fmt.Sprintf("%d.", m) + result
		} else {
			result = fmt.Sprintf("%03d.", m) + result
		}
	}
	return result[:len(result)-1]
}

// @lc code=end
