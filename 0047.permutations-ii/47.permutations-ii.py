#
# @lc app=leetcode id=47 lang=python3
#
# [47] Permutations II
#
from typing import List

# @lc code=start


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        if n == 1:
            return [nums]
        nums.sort()
        index, num, used, result = 0, [0 for i in range(n)], [False for i in range(n)], []
        self.generate(nums, index, num, used, result)
        return result

    def generate(self, nums: List[int], index: int, num: List[int], used: List[bool], result: List[List[int]]):
        if index == len(nums):
            tmp = num[:]
            result.append(tmp)
        else:
            for i in range(len(nums)):
                if not used[i]:
                    if i > 0 and nums[i] == nums[i-1] and not used[i-1]:
                        continue
                    used[i] = True
                    num[index] = nums[i]
                    self.generate(nums, index+1, num, used, result)
                    used[i] = False
# @lc code=end
