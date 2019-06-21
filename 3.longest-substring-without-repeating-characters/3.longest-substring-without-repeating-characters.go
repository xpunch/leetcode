/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	cache := make(map[byte]int)
	var result, longest int
	var i int
	for j, c := range []byte(s {
		if x, ok := cache[c]; ok && x >= i {
			i = x + 1
		}
		longest = j - i + 1
		if longest > result {
			result = longest
		}
		cache[c] = j
	}
	return result
}
