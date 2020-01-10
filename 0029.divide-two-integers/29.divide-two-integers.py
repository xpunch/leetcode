#
# @lc app=leetcode id=29 lang=python3
#
# [29] Divide Two Integers
#

# @lc code=start


class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == 0:
            return 0
        if divisor == 1:
            return dividend
        if divisor == -1:
            if dividend == -2147483648:
                return 2147483647
            else:
                return -dividend
        quotient = 0
        is_quotient_positive = (dividend > 0) == (divisor > 0)
        dvd, dvs = abs(dividend), abs(divisor)
        while dvd >= dvs:
            i = 1
            tmp = dvs
            while tmp << 1 <= dvd:
                i <<= 1
                tmp <<= 1
            dvd -= tmp
            quotient += i
        if is_quotient_positive:
            return quotient
        return -quotient
# @lc code=end
