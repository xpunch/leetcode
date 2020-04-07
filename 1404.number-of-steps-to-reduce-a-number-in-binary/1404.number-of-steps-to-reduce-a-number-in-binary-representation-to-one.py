#
# @lc app=leetcode id=1404 lang=python3
#
# [1404] Number of Steps to Reduce a Number in Binary Representation to One
#

# @lc code=start


class Solution:
    def numSteps(self, s: str) -> int:
        if s == '' or s == '0':
            return 1
        result, num = 0, self.getNum(s)
        while num != 1:
            if num % 2 == 0:
                num = num//2
            else:
                num += 1
            result += 1
        return result

    def getNum(self, s: str) -> int:
        n = len(s)
        if n == 0:
            return 0
        num = 0
        for i in range(n):
            if s[i] == '0':
                continue
            num += (1 << (n-i-1))
        return num
# @lc code=end
