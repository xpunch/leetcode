package xocde

import (
	"sort"
)

/*
 * @lc app=leetcode id=1363 lang=golang
 *
 * [1363] Largest Multiple of Three
 */

// @lc code=start
func largestMultipleOfThree(digits []int) string {
	if len(digits) == 0 {
		return ""
	}
	total := sum(digits)
	if total == 0 {
		return "0"
	}
	sort.Ints(digits)
	mod := total % 3
	if mod == 0 {
		return toString(digits)
	} else {
		var b1, b2 []int
		for _, d := range digits {
			if d%3 == 1 {
				b1 = append(b1, d)
			}
			if d%3 == 2 {
				b2 = append(b2, d)
			}
		}
		if mod == 1 {
			if len(b1) > 0 {
				return toString(digits, b1[0])
			}
			if len(b2) > 1 {
				return toString(digits, b2[0], b2[1])
			}
		} else if mod == 2 {
			if len(b2) > 0 {
				return toString(digits, b2[0])
			}
			if len(b1) > 1 {
				return toString(digits, b1[0], b1[1])
			}
		}
	}
	return ""
}

func sum(digits []int) int {
	var total int
	for _, d := range digits {
		total += d
	}
	return total
}

func toString(digits []int, excludes ...int) string {
	bytes := make([]byte, len(digits)-len(excludes))
	n, j := len(excludes), 0
	for i, d := range digits {
		if j < n && d == excludes[j] {
			j++
			continue
		}
		bytes[len(digits)-len(excludes)-1-i+j] = byte(d) + '0'
	}
	return string(bytes)
}

// @lc code=end
