/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	var cache map[byte]interface{}
	var result int
	for i := 0; i < len(s); i++ {
		cache = make(map[byte]interface{})
		var j = i
		for ; j < len(s); j++ {
			if _, ok := cache[s[j]]; ok {
				break
			} else {
				cache[s[j]] = struct{}{}
			}
		}
		longest := j - i
		if longest > result {
			result = longest
		}
	}
	return result
}
