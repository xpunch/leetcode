package main

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < m || len(nums2) < n || len(nums1) < m+n {
		return
	}
	m2, n2 := m-1, n-1
	for m2 >= -1 && n2 >= 0 {
		if m2 == -1 {
			nums1[n2] = nums2[n2]
			n2--
			continue
		}
		if nums1[m2] >= nums2[n2] {
			nums1[m2+n2+1] = nums1[m2]
			m2--
		} else {
			nums1[m2+n2+1] = nums2[n2]
			n2--
		}
	}
}

// @lc code=end
