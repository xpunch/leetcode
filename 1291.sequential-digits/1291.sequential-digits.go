package xcode

/*
 * @lc app=leetcode id=1291 lang=golang
 *
 * [1291] Sequential Digits
 */

// @lc code=start
func sequentialDigits(low int, high int) []int {
	n1, n2 := getLength(low), getLength(high)
	starts := [10]int{0, 1, 12, 123, 1234, 12345, 123456, 1234567, 12345678, 123456789}
	steps := [10]int{0, 1, 11, 111, 1111, 11111, 111111, 1111111, 11111111, 111111111}
	ends := [10]int{0, 9, 89, 789, 6789, 56789, 456789, 3456789, 23456789, 123456789}
	var result []int
	for i := n1; i <= n2; i++ {
		end := ends[i]
		if low >= end {
			if low == end {
				result = append(result, end)
			}
			continue
		}
		for start, step, j := starts[i], steps[i], 0; j <= 9-i && start <= high; j++ {
			if start >= low {
				result = append(result, start)
			}
			start += step
		}
	}
	return result
}

func getLength(n int) int {
	if n == 0 {
		return 0
	}
	if n >= 1000000000 {
		return 9
	}
	len := 1
	for i := n; i >= 10; i = i / 10 {
		len++
	}
	return len
}

// @lc code=end
