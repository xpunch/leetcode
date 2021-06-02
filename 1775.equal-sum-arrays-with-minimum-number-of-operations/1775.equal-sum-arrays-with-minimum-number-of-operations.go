/*
 * @lc app=leetcode id=1775 lang=golang
 *
 * [1775] Equal Sum Arrays With Minimum Number of Operations
 */
package xcode

// @lc code=start
func minOperations(nums1 []int, nums2 []int) int {
	len1, len2 := len(nums1), len(nums2)
	if len1 > 6*len2 || 6*len1 < len2 {
		return -1
	}
	sum1, sum2 := 0, 0
	for i := 0; i < len1; i++ {
		sum1 += nums1[i]
	}
	for i := 0; i < len2; i++ {
		sum2 += nums2[i]
	}
	if sum1 == sum2 {
		return 0
	}
	var difference int
	var count int
	var steps [5]int
	if sum1 > sum2 {
		for _, i := range nums1 {
			if i == 1 {
				continue
			}
			steps[i-2]++
		}
		for _, i := range nums2 {
			if i == 6 {
				continue
			}
			steps[5-i]++
		}
		difference = sum1 - sum2
	} else {
		for _, i := range nums1 {
			if i == 6 {
				continue
			}
			steps[5-i]++
		}
		for _, i := range nums2 {
			if i == 1 {
				continue
			}
			steps[i-2]++
		}
		difference = sum2 - sum1
	}
	for i := 5; i >= 1; i-- {
		max := steps[i-1] * i
		if max > difference {
			if difference%i == 0 {
				return count + difference/i
			}
			return count + difference/i + 1
		}
		count += steps[i-1]
		difference -= max
	}
	return count
}

// @lc code=end
