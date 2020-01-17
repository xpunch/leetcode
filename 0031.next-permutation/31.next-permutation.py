#
# @lc app=leetcode id=31 lang=python3
#
# [31] Next Permutation
#

# @lc code=start
from typing import List


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        n = len(nums)
        if n > 1:
            self.arrange(nums, 0, n)

    def arrange(self, nums: List[int], l, n: int) -> None:
        if l == n-1:
            return
        tp = l
        for i in range(l, n-1):
            if nums[i+1] >= nums[i]:
                tp = i+1
            else:
                break
        if tp == n-1:
            for i in range(n-1, l, -1):
                if nums[i] != nums[i-1]:
                    nums[i], nums[i-1] = nums[i-1], nums[i]
                    break
        else:
            isDesc = True
            for i in range(tp+1, n-1):
                if nums[i+1] > nums[i]:
                    isDesc = False
                    break
            if isDesc:
                if tp == 0:
                    for i in range((n)//2):
                        nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
                else:
                    p = tp
                    for i in range(tp, l, -1):
                        if nums[i] >= nums[i+1]:
                            p -= 1
                        else:
                            break
                    for i in range(n-1, p, -1):
                        if nums[i] > nums[p]:
                            nums[p], nums[i] = nums[i], nums[p]
                            break
                    for i in range((n-p-1)//2):
                        nums[p+1+i], nums[n-1-i] = nums[n-1-i], nums[p+1+i]
            else:
                self.arrange(nums, tp+1, n)

# @lc code=end
