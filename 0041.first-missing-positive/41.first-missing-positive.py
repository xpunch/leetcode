#
# @lc app=leetcode id=41 lang=python3
#
# [41] First Missing Positive
#

# @lc code=start


class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 0:
            return 1
        cache = [False for i in range(n)]
        for i in nums:
            if i >= 1 and i <= n:
                cache[i-1] = True
        for i in range(n):
            if not cache[i]:
                return i+1
        return n+1

# @lc code=end
