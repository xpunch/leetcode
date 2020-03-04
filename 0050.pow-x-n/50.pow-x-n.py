#
# @lc app=leetcode id=50 lang=python3
#
# [50] Pow(x, n)
#

# @lc code=start


class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0:
            return 1
        if n == 1:
            return x
        if x == 0:
            return x
        y, m = (x if n > 0 else 1/x), abs(n)
        result = self.myPow(y, m >> 1)
        result *= result
        if(m & 0x01 == 1):
            result *= y
        return result

# @lc code=end
