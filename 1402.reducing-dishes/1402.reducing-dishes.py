#
# @lc app=leetcode id=1402 lang=python3
#
# [1402] Reducing Dishes
#
from typing import List
# @lc code=start


class Solution:
    def maxSatisfaction(self, satisfaction: List[int]) -> int:
        satisfaction.sort()
        total, cur, result,  n = 0, 0, 0,  len(satisfaction)
        for i in range(n):
            total += satisfaction[i]
        while total < 0 and cur < n:
            total -= satisfaction[cur]
            cur += 1
        print(satisfaction[cur:])
        for i in range(cur, n):
            result += (i-cur+1)*satisfaction[i]
        return result
# @lc code=end
