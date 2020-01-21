#
# @lc app=leetcode id=33 lang=python3
#
# [33] Search in Rotated Sorted Array
#

# @lc code=start


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)
        if n == 0:
            return -1
        result = -1
        l, r = 0, n-1
        while l <= r and l >= 0 and r < n:
            if r-l == 0:
                return 0 if nums[l] == target else -1
            if target < nums[l] and target > nums[r]:
                return -1
            if nums[l] == target:
                return l
            if nums[r] == target:
                return r
            if r-l == 1:
                return -1
            m = (l+r+1)//2
            if nums[m] == target:
                return m
            if nums[m] > nums[l]:
                if target > nums[l]:
                    if target > nums[m]:
                        l = m+1
                    else:
                        r = m-1
                else:
                    l = m+1
            else:
                if target > nums[l]:
                    r = m-1
                else:
                    if target > nums[m]:
                        l = m+1
                    else:
                        r = m-1
        return result
# @lc code=end
