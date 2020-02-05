#
# @lc app=leetcode id=43 lang=python3
#
# [43] Multiply Strings
#

# @lc code=start


class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        n1, n2 = len(num1), len(num2)
        if n1 == 0 or n2 == 0:
            return ''
        if num1 == '0' or num2 == '0':
            return '0'
        num = [0 for i in range(n1+n2)]
        for i in range(n1-1, -1, -1):
            d1 = ord(num1[i])-48
            for j in range(n2-1, -1, -1):
                d2 = ord(num2[j])-48
                num[i+j+1] += d1*d2
                num[i+j] += num[i+j+1]//10
                num[i+j+1] = num[i+j+1] % 10
        start = 0
        for i in range(n1+n2):
            if num[i] > 0:
                start = i
                break
        result = ''
        for i in range(start, n1+n2):
            result += chr(num[i]+48)
        return result

# @lc code=end
