/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; ; i++ {
		for j, s := range strs {
			if i == len(s) {
				return s
			}
			if j != 0 && strs[j-1][i] != s[i] {
				return s[0:i]
			}
		}
	}
}
