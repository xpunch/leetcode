#
# @lc app=leetcode id=55 lang=python3
#
# [55] Jump Game
#
from typing import List
# @lc code=start


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        n = len(nums)
        if n <= 1:
            return True
        if nums[0] == 0:
            return False
        i, furthest = 1, nums[0]
        while i < n and i <= furthest:
            further = i+nums[i]
            if further >= n-1:
                return True
            if further > furthest:
                furthest = further
            i += 1
        return False
# @lc code=end
