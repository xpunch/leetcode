/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */
package xcode

// @lc code=start
func findWords(words []string) []string {
	chars := [26]rune{2, 3, 3, 2, 1, 2, 2, 2, 1, 2, 2, 2, 3, 3, 1, 1, 1, 1, 2, 1, 1, 3, 1, 3, 1, 3}
	ans := make([]string, 0, len(words))
	for _, w := range words {
		var line rune
		var match bool = true
		for i, c := range w {
			var tmp rune
			if c >= 97 {
				tmp = c - 97
			} else {
				tmp = c - 65
			}
			l := chars[tmp]
			if i == 0 {
				line = l
			} else if l != line {
				match = false
				break
			}
		}
		if match {
			ans = append(ans, w)
		}
	}
	return ans
}

// @lc code=end
