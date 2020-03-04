#
# @lc app=leetcode id=46 lang=python3
#
# [46] Permutations
#

# @lc code=start

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        if n == 1:
            return [nums]
        index, num, used, result = 0, [0 for i in range(n)], [
            False for i in range(n)], []
        self.generate(nums, index, num, used, result)
        return result

    def generate(self, nums: List[int], index: int, num: List[int], used: List[bool], result: List[List[int]]):
        if index == len(nums):
            tmp = num[:]
            result.append(tmp)
        else:
            for i in range(len(nums)):
                if not used[i]:
                    used[i] = True
                    num[index] = nums[i]
                    self.generate(nums, index+1, num, used, result)
                    used[i] = False

# @lc code=end
