/*
 * @lc app=leetcode id=1374 lang=golang
 *
 * [1374] Generate a String With Characters That Have Odd Counts
 */

// @lc code=start
func generateTheString(n int) string {
	result := make([]byte, 0, n)
	alphabets := []byte("abcdefghijklmnopqrstuvwxyz")
	cache := make(map[byte]int, 0)
	for _, v := range alphabets {
		cache[v] = 0
	}
	usedCharCount := len(alphabets)
	if n%2 != 0 {
		usedCharCount--
	}
	for i, j := 0, 0; i < n; {
		ch := alphabets[j]
		chCount := cache[ch]
		if chCount%2 == 0 {
			result = append(result, ch)
			i++
			cache[ch] = chCount + 1
		} else {
			result = append(result, ch, ch)
			i += 2
			cache[ch] = chCount + 2
		}
		j = (j + 1) % usedCharCount
	}
	return string(result)
}

// @lc code=end
