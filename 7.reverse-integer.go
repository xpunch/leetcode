package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	source := strconv.Itoa(x)
	start := 0
	if source[0] == '-' {
		start = 1
	}
	target := revertStr([]byte(source), start)
	result, _ := strconv.Atoi(target)
	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}

func revertStr(str []byte, start int) string {
	for from, to := start, len(str)-1; from < to; from, to = from+1, to-1 {
		str[from], str[to] = str[to], str[from]
	}
	return string(str)
}
