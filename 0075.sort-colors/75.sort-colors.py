#
# @lc app=leetcode id=75 lang=python3
#
# [75] Sort Colors
#
from typing import List
# @lc code=start


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        cur, l, r = 0, 0, n-1
        while cur <= r:
            if nums[cur] == 0:
                if cur > l:
                    nums[l], nums[cur] = nums[cur], nums[l]
                else:
                    cur += 1
                l += 1
            elif nums[cur] == 2:
                if cur < r:
                    nums[r], nums[cur] = nums[cur], nums[r]
                    r -= 1
                else:
                    cur += 1
            else:
                cur += 1
# @lc code=end
