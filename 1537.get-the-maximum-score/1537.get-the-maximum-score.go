/*
 * @lc app=leetcode id=1537 lang=golang
 *
 * [1537] Get the Maximum Score
 */
package xcode

// @lc code=start
func maxSum(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	var ans int
	i, j, sum1, sum2 := 0, 0, 0, 0
	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			if sum1 > sum2 {
				ans += sum1 + nums1[i]
			} else {
				ans += sum2 + nums2[j]
			}
			sum1 = 0
			sum2 = 0
			i++
			j++
		} else if nums1[i] < nums2[j] {
			sum1 += nums1[i]
			i++
		} else {
			sum2 += nums2[j]
			j++
		}
	}
	if i == n1 || j == n2 {
		if i == n1 {
			for k := j; k < n2; k++ {
				sum2 += nums2[k]
			}
		}
		if j == n2 {
			for k := i; k < n1; k++ {
				sum1 += nums1[k]
			}
		}
		if sum1 > sum2 {
			ans += sum1
		} else {
			ans += sum2
		}
	}
	return ans % 1000000007
}

// @lc code=end
