package main

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	buf := make([]byte, 0)
	for x > 0 {
		buf = append(buf, byte(x%10))
		x = x / 10
	}
	for i := 0; i < len(buf)/2; i++ {
		if buf[i] != buf[len(buf)-i-1] {
			return false
		}
	}
	return true
}
