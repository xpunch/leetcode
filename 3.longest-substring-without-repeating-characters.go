/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	cache := make(map[byte]byte)
	var result, longest int
	var i int
	j := i
	for ; i < len(s); i++ {
		if result > 0 {
			j = longest + i - 1
			delete(cache, s[i-1])
		}
		for ; j < len(s); j++ {
			if _, ok := cache[s[j]]; ok {
				break
			} else {
				cache[s[j]] = s[j]
			}
		}
		longest = j - i
		if longest > result {
			result = longest
		}
	}
	return result
}
