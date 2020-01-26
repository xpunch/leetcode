#
# @lc app=leetcode id=34 lang=python3
#
# [34] Find First and Last Position of Element in Sorted Array
#

# @lc code=start


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        n, result = len(nums), [-1, -1]
        if n == 0:
            return result
        if n == 1:
            return [0, 0] if nums[0] == target else result
        l, r = 0, n-1
        while l <= r and l >= 0 and r < n:
            m = (l+r+1)//2
            if target > nums[m]:
                l = m+1
            elif target < nums[m]:
                r = m-1
            else:
                left = self.findFirstIndex(nums, l, m, target)
                right = self.findLastIndex(nums, m, r, target)
                return [left, right]
        return result

    def findFirstIndex(self, nums: List[int], l, r, target: int) -> int:
        while l <= r:
            if r == l:
                return l
            if r-l == 1:
                return l if nums[l] == target else r
            m = (l+r+1)//2
            if target > nums[m]:
                l = m+1
            elif target < nums[m]:
                r = m-1
            else:
                r = m
        return l

    def findLastIndex(self, nums: List[int], l, r, target: int) -> int:
        while l <= r:
            if r == l:
                return r
            if r-l == 1:
                return r if nums[r] == target else l
            m = (l+r)//2
            if target > nums[m]:
                l = m+1
            elif target < nums[m]:
                r = m-1
            else:
                l = m
        return r
# @lc code=end
