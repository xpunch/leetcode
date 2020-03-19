#
# @lc app=leetcode id=69 lang=python3
#
# [69] Sqrt(x)
#

# @lc code=start


class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0
        l, r, y = 1, x, 1
        while l < r:
            y = (l+r)//2
            if y == l:
                return y
            square = y*y
            if square > x:
                r = y
            elif square == x:
                return y
            else:
                l = y
        return y
# @lc code=end
