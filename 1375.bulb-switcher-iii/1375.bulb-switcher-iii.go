/*
 * @lc app=leetcode id=1375 lang=golang
 *
 * [1375] Bulb Switcher III
 */

// @lc code=start
func numTimesAllBlue(light []int) int {
	result := 0
	lastBlue, maxLight := 0, 0
	cache := make(map[int]bool, 0)
	for i := 0; i < len(light); i++ {
		if light[i] > maxLight {
			maxLight = light[i]
		}
		cache[light[i]] = true
		if lastBlue+1 == light[i] {
			lastBlue = light[i]
		}
		if lastBlue != light[i] {
			continue
		}
		allBlue := true
		for i := lastBlue + 1; i <= maxLight; i++ {
			if b, ok := cache[i]; !ok || !b {
				allBlue = false
				break
			}
			lastBlue = i
		}
		if allBlue {
			result++
		}
	}
	return result
}

// @lc code=end
