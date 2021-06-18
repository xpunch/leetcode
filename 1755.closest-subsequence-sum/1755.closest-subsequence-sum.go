package xcode

import "sort"

/*
 * @lc app=leetcode id=1755 lang=golang
 *
 * [1755] Closest Subsequence Sum
 */

// @lc code=start
func minAbsDifference(nums []int, goal int) int {
	if goal == 0 {
		return 0
	}
	var sum int
	for _, i := range nums {
		sum += i
	}
	if sum == goal {
		return 0
	}
	var ans int = abs(sum - goal)
	n := len(nums)
	m := n >> 1

	h := make([]int, 0, 1<<m)
	for s := 0; s < (1 << m); s++ {
		t := 0
		for i := 0; i < m; i++ {
			if s&(1<<i) != 0 {
				t += nums[i]
			}
		}
		h = append(h, t)
	}

	sort.Ints(h)

	for s := 0; s < (1 << (n - m)); s++ {
		t := goal
		for i := m; i < n; i++ {
			if s&(1<<(i-m)) != 0 {
				t -= nums[i]
			}
		}
		l, r := 0, len(h)-1
		for l < r {
			mid := (l + r) >> 1
			if h[mid] < t {
				l = mid + 1
			} else {
				r = mid
			}
		}

		if ans > abs(h[l]-t) {
			ans = abs(h[l] - t)
		}

		if l >= 1 && ans > abs(h[l-1]-t) {
			ans = abs(h[l-1] - t)
		}
	}
	return ans
}

func abs(v int) int {
	if v >= 0 {
		return v
	} else {
		return -v
	}
}

// @lc code=end
