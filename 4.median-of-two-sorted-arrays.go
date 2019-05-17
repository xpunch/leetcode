import (
	"math"
)

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m1 := len(nums1)
	m2 := len(nums2)
	if m1 > m2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	if m1 == 0 {
		return (float64(nums2[(m2-1)/2]) + float64(nums2[m2/2])) / 2
	}
	var c1, c2 int
	l1 := nums1[0]
	l2 := nums2[0]
	if m1 == 1 {
		if m2 == 1 {
			return (float64(l1) + float64(l2)) / 2
		}
		c2 = m2/2 - 1
		l2 = nums2[c2]
		r2 := nums2[c2+1]
		if m2%2 == 0 {
			if l1 <= l2 {
				return float64(l2)
			} else if l1 <= r2 {
				return float64(l1)
			} else {
				return float64(r2)
			}
		} else {
			if l1 <= l2 {
				return float64(l2+r2) / 2
			} else if l1 <= nums2[c2+2] {
				return float64(l1+r2) / 2
			} else {
				return (float64(r2) + float64(nums2[c2+2])) / 2
			}
		}
	}
	r1 := nums1[m1-1]
	r2 := nums2[m2-1]
	if l1 >= l2 && l1 >= r2 {
		if m1 == m2 {
			return (float64(l1) + float64(r2)) / 2
		}
		return (float64(nums2[(m1+m2-1)/2]) + float64(nums2[(m1+m2)/2])) / 2
	}
	if r1 <= l2 && r1 <= r2 {
		if m1 == m2 {
			return (float64(r1) + float64(l2)) / 2
		}
		return (float64(nums2[(m1+m2-1)/2-m1]) + float64(nums2[(m1+m2)/2-m1])) / 2
	}
	lc1 := 0
	lc2 := m1 - 1
	c1 = (lc1 + lc2) / 2
	l1 = nums1[c1]
	r1 = nums1[c1+1]
	c2 = (m1+m2)/2 - c1 - 2
	l2 = nums2[c2]
	r2 = nums2[c2+1]
	for {
		if l1 <= r2 && l2 <= r1 {
			break
		}
		if l1+r1 > l2+r2 {
			if c1 != 0 {
				lc2 = c1
				c1 = (lc1 + lc2) / 2
				c2 = (m1+m2)/2 - c1 - 2
			} else {
				if c2 == m2-2 {
					break
				}
				c2 = c2 + 1
			}
		} else {
			if c1 != m1-2 {
				lc1 = c1
				c1 = (lc1 + lc2) / 2
				c2 = (m1+m2)/2 - c1 - 2
			} else {
				if c2 == 0 {
					break
				}
				c2 = c2 - 1
			}
		}
		l1 = nums1[c1]
		if m1 != c1+1 {
			r1 = nums1[c1+1]
		}
		l2 = nums2[c2]
		if m2 != c2+1 {
			r2 = nums2[c2+1]
		}
	}
	if (m1+m2)%2 == 0 {
		if c2 < (m1+m2)/2-c1-4 {
			return (float64(nums2[m2/2-2]) + float64(nums2[m2/2-1])) / 2
		}
		if c2 == (m1+m2)/2-c1-4 {
			return (float64(nums2[c2+2]) + float64(r2)) / 2
		}
		if c2 == (m1+m2)/2-c1-3 {
			return (float64(r1) + float64(r2)) / 2
		}
		if c2 == (m1+m2)/2-c1-1 {
			return (float64(l1) + float64(l2)) / 2
		}
		if c2 == (m1+m2)/2-c1 {
			return (float64(l2) + float64(nums2[c2-1])) / 2
		}
		if c2 > (m1+m2)/2-c1 {
			return (float64(nums2[(m2+1)/2]) + float64(nums2[(m2+1)/2+1])) / 2
		}
		return (math.Min(float64(r1), float64(r2)) + math.Max(float64(l1), float64(l2))) / 2
	} else {
		if c2 < (m1+m2)/2-c1-3 {
			return float64(nums2[m2/2-1-c1])
		}
		if c2 == (m1+m2)/2-c1-3 {
			return float64(r2)
		}
		if c2 == (m1+m2)/2-c1-1 {
			return float64(l1)
		}
		if c2 == (m1+m2)/2-c1 {
			return float64(l2)
		}
		if c2 > (m1+m2)/2-c1 {
			return float64(nums2[m2/2+1])
		}
		if l1 > r2 {
			return float64(l1)
		}
		if r1 < l2 {
			return float64(l2)
		}
		if r1 < r2 {
			return float64(r1)
		}
		return float64(r2)
	}
}
