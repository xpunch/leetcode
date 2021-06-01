/*
 * @lc app=leetcode id=466 lang=golang
 *
 * [466] Count The Repetitions
 */
package xcode

// @lc code=start
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	// 1. if s2 has x repetitions in [s1,n1], means [s2,n2] has x/n2 repetitions in [s1,n1]
	// 2. repetition index
	// 3. next index
	// s2 repetition count in target s1 repetition segment
	repetitionCount := make([]int, n1+1)
	// first s2 char index in next s1 repetition segment
	nextIndex := make([]int, n1+1)
	for seg, i2, count := 1, 0, 0; seg <= n1; seg++ {
		for i1 := 0; i1 < len(s1); i1++ {
			if s1[i1] == s2[i2] {
				i2++
			}
			if i2 == len(s2) {
				i2 = 0
				count++
			}
		}
		repetitionCount[seg] = count
		nextIndex[seg] = i2
		for i := 0; i < seg; i++ {
			if nextIndex[i] == i2 {
				// s1 repetition count for a loop pattern
				interval := seg - i
				// repeat loop of pattern
				repeat := (n1 - i) / interval
				patternCount := (repetitionCount[seg] - repetitionCount[i]) * repeat
				// after remove patterns in [s1,n1], repetition count in remained string
				remainCount := repetitionCount[i+(n1-i)%interval]
				return (patternCount + remainCount) / n2
			}
		}
	}
	return repetitionCount[n1] / n2
}

// @lc code=end
