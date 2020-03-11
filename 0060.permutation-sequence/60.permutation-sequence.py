#
# @lc app=leetcode id=60 lang=python3
#
# [60] Permutation Sequence
#

# @lc code=start
# 0 0 0
# 0 1 0
# 1 0 0
# 1 1 0
# 2 0 0
# 2 1 0


class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        if n == 1:
            return "1"
        result, nums = "", [str(i+1) for i in range(n)]
        cache, cursor, i = [0 for i in range(n)], n-2, 0
        while i < k-1:
            if cache[cursor] < n-cursor-1:
                i += 1
                cache[cursor] += 1
                if cursor < n-2:
                    cursor = n-2
            else:
                cache[cursor] = 0
                cursor -= 1
        for i in cache:
            result += nums[i]
            del nums[i]
        return result
# @lc code=end
