package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [15] 3Sum Closest
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i <= n-3; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum >= target {
				r--
			} else {
				l++
			}
		}
	}
	return closest
}

func abs(target int) int {
	if target < 0 {
		return -target
	}
	return target
}
