#
# @lc app=leetcode id=60 lang=python3
#
# [60] Permutation Sequence
#

# @lc code=start


class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        if n == 1:
            return "1"
        nums = [str(i+1) for i in range(n)]
        result, factorial = '', 1
        for i in range(2, n):
            factorial *= i
        k = k-1
        for i in range(n):
            if i == n-1:
                result += nums[0]
            else:
                ki = k // factorial
                k = k % factorial
                factorial = factorial//(n-i-1)
                result += nums[ki]
                del nums[ki]
        return result
# @lc code=end
