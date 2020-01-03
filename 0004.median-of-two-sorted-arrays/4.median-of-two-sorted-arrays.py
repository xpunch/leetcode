#
# @lc app=leetcode id=4 lang=python3
#
# [4] Median of Two Sorted Arrays
#

# @lc code=start


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        total = len(nums1) + len(nums2)
        if total % 2 != 0:
            mid = int(total // 2)
            s1, s2 = 0, 0
            while True:
                if s1 == len(nums1):
                    if s1 + s2 == mid:
                        return nums2[s2]
                    s2 += 1
                elif s2 == len(nums2):
                    if s1 + s2 == mid:
                        return nums1[s1]
                    s1 += 1
                elif nums1[s1] < nums2[s2]:
                    if s1 + s2 == mid:
                        return nums1[s1]
                    s1 += 1
                else:
                    if s1 + s2 == mid:
                        return nums2[s2]
                    s2 += 1
        else:
            mid = int(total//2)
            s1, s2, m1 = 0, 0, 0
            while True:
                if s1 == len(nums1):
                    if s1 + s2 == mid - 1:
                        m1 = nums2[s2]
                    elif s1 + s2 == mid:
                        return (m1 + nums2[s2])/2
                    s2 += 1
                elif s2 == len(nums2):
                    if s1 + s2 == mid - 1:
                        m1 = nums1[s1]
                    elif s1 + s2 == mid:
                        return (m1 + nums1[s1])/2
                    s1 += 1
                elif nums1[s1] < nums2[s2]:
                    if s1 + s2 == mid - 1:
                        m1 = nums1[s1]
                    elif s1 + s2 == mid:
                        return (m1 + nums1[s1])/2
                    s1 += 1
                else:
                    if s1 + s2 == mid - 1:
                        m1 = nums2[s2]
                    elif s1 + s2 == mid:
                        return (m1 + nums2[s2])/2
                    s2 += 1
        return 0
# @lc code=end
