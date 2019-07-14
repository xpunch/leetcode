package main

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	var ml, mr, l, r int
	for r = n - 1; l < r; {
		for r > l {
			var c int
			match := false
			padding := 0
			if (r-l)%2 == 0 {
				c = (l + r + 1) / 2
			} else {
				c = (r + l) / 2
				padding = 1
			}
			for i := 1; c-i+padding >= l && c+i <= r; i++ {
				match = true
				if s[c-i+padding] != s[c+i] {
					match = false
					break
				}
			}
			if match && r-l > mr-ml {
				ml = l
				mr = r
				break
			}
			r--
			if r-l < mr-ml {
				break
			}
		}
		l++
		r = n - 1
		if r-l < mr-ml {
			break
		}
	}
	return s[ml : mr+1]
}
