package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {
		return nil
	}
	result := make([]string, 0)
	for i := 1; i <= min(3, len(s)-3); i++ {
		seg := s[0:i]
		if valid(seg) {
			result = append(result, dsf(s, []string{seg}, i)...)
		}
	}
	return result
}

func dsf(s string, segs []string, start int) []string {
	if start == len(s) && len(segs) == 4 {
		return []string{build(segs)}
	}
	result := make([]string, 0)
	if len(s)-start > 3*(4-len(segs)) || len(s)-start < (4-len(segs)) {
		return result
	}
	for i := 1; i <= min(3, len(s)-start-3+len(segs)); i++ {
		seg := s[start : start+i]
		if valid(seg) {
			tmp := append([]string(nil), segs...)
			result = append(result, dsf(s, append(tmp, seg), start+i)...)
		}
	}
	return result
}

func min(t, s int) int {
	if t < s {
		return t
	}
	return s
}

func build(segs []string) string {
	return strings.Join(segs, ".")
}

func valid(s string) bool {
	switch len(s) {
	case 1:
		return true
	case 2:
		return s[0] != '0'
	case 3:
		switch s[0] {
		case '1':
			return true
		case '2':
			if s[1] < '5' {
				return true
			}
			if s[1] == '5' {
				return s[2] < '6'
			}
			return false
		}
		return false
	}
	return false
}

// @lc code=end
