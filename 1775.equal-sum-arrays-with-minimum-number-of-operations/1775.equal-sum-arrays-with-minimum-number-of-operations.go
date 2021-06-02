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
	cache1, cache2 := make([]int, 6), make([]int, 6)
	sum1, sum2 := 0, 0
	for i := 0; i < len1; i++ {
		sum1 += nums1[i]
		cache1[nums1[i]-1]++
	}
	for i := 0; i < len2; i++ {
		sum2 += nums2[i]
		cache2[nums2[i]-1]++
	}
	if sum1 == sum2 {
		return 0
	}
	var count int
	if sum1 > sum2 {
		// cache1 for sub: c1
		// cache2 for add: 5-c1
		c1, c2 := 5, 0
		for difference := sum1 - sum2; difference > 0; {
			if c1 >= 5-c2 {
				if cache1[c1] == 0 {
					c1--
					continue
				}
				difference -= c1
				cache1[c1]--
			} else {
				if cache2[c2] == 0 {
					c2++
					continue
				}
				difference -= 5 - c2
				cache2[c2]--
			}
			count++
		}
	} else {
		// c1 for add: 5-c1
		// c2 for sub: c2
		c1, c2 := 0, 5
		for difference := sum2 - sum1; difference > 0; {
			if 5-c1 >= c2 {
				if cache1[c1] == 0 {
					c1++
					continue
				}
				difference -= 5 - c1
				cache1[c1]--
			} else {
				if cache2[c2] == 0 {
					c2--
					continue
				}
				difference -= c2
				cache2[c2]--
			}
			count++
		}
	}
	return count
}

// @lc code=end
