#
# @lc app=leetcode id=35 lang=python3
#
# [35] Search Insert Position
#

# @lc code=start


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)-1
        while l <= r:
            if l == r or r-l == 1:
                if target <= nums[l]:
                    return l
                elif target <= nums[r]:
                    return r
                else:
                    return r+1
            m = (l+r+1)//2
            if target > nums[m]:
                l = m+1
            elif target < nums[m]:
                r = m-1
            else:
                return m
# @lc code=end
