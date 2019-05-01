/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	cache := make(map[byte]int)
	var result, longest int
	var i, j int
	for j < len(s) {
		if x, ok := cache[s[j]]; ok {
			if i <= x {
				longest = j - i
				if longest > result {
					result = longest
				}
				i = x + 1
				continue
			}
		}
		cache[s[j]] = j
		j++
	}
	longest = j - i
	if longest > result {
		result = longest
	}
	return result
}
